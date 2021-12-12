package task

import (
	"time"

	"github.com/sk-develop/motivation-management/shared/errors"
)

type UpdatedAt time.Time

func NewUpdatedAt(value time.Time) (*UpdatedAt, error) {
	now := time.Now()
	if !now.Before(value) {
		return nil, errors.ValidationError("Past date cannot be registered")
	}
	updatedAt := UpdatedAt(value)

	return &updatedAt, nil
}
