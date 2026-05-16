package pg

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/1DelFin1/testMicroservice/config"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	_defaultMaxPoolSize  = 4
	_defaultConnAttempts = 10
	_defaultConnTimeout  = time.Second
)

type Postgres struct {
	Builder squirrel.StatementBuilderType
	Pool    *pgxpool.Pool
}

func New(cfg *config.Config) (*Postgres, error) {
	poolConfig, err := pgxpool.ParseConfig(cfg.PG.URL)
	if err != nil {
		return nil, fmt.Errorf("pg - New - ParseConfig: %w", err)
	}

	poolConfig.MaxConns = _defaultMaxPoolSize

	var pool *pgxpool.Pool
	for attempts := _defaultConnAttempts; attempts > 0; attempts-- {
		pool, err = pgxpool.NewWithConfig(context.Background(), poolConfig)
		if err == nil {
			break
		}
		time.Sleep(_defaultConnTimeout)
	}

	if err != nil {
		return nil, fmt.Errorf("pg - New - connAttempts == 0: %w", err)
	}

	return &Postgres{
		Pool:    pool,
		Builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}, nil
}

func (p *Postgres) Close(log *slog.Logger) {
	if p.Pool == nil {
		return
	}
	p.Pool.Close()
	log.Info("pg pool closed")
}
