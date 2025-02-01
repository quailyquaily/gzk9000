package looper

import (
	"context"
	"log/slog"

	"github.com/lyricat/goutils/ai"
	"github.com/quailyquaily/gzk9000/core"
	"github.com/quailyquaily/gzk9000/loop"
	"github.com/quailyquaily/gzk9000/loop/telegram"
	"golang.org/x/sync/errgroup"
)

type Worker struct {
	cfg     Config
	loopMap map[string]*loop.LoopService
}

type Config struct {
	TelegramToken string
}

func New(
	cfg Config,
	aiInst *ai.Instant,

	factz core.FactService,

	memslices core.MemsliceStore,
) *Worker {
	loopMap := make(map[string]*loop.LoopService)

	tgAdapater := telegram.New(telegram.Config{
		Token: cfg.TelegramToken,
	}, memslices, factz)

	loopMap["telegrame"] = loop.New(aiInst, tgAdapater)
	return &Worker{
		cfg:     cfg,
		loopMap: loopMap,
	}
}

func (w *Worker) Run(ctx context.Context) error {
	g := errgroup.Group{}

	for _, loopz := range w.loopMap {
		g.Go(func() error {
			slog.Info("[looper] loop started", "adapter_name", loopz.GetAdapterName())
			if err := loopz.Start(ctx); err != nil {
				return err
			}
			return nil
		})
	}

	return g.Wait()
}
