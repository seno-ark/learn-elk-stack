package main

import (
	"context"
	"hb/app"
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"

	"go.elastic.co/apm/module/apmhttp/v2"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		slog.Info("Failed to read .env file")
	}
}

func main() {
	mux := apmhttp.Wrap(Mux())

	server := &http.Server{
		Addr:         ":9090",
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      mux,
	}

	app.Logger.Info("Starting server", "port", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func Mux() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.IndexHandler)

	var router http.Handler = mux
	router = LogMiddleware(router)

	return router
}

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := uuid.NewString()

		app.Logger.Info("API",
			slog.Any(r.Method, r.RequestURI),
			slog.Any("reqID", reqID),
		)
		ctx := context.WithValue(r.Context(), app.ContextRequestID, reqID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
