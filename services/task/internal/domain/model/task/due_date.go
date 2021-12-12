package task

import (
	"time"

	"github.com/sk-develop/motivation-management/shared/errors"
)

type DueDate time.Time

func NewDueDate(value time.Time) (*DueDate, error) {
	now := time.Now()
	if !now.Before(value) {
		return nil, errors.ValidationError("Past date cannot be registered")
	}
	dueDate := DueDate(value)

	return &dueDate, nil
}
