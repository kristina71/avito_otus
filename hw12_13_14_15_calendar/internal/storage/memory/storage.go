package memorystorage

import (
	"context"
	"sort"
	"sync"
	"time"

	"github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/logger"
	"github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/storage"
)

type Storage struct {
	mu     sync.RWMutex
	lastID int
	data   data
	logger logger.Logger
}

func New() *Storage {
	return &Storage{data: make(map[int]storage.Event)}
}

type Repository struct {
	repo storage.Storage
}

func New1(repo storage.Storage) *Repository {
	return &Repository{repo: repo}
}

type data map[int]storage.Event

func (s *Storage) Connect(_ context.Context) error {
	s.logger.Info("connected to memory storage")
	return nil
}

func (s *Storage) Close() error {
	s.logger.Info("closing memory storage")
	return nil
}

func (s *Storage) Create(_ context.Context, event storage.Event) (storage.Event, error) {
	s.logger.Info("create a new event")

	s.mu.Lock()
	defer s.mu.Unlock()

	id := s.newID()
	event.ID = id
	s.data[id] = storage.Event{
		ID:          id,
		Title:       event.Title,
		StartAt:     event.StartAt,
		EndAt:       event.EndAt,
		Description: event.Description,
		UserID:      event.UserID,
		RemindAt:    event.RemindAt,
	}
	return event, nil
<<<<<<< HEAD
}

func (s *Storage) Get(_ context.Context, id int) (storage.Event, error) {
	s.logger.Info("get the event")

	s.mu.Lock()
	defer s.mu.Unlock()

	event, ok := s.data[id]
	if !ok {
		return storage.Event{}, storage.ErrEvent404
	}
	return event, nil
}

func (s *Storage) Update(_ context.Context, id int, change storage.Event) error {
	s.logger.Info("update the event")

	s.mu.Lock()
	defer s.mu.Unlock()

	event, ok := s.data[id]
	if !ok {
		return storage.ErrEvent404
	}

	event.Title = change.Title
	event.StartAt = change.StartAt
	event.EndAt = change.EndAt
	event.Description = change.Description
	event.RemindAt = change.RemindAt
	s.data[id] = event

	return nil
}

func (s *Storage) Delete(_ context.Context, id int) error {
	s.logger.Info("delete the event")

	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.data, id)
	return nil
=======
>>>>>>> HW-12 Implement db logic
}
func (s *Storage) Get(_ context.Context, id int) (storage.Event, error) {
	s.logger.Info("get the event")

	s.mu.Lock()
	defer s.mu.Unlock()

	event, ok := s.data[id]
	if !ok {
		return storage.Event{}, storage.ErrEvent404
	}
	return event, nil
}

func (s *Storage) Update(_ context.Context, id int, change storage.Event) error {
	s.logger.Info("update the event")

	s.mu.Lock()
	defer s.mu.Unlock()

	event, ok := s.data[id]
	if !ok {
		return storage.ErrEvent404
	}

	event.Title = change.Title
	event.StartAt = change.StartAt
	event.EndAt = change.EndAt
	event.Description = change.Description
	event.RemindAt = change.RemindAt
	s.data[id] = event

	return nil
}

func (s *Storage) Delete(_ context.Context, id int) error {
	s.logger.Info("delete the event")

	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.data, id)
	return nil
}

func (s *Storage) DeleteAll(_ context.Context) error {
	s.logger.Info("delete all events")

<<<<<<< HEAD
func (s *Storage) DeleteAll(_ context.Context) error {
	s.logger.Info("delete all events")

=======
>>>>>>> HW-12 Implement db logic
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data = make(data)
	return nil
}

func (s *Storage) ListAll(_ context.Context) ([]storage.Event, error) {
	s.logger.Info("get all events")

	s.mu.Lock()
	defer s.mu.Unlock()

	result := make([]storage.Event, 0, len(s.data))
	for _, event := range s.data {
		result = append(result, event)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].StartAt.Before(result[j].StartAt)
	})
	return result, nil
}

func (s *Storage) ListDay(_ context.Context, date time.Time) ([]storage.Event, error) {
	s.logger.Info("get events by day")

	s.mu.Lock()
	defer s.mu.Unlock()

	var result []storage.Event
	year, month, day := date.Date()
	for _, event := range s.data {
		eventYear, eventMonth, eventDay := event.StartAt.Date()
		if eventYear == year && eventMonth == month && eventDay == day {
			result = append(result, event)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].StartAt.Before(result[j].StartAt)
	})
	return result, nil
}

func (s *Storage) ListWeek(_ context.Context, date time.Time) ([]storage.Event, error) {
	s.logger.Info("get events by week")

	s.mu.Lock()
	defer s.mu.Unlock()

	var result []storage.Event
	year, week := date.ISOWeek()
	for _, event := range s.data {
		eventYear, eventWeek := event.StartAt.ISOWeek()
		if eventYear == year && eventWeek == week {
			result = append(result, event)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].StartAt.Before(result[j].StartAt)
	})
	return result, nil
}

func (s *Storage) ListMonth(_ context.Context, date time.Time) ([]storage.Event, error) {
	s.logger.Info("get events by month")

	s.mu.Lock()
	defer s.mu.Unlock()

	var result []storage.Event
	year, month, _ := date.Date()
	for _, event := range s.data {
		eventYear, eventMonth, _ := event.StartAt.Date()
		if eventYear == year && eventMonth == month {
			result = append(result, event)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].StartAt.Before(result[j].StartAt)
	})
	return result, nil
}

func (s *Storage) IsTimeBusy(_ context.Context, userID int, start, stop time.Time, excludeID int) (bool, error) {
	s.logger.Info("is time to busy")

	s.mu.Lock()
	defer s.mu.Unlock()

	/*for _, event := range s.data {
		if event.UserID == userID && event.ID != string(excludeID) && event.StartAt.Before(stop) && event.EndAt.After(start) {
			return true, nil
		}
	}
	*/

	return false, nil
}

func (s *Storage) newID() int {
	s.lastID++
	return s.lastID
}

func (s *Storage) String() string {
	return "memory storage"
}
