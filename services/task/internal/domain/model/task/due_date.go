package task

import (
	"time"

	"github.com/sk-develop/motivation-management/shared/errors"
	"github.com/sk-develop/motivation-management/shared/logger"
)

type DueDate time.Time

func NewDueDate(value time.Time) (*DueDate, error) {
	now := time.Now()
	if !now.Before(value) {
		err := errors.NewValidationError("Past date cannot be registered")
		logger.Warn(err)

		return nil, err
	}
	dueDate := DueDate(value)

	return &dueDate, nil
}
