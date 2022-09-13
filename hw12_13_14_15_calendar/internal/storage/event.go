package storage

import (
	"context"
	"time"
)

type Event struct {
	ID int `db:"id" json:"id"`
	// short description
	Title string `json:"title"`
	// date and time
	StartAt time.Time `db:"start_at" json:"startAt"`
	// time duration (end time)
	EndAt time.Time `db:"end_at" json:"endAt"`
	// description
	Description string `db:"description" json:"description"`
	// owner id
	UserID   int       `db:"user_id" json:"userId"`
	RemindAt time.Time `db:"remind_at" json:"remindAt"`
}

//go:generate mockery --name=Storage --output ./mocks
type Storage interface {
	Create(ctx context.Context, event Event) (Event, error)
	Get(ctx context.Context, id int) (Event, error)
	Update(ctx context.Context, event Event) error
	Delete(ctx context.Context, id int) error

	DeleteAll(ctx context.Context) error
	ListAll(ctx context.Context) ([]Event, error)
	ListDay(ctx context.Context, date time.Time) ([]Event, error)
	ListWeek(ctx context.Context, date time.Time) ([]Event, error)
	ListMonth(ctx context.Context, date time.Time) ([]Event, error)
	Close(ctx context.Context) error
	IsTimeBusy(ctx context.Context, start, stop time.Time, excludeID int) (bool, error)
}
