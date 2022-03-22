package storage

import "github.com/TovarischSuhov/pos-telegram/domain"

type Storage interface {
	AddMessageToUserID(string, domain.MessageID, domain.UserID) error
	GetMessagesByUserID(domain.UserID) ([]string, error)
	CleanupMessagesByUserID(domain.UserID) error
	GetLastOffset() domain.Offset
	SetNewOffset(domain.Offset) error
}
