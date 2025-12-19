package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/danielsonng/ecomgo/internal/env"
	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()
	cfg := config{
		addr: ":8080",
		db: dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING", "host=localhost user=postgres password=postgres dbname=ecom sslmode=disable"),
		},
	}

	//Logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	var conn *pgx.Conn
	// Database
	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)
	logger.Info("Connected to database", "dsn", cfg.db.dsn)

	api := application{
		config: cfg,
		db:     conn,
	}

	h := api.mount()
	if err := api.run(h); err != nil {
		slog.Error("Server Failed to Initialize", "Error", err)
		os.Exit(1)
	}
}
