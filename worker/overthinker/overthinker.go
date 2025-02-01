package overthinker

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/lyricat/goutils/ai"
	"github.com/quailyquaily/gzk9000/core"
	"gorm.io/gorm"
)

type Worker struct {
	cfg       Config
	aiInst    *ai.Instant
	memslicez core.MemsliceService

	studygoals core.StudygoalStore
	agents     core.AgentStore
}

type Config struct {
	Name    string
	AIDebug bool
}

func New(
	cfg Config,
	aiInst *ai.Instant,
	memslicez core.MemsliceService,

	studygoals core.StudygoalStore,
	agents core.AgentStore,
) *Worker {
	return &Worker{
		cfg:        cfg,
		aiInst:     aiInst,
		memslicez:  memslicez,
		studygoals: studygoals,
		agents:     agents,
	}
}

func (w *Worker) Run(ctx context.Context) error {
	slog.Info("[overthinker] worker started")
	dur := time.Millisecond
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(dur):
			if err := w.run(ctx); err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					slog.Error("[overthinker] failed to run", "error", err)
					dur = time.Second * 30
				} else {
					dur = time.Second * 5
				}
			} else {
				// wait for 60s
				dur = time.Second * 60
			}
		}
	}
}

func (w *Worker) run(ctx context.Context) error {
	return nil
}
