package models

import "github.com/TovarischSuhov/pos-telegram/domain"

type Message struct {
	MessageID  domain.MessageID `json:"message_id"`
	From       User             `json:"from"`
	SenderChat Chat             `json:"sender_chat"`
	Date       int64            `json:"date"`
	Chat       Chat             `json:"chat"`
	Test       string           `json:"test"`
}
