package sqlstorage

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Storage struct {
	dsn string
	db  *sqlx.DB
}

func New() *Storage {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"cfg.Host", "cfg.Port", "cfg.User", "cfg.Password", "cfg.DBName",
	)
	return &Storage{
		dsn: dsn,
	}
}

func (s *Storage) Connect(ctx context.Context) error {
	db, err := sqlx.ConnectContext(ctx, "pgx", s.dsn)
	if err != nil {
		return err
	}

	s.db = db
	zap.L().Info("connect to db")
	return nil
}

func (s *Storage) Close(ctx context.Context) error {
	zap.L().Info("close sql connection to db")

	return s.db.Close()
}
