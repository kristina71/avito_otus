package sqlstorage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/storage"
)

type Storage struct {
	// dsn string
	db *sqlx.DB
	// logger logger.Logger
}

func NewDBStorage() *Storage {
	return &Storage{}
}

func New(db *sqlx.DB) *Storage {
	return &Storage{db: db}
}

const (
	tableName = "events"
)

/* func New(db *sqlx.DB) *Storage {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"cfg.Host", "cfg.Port", "cfg.User", "cfg.Password", "cfg.DBName",
	)
	return &Storage{
		dsn: dsn,
		db:  db,
	}
}*/

func (s *Storage) Connect(ctx context.Context, dsn string) (err error) {
	s.db, err = sqlx.ConnectContext(ctx, "pgx", dsn)
	//	s.db.SetMaxOpenConns(20)
	//	s.db.SetMaxIdleConns(5)
	//	s.db.SetConnMaxLifetime(time.Minute * 3)
	if err != nil {
		return fmt.Errorf("failed to connect to db: %w", err)
	}

	return s.db.PingContext(ctx)
}

func (s *Storage) Close(ctx context.Context) error {
	if err := s.db.Close(); err != nil {
		return fmt.Errorf("error during db connection pool closing: %w", err)
	}
	return nil
}

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

func (s *Storage) Create(ctx context.Context, ev storage.Event) (storage.Event, error) {
	/* var cntr int
	err := s.db.GetContext(ctx, &cntr, "select count(*) from events where start_at < $1 and end_at > $1", ev.StartAt)
	if err != nil {
		return storage.Event{}, err
	}
	if cntr > 0 {
		return storage.Event{}, storage.ErrEventExists
	}
	*/
	query, args, err := psql.Insert(tableName).
		Columns("title", "start_at", "end_at", "description", "user_id", "remind_at").
		Values(ev.Title, ev.StartAt, ev.EndAt, ev.Description, ev.UserID, ev.RemindAt).
		Suffix("RETURNING \"id\"").ToSql()
	// s.logger.Info(query)
	if err != nil {
		log.Println(err)
		return storage.Event{}, err
	}
	fmt.Println(query)
	var id int
	err = s.db.QueryRow(query, args...).Scan(&id)
	if err != nil {
		log.Println(err)
		return storage.Event{}, err
	}

	ev.ID = id

	return ev, nil
}

func (s *Storage) Get(ctx context.Context, id int) (storage.Event, error) {
	query, args, err := psql.Select("id", "title", "start_at", "end_at", "description", "user_id", "remind_at").
		From(tableName).Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		log.Println(err)
		return storage.Event{}, err
	}

	events := storage.Event{}
	err = s.db.Get(&events, query, args...)

	if errors.Is(err, sql.ErrNoRows) {
		return storage.Event{}, storage.ErrNoRows
	}
	return events, nil
}

func (s *Storage) Update(ctx context.Context, event storage.Event) error {
	query, args, err := psql.Update(tableName).
		Set("title", event.Title).
		Set("start_at", event.StartAt).
		Set("end_at", event.EndAt).
		Set("description", event.Description).
		Set("remind_at", event.RemindAt).
		Where(squirrel.Eq{"id": event.ID}).ToSql()
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = s.db.Exec(query, args...)
	return err
}

func (s *Storage) Delete(ctx context.Context, id int) error {
	query, args, err := psql.Delete(tableName).Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = s.db.Exec(query, args...)
	return err
}

func (s *Storage) DeleteAll(ctx context.Context) error {
	query, args, err := psql.Delete(tableName).ToSql()
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = s.db.Exec(query, args...)
	return err
}

func (s *Storage) ListAll(ctx context.Context) ([]storage.Event, error) {
	query, _, err := psql.Select("id", "title", "start_at", "end_at", "description", "user_id", "remind_at").
		From(tableName).OrderBy("id desc").ToSql()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	events := []storage.Event{}
	err = s.db.Select(&events, query)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return events, nil
}

func (s *Storage) ListDay(ctx context.Context, date time.Time) ([]storage.Event, error) {
	query, _, err := psql.Select("id", "title", "start_at", "end_at", "description", "user_id", "remind_at").
		From(tableName).Where(squirrel.Expr("start_at BETWEEN $1 AND $1 + (interval '1d')", date)).OrderBy("id desc").ToSql()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	events := []storage.Event{}
	err = s.db.Select(&events, query)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return events, nil
}

func (s *Storage) ListWeek(ctx context.Context, date time.Time) ([]storage.Event, error) {
	query, _, err := psql.Select("id", "title", "start_at", "end_at", "description", "user_id", "remind_at").
		From(tableName).Where(squirrel.Expr("start_at BETWEEN $1 AND $1 + (interval '7d')", date)).OrderBy("id desc").ToSql()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	events := []storage.Event{}
	err = s.db.Select(&events, query)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return events, nil
}

func (s *Storage) ListMonth(ctx context.Context, date time.Time) ([]storage.Event, error) {
	query, _, err := psql.Select("id", "title", "start_at", "end_at", "description", "user_id", "remind_at").
		From(tableName).Where(squirrel.Expr("start_at BETWEEN $1 AND $1 + (interval '1months')", date)).
		OrderBy("id desc").ToSql()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	events := []storage.Event{}
	err = s.db.Select(&events, query)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return events, nil
}

func (s *Storage) IsTimeBusy(ctx context.Context, start, stop time.Time, excludeID int) (bool, error) {
	return true, nil
}
