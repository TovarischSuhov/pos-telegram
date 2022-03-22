package storage

import (
	"flag"
	"log"
	"sync"

	"github.com/TovarischSuhov/pos-telegram/domain"
	"go.etcd.io/bbolt"
)

var filepath = flag.String("bbolt-path", "/tmp/bbolt", "Bbolt storage filepath")

var (
	messagesBucket = []byte("messages")
	systemBucket   = []byte("system")
	offsetField    = []byte("offset")
)

var (
	_storage *bboltStorage
	once     sync.Once
)

type bboltStorage struct {
	db *bbolt.DB
}

func GetStorage() Storage {
	once.Do(func() {
		db, err := bbolt.Open(*filepath, 0o600, nil)
		if err != nil {
			panic(err)
		}
		_storage = &bboltStorage{db: db}
	})
	return _storage
}

func (s *bboltStorage) GetLastOffset() domain.Offset {
	tx, err := s.db.Begin(true)
	if err != nil {
		log.Printf("[Error] fail to create tx : %+v\n", err)
		return 0
	}
	defer tx.Rollback()
	bucket, err := tx.CreateBucketIfNotExists(messagesBucket)
	if err != nil {
		log.Printf("[Error] fail to createBucket : %+v\n", err)
		return 0
	}
	result := bucket.Get(offsetField)
	return domain.Offset(result)
}

func (s *bboltStorage) SetNewOffset(offset domain.Offset) error {
	tx, err := s.db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	bucket, err := tx.CreateBucketIfNotExists(messagesBucket)
	if err != nil {
		return err
	}
	err = bucket.Put(offsetField, []byte(offset))
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (s *bboltStorage) AddMessageToUserID(
	message string,
	messageID domain.MessageID,
	userID domain.UserID,
) error {
	tx, err := s.db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	bucket, err := tx.CreateBucketIfNotExists(messagesBucket)
	if err != nil {
		return err
	}
	messages, err := bucket.CreateBucketIfNotExists([]byte(userID))
	if err != nil {
		return err
	}
	err = messages.Put([]byte(messageID), []byte(message))
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *bboltStorage) GetMessagesByUserID(
	userID domain.UserID,
) ([]string, error) {
	tx, err := s.db.Begin(false)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	bucket := tx.Bucket([]byte(userID))
	if bucket == nil {
		return nil, nil
	}
	result := []string{}
	cursor := bucket.Cursor()
	for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
		result = append(result, string(v))
	}
	return result, nil
}

func (s *bboltStorage) CleanupMessagesByUserID(userID domain.UserID) error {
	tx, err := s.db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	err = tx.DeleteBucket([]byte(userID))
	if err != nil {
		return err
	}
	return tx.Commit()
}
