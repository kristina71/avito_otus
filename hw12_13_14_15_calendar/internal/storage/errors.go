package storage

import "errors"

var (
	ErrEventExists = errors.New("event already exists")
	ErrEvent404    = errors.New("event not found")
	ErrNoUserID    = errors.New("no user id")
	ErrEmptyTitle  = errors.New("empty title")
	ErrStartInPast = errors.New("start in past")
	ErrDateBusy    = errors.New("err date busy")
	ErrNoRows      = errors.New("err no rows")
)
