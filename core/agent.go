package core

import (
	"context"
	"time"
)

const (
	AgentStatusInactive = 0
	AgentStatusActive   = 1
)

type (
	Agent struct {
		ID          uint64 `json:"id"`
		Name        string `json:"name"`
		DisplayName string `json:"display_name"`

		Status int `json:"status"`

		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	AgentStore interface {
		// SELECT * FROM @@table
		// WHERE id = @id;
		GetAgentByID(ctx context.Context, id uint64) (*Agent, error)

		// SELECT * FROM @@table WHERE status = 1;
		GetAllAgents(ctx context.Context) ([]*Agent, error)
	}
)
