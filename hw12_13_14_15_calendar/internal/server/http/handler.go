package internalhttp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/storage"
)

func (s *Server) createEvent(w http.ResponseWriter, r *http.Request) {
	ev := storage.Event{}
	err := json.NewDecoder(r.Body).Decode(&ev)

	/* type event struct {
		id          int       `json:"id"`
		title       string    `json:"title"`
		description string    `json:"description"`
		startAt     time.Time `json:"start_at"`
		endAt       time.Time `json:"end_at"`
		userID      int       `json:"user_id"`
		remindAt    time.Time `json:"remind_at"`
	} */

	if err != nil {
		s.logger.Error(fmt.Sprintf("error to get request body: %v", err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ev, err = s.app.Create(r.Context(), ev.UserID, ev.Title, ev.Description, ev.StartAt, ev.EndAt)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error to create event: %v", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *Server) updateEvent(w http.ResponseWriter, r *http.Request) {
	type id struct {
		id int `json:"id"`
	}
	var idEv id

	ev := storage.Event{}
	err := json.NewDecoder(r.Body).Decode(&ev)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error to get request body: %v", err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.app.Update(r.Context(), idEv.id, ev)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error to update event: %v", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *Server) deleteEvent(w http.ResponseWriter, r *http.Request) {
	type id struct {
		id int `json:"id"`
	}
	var idEv id
	err := json.NewDecoder(r.Body).Decode(&idEv)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error to get request body: %v", err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.app.Delete(r.Context(), idEv.id)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error to update event: %v", err))
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
	events, err := s.app.ListDay(r.Context(), day)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error to get list of events: %v", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(events)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error to get list of events: %v", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *Server) getEventsPerWeek(w http.ResponseWriter, r *http.Request) {
	var day time.Time
	err := json.NewDecoder(r.Body).Decode(&day)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error to get request body: %v", err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	events, err := s.app.ListWeek(r.Context(), day)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error to get list of events: %v", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(events)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error to get list of events: %v", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *Server) getEventsPerMonth(w http.ResponseWriter, r *http.Request) {
	var day time.Time
	err := json.NewDecoder(r.Body).Decode(&day)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error to get request body: %v", err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	events, err := s.app.ListMonth(r.Context(), day)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error to get list of events: %v", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(events)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error to get list of events: %v", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}