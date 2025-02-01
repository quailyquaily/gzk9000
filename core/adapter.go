package core

import (
	"time"

	"github.com/lyricat/goutils/structs"
)

const (
	AdapterTypeTelegram = "telegram"
)

type (
	Adapter struct {
		ID   uint64 `json:"id"`
		Type string `json:"type"`

		Settings structs.JSONMap `json:"settings"`

		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
