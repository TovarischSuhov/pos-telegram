package main

import (
	"flag"
	"log"

	"github.com/TovarischSuhov/pos-telegram/storage"
	"github.com/TovarischSuhov/pos-telegram/telegram"
	_ "github.com/mailru/easyjson/gen"
)

func main() {
	flag.Parse()
	go PullEvents()
}

func PullEvents() {
	api := telegram.GetAPI()
	store := storage.GetStorage()

	offset := store.GetLastOffset()
	for {
		updates, err := api.GetUpdates(offset)
		if err != nil {
			log.Printf("[ERROR] fail to get updates: %+v", err)
			return
		}
		for _, value := range updates {
			if value.UpdateID > offset {
				offset = value.UpdateID
			}
			err := parseCommand(value.Message)
			if err != nil {
				log.Printf("[ERROR] fail to parse command: %+v", err)
				return
			}
		}
		err := store.SetNewOffset(offset)
		if err != nil {
			log.Printf("[ERROR] fail to update offset: %+v", err)
			return
		}
	}
}

func parseCommand(message string) error {
}
