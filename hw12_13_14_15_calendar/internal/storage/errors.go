package storage

import "errors"

var (
	ErrEventExists       = errors.New("event already exists")
	ErrEvent404          = errors.New("event not found")
	ErrWrongDuration     = errors.New("wrong duration")
	ErrWrongId           = errors.New("wrong id")
	ErrEmptyTitle        = errors.New("empty title")
	ErrDateBusy          = errors.New("err date busy")
	ErrNoRows            = errors.New("err no rows")
	ErrCanceledByContext = errors.New("cancel by context")
)
