package task

import (
	"time"

	"github.com/sk-develop/motivation-management/shared/errors"
	"github.com/sk-develop/motivation-management/shared/logger"
)

type CreatedAt time.Time

func NewCreatedAt(value time.Time) (*CreatedAt, error) {
	now := time.Now()
	if !now.Before(value) {
		err := errors.NewValidationError("Past date cannot be registered")
		logger.Warn(err)

		return nil, err
	}
	createdAt := CreatedAt(value)

	return &createdAt, nil
}
