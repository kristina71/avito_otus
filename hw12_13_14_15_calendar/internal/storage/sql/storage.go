package sqlstorage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gofrs/uuid"

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

/* func New(db *sqlx.DB) *Event {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"cfg.Host", "cfg.Port", "cfg.User", "cfg.Password", "cfg.DBName",
	)
	return &Event{
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

func (s *Storage) Create(ctx context.Context, ev *storage.Event) error {
	/* events, err := s.GetEventsPerDay(ctx, ev.StartAt)
	for _, event := range events {
		if event.ID == ev.ID {
			return storage.Event{}, storage.ErrEventExists
		} else if ev.StartAt == event.StartAt || ev.StartAt.Sub(event.StartAt) < ev.Duration {
			return storage.Event{}, storage.ErrDateBusy
		}
	}
	*/
	query, args, err := psql.Insert(tableName).
		Columns("id", "title", "start_at", "duration", "description", "user_id", "remind_at").
		Values(ev.ID, ev.Title, ev.StartAt, ev.Duration, ev.Description, ev.UserID, ev.RemindAt).
		Suffix("RETURNING id").ToSql()
	// s.logger.Info(query)
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = s.db.Exec(query, args...)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *Storage) Get(ctx context.Context, ev *storage.Event) (uuid.UUID, error) {
	query, args, err := psql.Select("id", "title", "start_at", "duration", "description", "user_id", "remind_at").
		From(tableName).Where(squirrel.Eq{"id": ev.ID}).ToSql()
	if err != nil {
		log.Println(err)
		return uuid.Nil, err
	}

	events := storage.Event{}
	err = s.db.Get(&events, query, args...)

	if errors.Is(err, sql.ErrNoRows) {
		return uuid.Nil, storage.ErrNoRows
	}
	return events.ID, nil
}

func (s *Storage) Update(ctx context.Context, event *storage.Event) error {
	query, args, err := psql.Update(tableName).
		Set("title", event.Title).
		Set("start_at", event.StartAt).
		Set("duration", event.Duration).
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

func (s *Storage) Delete(ctx context.Context, id uuid.UUID) error {
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
	query, _, err := psql.Select("id", "title", "start_at", "duration", "description", "user_id", "remind_at").
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

func (s *Storage) GetEventsPerDay(ctx context.Context, date time.Time) ([]storage.Event, error) {
	query, _, err := psql.Select("id", "title", "start_at", "duration", "description", "user_id", "remind_at").
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

func (s *Storage) GetEventsPerWeek(ctx context.Context, date time.Time) ([]storage.Event, error) {
	query, _, err := psql.Select("id", "title", "start_at", "duration", "description", "user_id", "remind_at").
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

func (s *Storage) GetEventsPerMonth(ctx context.Context, date time.Time) ([]storage.Event, error) {
	query, _, err := psql.Select("id", "title", "start_at", "duration", "description", "user_id", "remind_at").
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

func (s *Storage) ListForScheduler(ctx context.Context, remindFor time.Duration, period time.Duration) ([]storage.Notification, error) {
	from := time.Now().Add(remindFor)
	to := from.Add(period)

	query, _, err := psql.Select("id", "title", "start_at", "user_id").
		From(tableName).Where(squirrel.Expr("start_at BETWEEN $1 AND $2", from, to)).
		OrderBy("id desc").ToSql()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	events := []storage.Notification{}
	err = s.db.Select(&events, query)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return events, nil
}
