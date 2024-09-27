package models

import "errors"

var (
	ErrEmailTaken      = errors.New("models: email address is already in use")
	ErrInvalidPassword = errors.New("models: invalid password")
	ErrNotFound        = errors.New("models: resource not found")
)
