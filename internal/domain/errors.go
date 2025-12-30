package domain

import (
	"errors"
)

var (
	ErrShowIDValidate      = errors.New("show id is required")
	ErrSeatNumbersValidate = errors.New("seats numbers is required")

	ErrSeatUnavailable = errors.New("seats unavailable")
	ErrShowNotFound    = errors.New("show not found")
	ErrDeadlock        = errors.New("deadlock detected")
)
