package models

import "github.com/TovarischSuhov/pos-telegram/domain"

type User struct {
	ID           domain.UserID `json:"id"`
	IsBot        bool          `json:"is_bot"`
	FirstName    string        `json:"first_name"`
	LastName     string        `json:"last_name"`
	Username     string        `json:"username"`
	LanguageCode string        `json:"language_code"`
}
