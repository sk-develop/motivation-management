package value

import (
	"github.com/sk-develop/motivation-management/shared/errors"
	"github.com/sk-develop/motivation-management/shared/logger"
)

type UserID string

func (taskValue *taskValue) NewUserID(value string) (*UserID, error) {
	if len(value) == 0 {
		err := errors.NewValidationError("User ID is required")
		logger.Warn(err)

		return nil, err
	}

	if len(value) > lengthLimit {
		err := errors.NewValidationError("User ID must be less than 255 characters")
		logger.Warn(err)

		return nil, err
	}
	userID := UserID(value)

	return &userID, nil
}
