package app

import (
	"context"
	"fmt"
	"time"

	"github.com/gofrs/uuid"

	"github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/storage"
)

type App struct {
	logger  Logger
	storage storage.Storage
}

type Logger interface {
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Debug(msg string)
}

func New(logger Logger, storage storage.Storage) *App {
	return &App{
		logger:  logger,
		storage: storage,
	}
}

func (a *App) Create(
	ctx context.Context, event *storage.Event) error {
	var err error

	if event.Title == "" {
		err = storage.ErrEmptyTitle
		return err
	}
	if event.Duration <= 0 {
		err = storage.ErrWrongDuration
		return err
	}

	id, err := uuid.NewV4()
	if err != nil {
		err = storage.ErrWrongId
		return err
	}

	ev := &storage.Event{
		ID:          id,
		Title:       event.Title,
		StartAt:     event.StartAt,
		Duration:    event.Duration,
		Description: event.Description,
		UserID:      event.UserID,
		RemindAt:    event.RemindAt,
	}

	fmt.Println(ev)
	a.logger.Info(
		fmt.Sprintf(
			"%s from %s to %d created",
			event.Title,
			event.StartAt.Format(time.UnixDate),
			event.Duration,
		),
	)
	return a.storage.Create(ctx, ev)
}

func (a *App) Close(ctx context.Context) error {
	a.logger.Info("storage closing...")
	return a.storage.Close(ctx)
}

func (a *App) Get(ctx context.Context, event *storage.Event) (uuid.UUID, error) {
	a.logger.Info(fmt.Sprintf("event %d found", event.ID))
	return a.storage.Get(ctx, event)
}

func (a *App) Update(ctx context.Context, event *storage.Event) error {
	a.logger.Info(
		fmt.Sprintf(
			"%s from %s to %d has been updated",
			event.Title,
			event.StartAt.Format(time.UnixDate),
			event.Duration,
		),
	)

	return a.storage.Update(ctx, event)
}

func (a *App) Delete(ctx context.Context, id uuid.UUID) error {
	a.logger.Info(
		fmt.Sprintf(
			"event %d has been deleted",
			id,
		),
	)
	return a.storage.Delete(ctx, id)
}

func (a *App) DeleteAll(ctx context.Context) error {
	a.logger.Info("all events have been deleted")
	return a.storage.DeleteAll(ctx)
}

func (a *App) ListAll(ctx context.Context) ([]storage.Event, error) {
	a.logger.Info("get all events")
	return a.storage.ListAll(ctx)
}

func (a *App) GetEventsPerDay(ctx context.Context, day time.Time) ([]storage.Event, error) {
	a.logger.Info(
		fmt.Sprintf(
			"get events by day %s",
			day.Format(time.UnixDate),
		),
	)
	return a.storage.GetEventsPerDay(ctx, day)
}

func (a *App) GetEventsPerWeek(ctx context.Context, beginDate time.Time) ([]storage.Event, error) {
	a.logger.Info(
		fmt.Sprintf(
			"get events by week %s",
			beginDate.Format(time.UnixDate),
		),
	)
	return a.storage.GetEventsPerWeek(ctx, beginDate)
}

func (a *App) GetEventsPerMonth(ctx context.Context, beginDate time.Time) ([]storage.Event, error) {
	a.logger.Info(
		fmt.Sprintf(
			"get events by month %s",
			beginDate.Format(time.UnixDate),
		),
	)
	return a.storage.GetEventsPerMonth(ctx, beginDate)
}
