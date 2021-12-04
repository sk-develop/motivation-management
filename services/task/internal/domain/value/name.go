package value

import (
	"github.com/sk-develop/motivation-management/shared/errors"
	"github.com/sk-develop/motivation-management/shared/logger"
)

type Name string

func (tv *taskValue) NewName(value string) (*Name, error) {
	if len(value) == 0 {
		err := errors.NewValidationError("Name is required")
		logger.Warn(err)

		return nil, err
	}

	if len(value) >= lengthLimit {
		err := errors.NewValidationError("Name must be 255 characters or less")
		logger.Warn(err)

		return nil, err
	}
	title := Name(value)

	return &title, nil
}
