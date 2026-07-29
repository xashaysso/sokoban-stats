package services

import "errors"

var (
	ErrInvalidLevelID  = errors.New("invalid level id")
	ErrInvalidUsername = errors.New("invalid username")
	ErrInvalidSteps    = errors.New("can't pass level with 0 steps")
)
