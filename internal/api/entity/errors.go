package entity

import "errors"

var (
	ErrHolderNotFound       = errors.New("holder not found")
	ErrUpdateScoresConflict = errors.New("update scores conflict")
)
