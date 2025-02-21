package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)


type application struct {
	logger *slog.Logger
}


func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	// loggerHandler := slog.NewTextHandler(os.Stdout, nil)
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
		AddSource: true,
	}))

	app := application{
		logger: logger,
	}

	mux := app.routes()

	logger.Info("starting server", "addr", *addr)

	// Use the default http server for this basic example.
	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
