package core

import (
	"context"
	"time"
)

const (
	StudygoalStatusInactive = 0
	StudygoalStatusActive   = 1
)

type (
	Studygoal struct {
		ID      uint64 `json:"id"`
		AgentID uint64 `json:"agent_id"`

		Content string `json:"goal_content"`

		Iteration int `json:"iteration"`

		PriorityScore float64 `json:"priority_score"`

		Status int `json:"status"`

		CreatedAt *time.Time `json:"created_at"`
		UpdatedAt *time.Time `json:"updated_at"`
	}

	StudygoalStore interface {
		// INSERT INTO @@table (
		//   agent_id, content, iteration, status,
		//   created_at, updated_at
		// ) VALUES (
		//   @goal.AgentID, @goal.Content, 0, @goal.Status,
		//   NOW(), NOW()
		// ) RETURNING id;
		CreateStudygoal(ctx context.Context, goal *Studygoal) error

		// SELECT * FROM @@table
		// WHERE id = @id;
		GetStudygoalByID(ctx context.Context, id uint64) (*Studygoal, error)

		// SELECT * FROM @@table
		// WHERE agent_id = @agentID;
		GetStudygoalsByAgentID(ctx context.Context, agentID uint64) ([]*Studygoal, error)

		// SELECT * FROM @@table
		// WHERE agent_id = @agentID AND status = 1;
		GetActiveStudygoalsByAgentID(ctx context.Context, agentID uint64) ([]*Studygoal, error)
	}
)
