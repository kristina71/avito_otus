package internalhttp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/storage"
)

func (s *Server) createEvent(w http.ResponseWriter, r *http.Request) {
	ev := storage.Event{}
	err := json.NewDecoder(r.Body).Decode(&ev)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error to get request body: %v", err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.app.Create(r.Context(), &ev)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error to create event: %v", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *Server) updateEvent(w http.ResponseWriter, r *http.Request) {
	ev := storage.Event{}
	err := json.NewDecoder(r.Body).Decode(&ev)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error to get request body: %v", err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.app.Update(r.Context(), &ev)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error to update event: %v", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *Server) deleteEvent(w http.ResponseWriter, r *http.Request) {
	type id struct {
		ID uuid.UUID `json:"id"`
	}
	var idEv id
	err := json.NewDecoder(r.Body).Decode(&idEv)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error to get request body: %v", err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.app.Delete(r.Context(), idEv.ID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error to delete event: %v", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *Server) getEvents(w http.ResponseWriter, r *http.Request) {
	events, err := s.app.ListAll(r.Context())
	s.writeHeader(w, err, events)
}

func (s *Server) deleteAllEvents(w http.ResponseWriter, r *http.Request) {
	err := s.app.DeleteAll(r.Context())
	if err != nil {
		s.logger.Error(fmt.Sprintf("error to delete events: %v", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *Server) getEventsPerDay(w http.ResponseWriter, r *http.Request) {
	var day time.Time
	err := json.NewDecoder(r.Body).Decode(&day)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error to get request body: %v", err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	events, err := s.app.GetEventsPerDay(r.Context(), day)
	s.writeHeader(w, err, events)
}

func (s *Server) getEventsPerWeek(w http.ResponseWriter, r *http.Request) {
	var day time.Time
	err := json.NewDecoder(r.Body).Decode(&day)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error to get request body: %v", err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	events, err := s.app.GetEventsPerWeek(r.Context(), day)
	s.writeHeader(w, err, events)
}

func (s *Server) getEventsPerMonth(w http.ResponseWriter, r *http.Request) {
	var day time.Time
	err := json.NewDecoder(r.Body).Decode(&day)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error to get request body: %v", err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	events, err := s.app.GetEventsPerMonth(r.Context(), day)
	s.writeHeader(w, err, events)
}

func (s *Server) writeHeader(w http.ResponseWriter, err error, events []storage.Event) {
	if err != nil {
		s.logger.Error(fmt.Sprintf("error to get list of events per month: %v", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(events)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error to get list of events per month: %v", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
