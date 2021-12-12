package task

import (
	"errors"
	"time"
)

type UpdatedAt time.Time

func NewUpdatedAt(value time.Time) (*UpdatedAt, error) {
	now := time.Now()
	if !now.Before(value) {
		return nil, errors.New("Past date cannot be registered")
	}
	updatedAt := UpdatedAt(value)

	return &updatedAt, nil
}
