package timer

import (
	"context"
	"log/slog"
	"time"
)

type (
	Worker struct {
		cfg Config
	}
	Config struct {
	}
)

func New(
	cfg Config,
) *Worker {
	return &Worker{
		cfg: cfg,
	}
}

func (w *Worker) Run(ctx context.Context) error {
	slog.Info("[timer] worker started")
	dur := time.Millisecond
	var ticker uint64
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(dur):
			if err := w.run(ctx, ticker); err != nil {
				dur = time.Second * 30
				ticker += 30
			} else {
				dur = time.Second * 1
				ticker += 1
			}
		}
	}
}

func (w *Worker) run(ctx context.Context, ticker uint64) error {
	// TODO: implement timer worker
	return nil
}
