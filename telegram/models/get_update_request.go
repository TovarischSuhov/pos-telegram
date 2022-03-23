package models

import "github.com/TovarischSuhov/pos-telegram/domain"

type GetUpdateRequest struct {
	Offset         domain.Offset `json:"offset,omitempty"`
	Limit          int64         `json:"limit,omitempty"`
	Timeout        int64         `json:"timeout,omitempty"`
	AllowedUpdates []string      `json:"allowed_updates,omitempty"`
}
