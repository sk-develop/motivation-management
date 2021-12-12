package task

import (
	"errors"
	"time"
)

type DueDate time.Time

func NewDueDate(value time.Time) (*DueDate, error) {
	now := time.Now()
	if !now.Before(value) {
		return nil, errors.New("Past date cannot be registered")
	}
	dueDate := DueDate(value)

	return &dueDate, nil
}
