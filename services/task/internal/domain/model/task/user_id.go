package task

import "errors"

type UserID string

func NewUserID(value string) (*UserID, error) {
	if len(value) == 0 {
		return nil, errors.New("User ID is required")
	}

	if len(value) > maxLength {
		return nil, errors.New("User ID must be less than 255 characters")
	}
	userID := UserID(value)

	return &userID, nil
}
