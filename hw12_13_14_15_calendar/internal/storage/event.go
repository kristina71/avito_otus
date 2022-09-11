package storage

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
)

type Event struct {
	ID uuid.UUID `db:"id" json:"id"`
	// short description
	Title string `json:"title"`
	// date and time
	StartAt time.Time `db:"start_at" json:"startAt"`
	// time duration (end time)
	Duration int `db:"duration"`
	// description
	Description string `db:"description" json:"description"`
	// owner id
	UserID   uuid.UUID `db:"user_id" json:"userId"`
	RemindAt int       `db:"remind_at" json:"remindAt"`
}

type Notification struct {
	ID        uuid.UUID `db:"id" json:"id"`
	Title     string    `db:"title" json:"title"`
	TimeStart time.Time `db:"start_at" json:"start_at"`
	UserID    uuid.UUID `db:"user_id" json:"user_id"`
}

//go:generate mockery --name=Storage --output ./mocks
type Storage interface {
	Create(ctx context.Context, event *Event) error
	Get(ctx context.Context, event *Event) (uuid.UUID, error)
	Update(ctx context.Context, event *Event) error
	Delete(ctx context.Context, id uuid.UUID) error

	DeleteAll(ctx context.Context) error
	ListAll(ctx context.Context) ([]Event, error)
	GetEventsPerDay(ctx context.Context, date time.Time) ([]Event, error)
	GetEventsPerWeek(ctx context.Context, date time.Time) ([]Event, error)
	GetEventsPerMonth(ctx context.Context, date time.Time) ([]Event, error)
	Close(ctx context.Context) error
	ListForScheduler(ctx context.Context, remindFor time.Duration, period time.Duration) ([]Notification, error)
}
