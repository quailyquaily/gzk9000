package sys

import (
	"net/http"

	"github.com/lyricat/goutils/httphelper/render"
	"github.com/quailyquaily/gzk9000/store"
)

func RenderHealthCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result := map[string]string{"status": "ok"}
		// query db for testing connection
		if err := store.Transaction(func(tx *store.Handler) error {
			sql := tx.Exec("SELECT 1")
			if sql.Error != nil {
				return sql.Error
			}
			return nil
		}); err != nil {
			render.Error(w, http.StatusInternalServerError, err)
			return
		}
		render.JSON(w, result)
	}
}

func RenderRobots() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// disable indexing
		render.Text(w, []byte("User-agent: *\nDisallow: /\n"))
	}
}

func RenderRoot() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.Html(w, []byte(`
			<html><head><title>gzk9000</title></head>
				<body>
					<p>gzk9000 service, standing by.</p>
					<div>
					<pre>
        ___________________________________
       |                                   |
       |             G Z K  9 0 0 0        |
       |___________________________________|
       |                                   |
       |             .--------.            |
       |             | ('  ') |            |
       |             | ('  ') |            |
       |             | ('  ') |            |
       |             | ('  ') |            |
       |             | ( __ ) |            |
       |             '--------'            |
       |___________________________________|

					</pre>
					</div>
				</body>
			</html>`))
	}
}
