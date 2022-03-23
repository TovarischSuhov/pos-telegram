package models

type Update struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}

type UpdateResponse []Update
