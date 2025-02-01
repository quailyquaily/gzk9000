package core

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"time"

	"github.com/lib/pq"
)

const (
	MemsliceResponseTypeMonoLog = iota
	MemsliceResponseTypeDialog
)

type UInt64Array []uint64

// Scan implements the sql.Scanner interface.
// It first scans into a pq.Int64Array, then converts each int64 to uint64.
func (a *UInt64Array) Scan(src interface{}) error {
	if src == nil {
		*a = nil
		return nil
	}

	// Let pq.Int64Array handle the actual parsing of the Postgres array string (e.g. "{1,2,3}").
	var int64Arr pq.Int64Array
	if err := int64Arr.Scan(src); err != nil {
		return fmt.Errorf("UInt64Array scan error: %w", err)
	}

	// Convert each int64 to uint64 (failing if negative).
	out := make(UInt64Array, len(int64Arr))
	for i, v := range int64Arr {
		if v < 0 {
			return fmt.Errorf("UInt64Array: cannot convert negative value %d to uint64", v)
		}
		out[i] = uint64(v)
	}

	*a = out
	return nil
}

// Value implements the driver.Valuer interface.
// It converts the []uint64 back to a pq.Int64Array for insertion into a BIGINT[] column.
func (a UInt64Array) Value() (driver.Value, error) {
	// We need to ensure that none of the uint64 values exceed max int64.
	int64Arr := make(pq.Int64Array, len(a))
	for i, v := range a {
		if v > math.MaxInt64 {
			return nil, fmt.Errorf("UInt64Array: value %d exceeds max int64", v)
		}
		int64Arr[i] = int64(v)
	}
	// Let pq.Int64Array handle final conversion to driver.Value for Postgres.
	return int64Arr.Value()
}

type (
	Memslice struct {
		ID uint64 `json:"id"`

		ResponseType int `json:"response_type"`

		AgentID   uint64 `json:"agent_id"`
		SpeakerID uint64 `json:"speaker_id"`

		IsMonolog bool `json:"is_monolog"`

		IncludedFactIDs UInt64Array `gorm:"type:BIG[]" json:"included_fact_ids"`
		ExternalFactIDs UInt64Array `gorm:"type:BIG[]" json:"external_fact_ids"`

		RelatedMemsliceIDs UInt64Array `gorm:"type:BIG[]" json:"related_memslice_ids"`

		Content string `json:"content"`

		Status int `json:"status"`

		CreatedAt *time.Time `json:"created_at"`
		UpdatedAt *time.Time `json:"updated_at"`

		// Virtual fields
		IncludedFacts []*Fact `gorm:"-" json:"included_facts"`
		ExternalFacts []*Fact `gorm:"-" json:"external_facts"`
	}

	MemsliceStore interface {
		// INSERT INTO @@table (
		//   response_type, agent_id, speaker_id,
		//   is_monolog, i
		//   ncluded_fact_ids, external_fact_ids,
		//   related_memslice_ids,
		//   content, status,
		//   created_at, updated_at
		// ) VALUES (
		//   @memslice.ResponseType, @memslice.AgentID, @memslice.SpeakerID,
		//   @memslice.IsMonolog,
		//   @memslice.IncludedFactIDs, @memslice.ExternalFactIDs,
		//   @memslice.RelatedMemsliceIDs,
		//   @memslice.Content, 0,
		//   NOW(), NOW()
		// ) RETURNING id;
		CreateMemslice(ctx context.Context, memslice *Memslice) error

		// SELECT * FROM @@table
		// WHERE
		//   agent_id = @agentID AND
		//   created_at >= @start AND
		//   created_at <= @end
		// ORDER BY created_at DESC;
		GetMemslicesByRange(ctx context.Context, agentID uint64, start, end *time.Time) ([]*Memslice, error)
	}

	MemsliceService interface {
		GetMemslicesByRange(ctx context.Context, agentID uint64, start, end *time.Time) ([]*Memslice, error)
	}
)
