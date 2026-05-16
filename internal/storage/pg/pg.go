package pg

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/1DelFin1/testMicroservice/config"
	"github.com/jackc/pgx/v5"
)

func GetPGConn(cfg *config.Config) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), cfg.PG.URL)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	return conn, nil
}

func Shutdown(log *slog.Logger, conn *pgx.Conn) {
	for i := 0; i < 5; i++ {
		if err := conn.Close(context.Background()); err != nil {
			continue
		}
		return
	}
	log.Error("unable to close pg connection")
}
