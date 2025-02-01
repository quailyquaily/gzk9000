package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/lyricat/goutils/ai"
	"github.com/quailyquaily/gzk9000/handler"
	memsliceZ "github.com/quailyquaily/gzk9000/service/memslice"
	"github.com/quailyquaily/gzk9000/session"
	"github.com/quailyquaily/gzk9000/store/agent"
	"github.com/quailyquaily/gzk9000/store/fact"
	"github.com/quailyquaily/gzk9000/store/memslice"
	"github.com/quailyquaily/gzk9000/store/studygoal"
	"github.com/quailyquaily/gzk9000/worker"
	"github.com/quailyquaily/gzk9000/worker/goalfinder"

	"github.com/quailyquaily/gzk9000/store"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
)

// newserverCmd creates the 'server' subcommand
func NewCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "Run the server that dispatches tasks to proxies and manages results",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			se := session.From(ctx)

			h := store.WithContext(ctx)

			facts := fact.New(h)
			agents := agent.New(h)
			memslices := memslice.New(h)
			studygoals := studygoal.New(h)

			memslicez := memsliceZ.New(memsliceZ.Config{}, memslices, facts)

			aiInst := ai.New(ai.Config{
				Provider:        "susanoo",
				SusanooEndpoint: viper.GetString("susanoo.endpoint"),
				SusanooApiKey:   viper.GetString("susanoo.api_key"),
				Debug:           true,
			})

			// handle signals
			ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
			defer stop()

			var svr *http.Server
			g, _ := errgroup.WithContext(ctx)

			workers := []worker.Worker{
				goalfinder.New(goalfinder.Config{}, aiInst, memslicez, studygoals, agents),
			}

			for idx := range workers {
				w := workers[idx]
				g.Go(func() error {
					return w.Run(ctx)
				})
			}

			g.Go(func() error {
				var err error
				mux := chi.NewMux()
				mux.Use(middleware.Recoverer)
				mux.Use(middleware.StripSlashes)
				mux.Use(cors.AllowAll().Handler)
				mux.Use(middleware.Logger)
				mux.Use(middleware.NewCompressor(5).Handler)
				{
					restsvr := handler.New(handler.Config{}, se)
					restHandler := restsvr.HandleRest()
					mux.Mount("/", restHandler)
				}

				port := 8080
				if len(args) > 0 {
					port, err = strconv.Atoi(args[0])
					if err != nil {
						port = 8080
					}
				}

				// launch server
				if err != nil {
					panic(err)
				}
				addr := fmt.Sprintf(":%d", port)

				svr = &http.Server{
					Addr:    addr,
					Handler: mux,
				}

				slog.Info("[server] run httpd server", "addr", addr)
				if err := svr.ListenAndServe(); err != http.ErrServerClosed {
					slog.Error("[server] server aborted", "error", err)
				}
				return nil
			})

			g.Go(func() error {
				<-ctx.Done() // block until ctx is canceled
				slog.Info("[server] shutting down gracefully...")

				// We use a fresh context for Shutdown with a timeout
				shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				defer cancel()

				if err := svr.Shutdown(shutdownCtx); err != nil {
					slog.Error("[server] forced to shutdown", "error", err)
					return err
				}
				slog.Info("[server] graceful shutdown complete")
				return nil
			})

			if err := g.Wait(); err != nil {
				slog.Error("[server] run httpd & worker", "error", err)
			}
			slog.Info("[server] shutdown totally complete")
		},
	}
}
