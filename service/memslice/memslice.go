package memslice

import (
	"context"
	"time"

	"github.com/quailyquaily/gzk9000/core"
)

func New(
	cfg Config,
	memslices core.MemsliceStore,
	facts core.FactStore,
) *service {
	return &service{
		cfg: cfg,

		memslices: memslices,
		facts:     facts,
	}
}

type Config struct {
}

type service struct {
	cfg       Config
	memslices core.MemsliceStore
	facts     core.FactStore
}

func (s *service) GetMemslicesByRange(ctx context.Context, agentID uint64, start, end *time.Time) ([]*core.Memslice, error) {
	mmslcs, err := s.memslices.GetMemslicesByRange(ctx, agentID, start, end)
	if err != nil {
		return nil, err
	}
	s.WithFacts(ctx, mmslcs...)
	return mmslcs, nil
}

func (s *service) WithFacts(ctx context.Context, items ...*core.Memslice) error {
	if items == nil {
		return nil
	}

	for _, item := range items {
		ids := make([]uint64, 0)
		for _, factID := range item.IncludedFactIDs {
			ids = append(ids, factID)
			fts, err := s.facts.GetFactsByIDs(ctx, ids)
			if err != nil {
				return err
			}
			item.IncludedFacts = append(item.IncludedFacts, fts...)
		}
	}

	return nil
}
