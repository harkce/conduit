package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/harkce/conduit/internal/config"
	"github.com/harkce/conduit/internal/handler"
)

func main() {
	cfg := config.Load()
	mux := handler.NewRouter()

	slog.Info("listening", "addr", cfg.Addr)
	if err := http.ListenAndServe(cfg.Addr, mux); err != nil {
		slog.Error("failed to start server", "error", err)
		os.Exit(1)
	}
}
