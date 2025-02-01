package handler

import (
	"errors"
	"net/http"

	"github.com/lyricat/goutils/httphelper/render"
	"github.com/quailyquaily/gzk9000/handler/sys"
	"github.com/quailyquaily/gzk9000/session"

	"github.com/go-chi/chi"
)

func New(
	cfg Config,
	session *session.Session,
) Server {

	return Server{
		cfg:     cfg,
		session: session,
	}
}

type (
	Config struct {
	}

	Server struct {
		cfg Config

		session *session.Session
	}
)

func (s Server) HandleRest() http.Handler {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Get("/", sys.RenderRoot())
	})

	r.Route("/_hc", func(r chi.Router) {
		// return disallow all
		r.Get("/", sys.RenderHealthCheck())
		r.Head("/", sys.RenderHealthCheck())
	})

	r.Route("/robots.txt", func(r chi.Router) {
		// return disallow all
		r.Get("/", sys.RenderRobots())
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		render.NotFound(w, http.StatusNotFound, errors.New("not found"))
	})

	return r
}
