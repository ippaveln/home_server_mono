package http_server

import (
	"fmt"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/docgen"
	"github.com/go-chi/render"
	"github.com/ippaveln/home_server_mono/internal/app/config"
	"github.com/ippaveln/home_server_mono/internal/services/connector"
	"github.com/ippaveln/home_server_mono/internal/services/service"
)

func (server *HttpServer) Run(config *config.Config, conn *connector.Connector, log *slog.Logger, wg *sync.WaitGroup) {
	server.port = int(config.Server.Ports.Http)
	server.log = log
	server.wg = wg
	server.connector = conn

	r := chi.NewRouter()
	r.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Recoverer,
		middleware.URLFormat,
		middleware.Timeout(60*time.Second),
		middleware.Heartbeat("/ping"),
		newLoggerMiddleWare(log),
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		}),
		render.SetContentType(render.ContentTypeJSON),
	)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	r.Post("/ha/ya_station", postDataToYandexStation(server))

	server.router = r

	go http.ListenAndServe(":8013", r)

	fmt.Println(docgen.MarkdownRoutesDoc(r, docgen.MarkdownOpts{}))
	fmt.Println(docgen.JSONRoutesDoc(r))

}

func (server *HttpServer) Status() service.Status {
	return server.status
}

func (server *HttpServer) Stop() {}

func newLoggerMiddleWare(log *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		log = log.With(
			slog.String("component", "http_server/middleware/logger"),
		)

		log.Info("logger middleware  enabled")

		log.Info("logger middleware enabled")

		// код самого обработчика
		fn := func(w http.ResponseWriter, r *http.Request) {
			// собираем исходную информацию о запросе
			entry := log.With(
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.String("remote_addr", r.RemoteAddr),
				slog.String("user_agent", r.UserAgent()),
				slog.String("request_id", middleware.GetReqID(r.Context())),
			)

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			t1 := time.Now()

			defer func() {
				entry.Info("request completed",
					slog.Int("status", ww.Status()),
					slog.Int("", ww.BytesWritten()),
					slog.String("duration", time.Since(t1).String()),
				)
			}()

			next.ServeHTTP(ww, r)
		}

		return http.HandlerFunc(fn)
	}
}

func postDataToYandexStation(server *HttpServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// data := &YandexStation{}
	}
}
