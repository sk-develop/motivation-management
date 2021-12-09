package task

import (
	"github.com/sk-develop/motivation-management/shared/errors"
)

type UserID string

func NewUserID(value string) (*UserID, error) {
	if len(value) == 0 {
		return nil, errors.NewValidationError("User ID is required")
	}

	if len(value) > lengthLimit {
		return nil, errors.NewValidationError("User ID must be less than 255 characters")
	}
	userID := UserID(value)

	return &userID, nil
}
