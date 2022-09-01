package internalhttp

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/app"
)

type Server struct {
	server *http.Server
	router *mux.Router
	logger Logger
	app    *app.App
}

type Logger interface {
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Debug(msg string)
}

/*
type Application interface { // TODO
}
*/

/*func HelloWorld(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := r.Body.Close(); err != nil {
			zap.L().Error("unable to close response body for request", zap.Error(err))
		}
	}()
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("Hello World!")); err != nil {
		zap.L().Error("http write error", zap.Error(err))
	}
}
*/

func NewServer(logger Logger, app *app.App) *Server {
	service := &Server{logger: logger, app: app}
	service.router = mux.NewRouter()
	service.router.Use(loggingMiddleware(service.logger))
	service.router.HandleFunc("/calendar/add", service.createEvent).Methods("POST")
	service.router.HandleFunc("/calendar/update", service.updateEvent).Methods("PUT")
	service.router.HandleFunc("/calendar/delete/{eventId}", service.deleteEvent).Methods("DELETE")
	service.router.HandleFunc(
		"/calendar/find/{period:[a-zA-Z]+}/{year:[0-9]{4}}/{month:[0-9]{2}}/{day:[0-9]{2}}",
		service.getEventsPerDay,
	).Methods("GET")
	service.router.HandleFunc(
		"/calendar/find/{period:[a-zA-Z]+}/{year:[0-9]{4}}/{month:[0-9]{2}}/{day:[0-9]{2}}",
		service.getEventsPerWeek,
	).Methods("GET")
	service.router.HandleFunc(
		"/calendar/find/{period:[a-zA-Z]+}/{year:[0-9]{4}}/{month:[0-9]{2}}/{day:[0-9]{2}}",
		service.getEventsPerMonth,
	).Methods("GET")

	return service
}

func (s *Server) Start(ctx context.Context, addr string) error {
	s.logger.Info("HTTP server starting..." + addr)

	s.server = &http.Server{
		Handler:      s.router,
		Addr:         addr,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	errChan := make(chan error)

	go func() {
		if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errChan <- err
		}
	}()

	select {
	case <-ctx.Done():
	case err := <-errChan:
		s.logger.Error("Failed to start http server")
		return err
	}
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.logger.Info("HTTP server stopping..." + s.server.Addr)
	err := s.server.Shutdown(ctx)
	s.logger.Info("HTTP server stopped")
	return err
}
