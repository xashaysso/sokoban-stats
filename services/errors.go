package services

import "errors"

var (
	ErrInvalidLevelID  = errors.New("invalid level id")
	ErrInvalidUsername = errors.New("invalid username")
	ErrInvalidSteps    = errors.New("invalid steps value")
	ErrInvalidTime     = errors.New("invalid time value")
)
