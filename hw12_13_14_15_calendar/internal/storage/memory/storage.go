package memorystorage

import (
	"context"
	"sort"
	"sync"
	"time"

	"github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/logger"

	"github.com/gofrs/uuid"
	"github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/storage"
)

type Storage struct {
	mu        sync.Mutex
	mapEvents map[uuid.UUID]storage.Event
	logger    logger.Logger
}

func New() *Storage {
	return &Storage{
		mu:        sync.Mutex{},
		mapEvents: make(map[uuid.UUID]storage.Event),
	}
}

type Repository struct {
	repo storage.Storage
}

func New1(repo storage.Storage) *Repository {
	return &Repository{repo: repo}
}

func (s *Storage) Connect(_ context.Context) error {
	s.logger.Info("connected to memory storage")
	return nil
}

func (s *Storage) Close(_ context.Context) error {
	s.logger.Info("closing memory storage")
	return nil
}

func (s *Storage) Create(ctx context.Context, event *storage.Event) error {
	select {
	case <-ctx.Done():
		return storage.ErrCanceledByContext
	default:
		s.mu.Lock()
		defer s.mu.Unlock()
		for _, evValue := range s.mapEvents {
			if event.StartAt.Equal(evValue.StartAt) && event.ID != evValue.ID {
				return storage.ErrDateBusy
			} else if event.StartAt.Equal(evValue.StartAt) && event.ID == evValue.ID {
				return storage.ErrEventExists
			}
		}
		s.mapEvents[event.ID] = *event
	}
	return nil
}

func (s *Storage) Get(ctx context.Context, event *storage.Event) (uuid.UUID, error) {
	select {
	case <-ctx.Done():
		return uuid.Nil, storage.ErrCanceledByContext
	default:
		s.mu.Lock()
		defer s.mu.Unlock()
		if _, ok := s.mapEvents[event.ID]; ok {
			return event.ID, nil
		}
	}
	return uuid.Nil, storage.ErrEvent404
}

func (s *Storage) ListAll(ctx context.Context) ([]storage.Event, error) {
	select {
	case <-ctx.Done():
		return []storage.Event{}, storage.ErrCanceledByContext
	default:
		result := make([]storage.Event, 0, len(s.mapEvents))
		for _, event := range s.mapEvents {
			result = append(result, event)
		}
		sort.Slice(result, func(i, j int) bool {
			return result[i].StartAt.Before(result[j].StartAt)
		})
		return result, nil
	}
	return []storage.Event{}, storage.ErrEvent404
}

func (s *Storage) DeleteAll(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return storage.ErrCanceledByContext
	default:
		s.mu.Lock()
		defer s.mu.Unlock()
	}
	return nil
}

func (s *Storage) Delete(ctx context.Context, id uuid.UUID) error {
	select {
	case <-ctx.Done():
		return storage.ErrCanceledByContext
	default:
		s.mu.Lock()
		defer s.mu.Unlock()
		if _, ok := s.mapEvents[id]; !ok {
			return storage.ErrEvent404
		}
		delete(s.mapEvents, id)
	}
	return nil
}

func (s *Storage) Update(ctx context.Context, event *storage.Event) error {
	select {
	case <-ctx.Done():
		return storage.ErrCanceledByContext
	default:
		s.mu.Lock()
		defer s.mu.Unlock()
		if _, ok := s.mapEvents[event.ID]; !ok {
			return storage.ErrEvent404
		}
		s.mapEvents[event.ID] = *event
	}
	return nil
}

func (s *Storage) GetEventsPerDay(ctx context.Context, day time.Time) ([]storage.Event, error) {
	eventsPerDay := make([]storage.Event, 0)
	select {
	case <-ctx.Done():
		return nil, storage.ErrCanceledByContext
	default:
		s.mu.Lock()
		defer s.mu.Unlock()
		for _, eventStruct := range s.mapEvents {
			if eventStruct.StartAt.Year() == day.Year() && eventStruct.StartAt.Month() == day.Month() &&
				eventStruct.StartAt.Day() == day.Day() {
				eventsPerDay = append(eventsPerDay, eventStruct)
			}
		}
	}
	return eventsPerDay, nil
}

func (s *Storage) GetEventsPerWeek(ctx context.Context, beginDate time.Time) ([]storage.Event, error) {
	eventsPerWeek := make([]storage.Event, 0)
	select {
	case <-ctx.Done():
		return nil, storage.ErrCanceledByContext
	default:
		s.mu.Lock()
		defer s.mu.Unlock()
		endDay := beginDate.AddDate(0, 0, 7)
		for _, eventStruct := range s.mapEvents {
			if (eventStruct.StartAt.After(beginDate) || eventStruct.StartAt.Equal(beginDate)) &&
				(eventStruct.StartAt.Before(endDay) || eventStruct.StartAt.Equal(endDay)) {
				eventsPerWeek = append(eventsPerWeek, eventStruct)
			}
		}
	}
	return eventsPerWeek, nil
}

func (s *Storage) GetEventsPerMonth(ctx context.Context, beginDate time.Time) ([]storage.Event, error) {
	eventsPerMonth := make([]storage.Event, 0)
	select {
	case <-ctx.Done():
		return nil, storage.ErrCanceledByContext
	default:
		s.mu.Lock()
		defer s.mu.Unlock()
		endDay := beginDate.AddDate(0, 1, 0)
		for _, eventStruct := range s.mapEvents {
			if (eventStruct.StartAt.After(beginDate) || eventStruct.StartAt.Equal(beginDate)) &&
				(eventStruct.StartAt.Before(endDay) || eventStruct.StartAt.Equal(endDay)) {
				eventsPerMonth = append(eventsPerMonth, eventStruct)
			}
		}
	}
	return eventsPerMonth, nil
}

func (s *Storage) String() string {
	return "memory storage"
}

func (s *Storage) ListForScheduler(context.Context, time.Duration, time.Duration) ([]storage.Notification, error) {
	return nil, nil
}
