package task

import (
	"errors"
	"time"
)

type CreatedAt time.Time

func NewCreatedAt(value time.Time) (*CreatedAt, error) {
	now := time.Now()
	if !now.Before(value) {
		return nil, errors.New("Past date cannot be registered")
	}
	createdAt := CreatedAt(value)

	return &createdAt, nil
}
