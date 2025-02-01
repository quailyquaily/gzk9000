package core

import (
	"context"
	"time"
)

type (
	Fact struct {
		ID uint64 `json:"id"`

		AgentID uint64 `json:"agent_id"`

		Content string `json:"content"`

		Sentiment float64 `json:"sentiment"`

		CreatedAt *time.Time `json:"created_at"`
		UpdatedAt *time.Time `json:"updated_at"`
	}

	FactStore interface {
		// INSERT INTO @@table (
		//   agent_id, content, sentiment,
		//   created_at, updated_at
		// ) VALUES (
		//   @fact.AgentID, @fact.Content, @fact.Sentiment,
		//   NOW(), NOW()
		// ) RETURNING id;
		CreateFact(ctx context.Context, fact *Fact) error

		// SELECT * FROM @@table
		// WHERE id = @id;
		GetFactByID(ctx context.Context, id uint64) (*Fact, error)

		// SELECT * FROM @@table
		// WHERE id IN (@ids);
		GetFactsByIDs(ctx context.Context, ids []uint64) ([]*Fact, error)
	}

	FactService interface {
		CreateFact(ctx context.Context, fact *Fact) error
	}
)
