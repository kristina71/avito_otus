package app

import (
	"context"
	"fmt"
	"time"

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
	ctx context.Context, userID int, title string,
	description string, start time.Time, stop time.Time, remindAt time.Time,
) (storage.Event, error) {
	var err error
	if userID == 0 {
		err = storage.ErrNoUserID
		return storage.Event{}, err
	}
	if title == "" {
		err = storage.ErrEmptyTitle
		return storage.Event{}, err
	}
	if start.After(stop) {
		start, stop = stop, start
	}
	if time.Now().After(start) {
		err = storage.ErrStartInPast
		return storage.Event{}, err
	}
	/* isBusy, err := a.storage.IsTimeBusy(ctx, start, stop, 0)
	if err != nil {
		return storage.Event{}, err
	}
	if isBusy {
		err = storage.ErrDateBusy
		return storage.Event{}, err
	}
	*/
	event := storage.Event{
		ID:          userID,
		Title:       title,
		Description: description,
		StartAt:     start,
		EndAt:       stop,
		RemindAt:    remindAt,
	}

	a.logger.Info(
		fmt.Sprintf(
			"%s from %s to %s created",
			event.Title,
			event.StartAt.Format(time.UnixDate),
			event.EndAt.Format(time.UnixDate),
		),
	)
	return a.storage.Create(ctx, event)
}

func (a *App) Close(ctx context.Context) error {
	a.logger.Info("storage closing...")
	return a.storage.Close(ctx)
}

func (a *App) Get(ctx context.Context, id int) (storage.Event, error) {
	a.logger.Info(fmt.Sprintf("event %d found", id))
	return a.storage.Get(ctx, id)
}

func (a *App) Update(ctx context.Context, event storage.Event) error {
	a.logger.Info(
		fmt.Sprintf(
			"%s from %s to %s has been updated",
			event.Title,
			event.StartAt.Format(time.UnixDate),
			event.EndAt.Format(time.UnixDate),
		),
	)

	return a.storage.Update(ctx, event)
}

func (a *App) Delete(ctx context.Context, id int) error {
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

func (a *App) ListDay(ctx context.Context, date time.Time) ([]storage.Event, error) {
	a.logger.Info(
		fmt.Sprintf(
			"get events by day %s",
			date.Format(time.UnixDate),
		),
	)
	return a.storage.ListDay(ctx, date)
}

func (a *App) ListWeek(ctx context.Context, date time.Time) ([]storage.Event, error) {
	a.logger.Info(
		fmt.Sprintf(
			"get events by week %s",
			date.Format(time.UnixDate),
		),
	)
	return a.storage.ListWeek(ctx, date)
}

func (a *App) ListMonth(ctx context.Context, date time.Time) ([]storage.Event, error) {
	a.logger.Info(
		fmt.Sprintf(
			"get events by month %s",
			date.Format(time.UnixDate),
		),
	)
	return a.storage.ListMonth(ctx, date)
}

/*
func (a *App) IsTimeBusy(ctx context.Context, date time.Time) (bool, error) {
	return a.storage.IsTimeBusy(ctx, date)
}
*/
/*

func startOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

func endOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

func startOfWeek(t time.Time) time.Time {
	return t.Truncate(24 * 7 * time.Hour)
}

func endOfWeek(t time.Time) time.Time {
	return startOfWeek(t).AddDate(0, 0, 7).Add(-time.Nanosecond)
}

func startOfMonth(t time.Time) time.Time {
	y, m, _ := t.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, t.Location())
}

func endOfMonth(t time.Time) time.Time {
	return startOfMonth(t).AddDate(0, 1, 0).Add(-time.Nanosecond)
}
*/
