package value

import (
	"time"

	"github.com/sk-develop/motivation-management/shared/errors"
	"github.com/sk-develop/motivation-management/shared/logger"
)

type UpdatedAt time.Time

func (tv *taskValue) NewUpdatedAt(value time.Time) (*UpdatedAt, error) {
	now := time.Now()
	if !now.Before(value) {
		err := errors.NewValidationError("Past date cannot be registered")
		logger.Warn(err)

		return nil, err
	}
	updatedAt := UpdatedAt(value)

	return &updatedAt, nil
}
