package task

import (
	"time"

	"github.com/sk-develop/motivation-management/shared/errors"
)

type CreatedAt time.Time

func NewCreatedAt(value time.Time) (*CreatedAt, error) {
	now := time.Now()
	if !now.Before(value) {
		return nil, errors.ValidationError("Past date cannot be registered")
	}
	createdAt := CreatedAt(value)

	return &createdAt, nil
}
