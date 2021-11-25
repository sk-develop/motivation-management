package value

import (
	"github.com/sk-develop/motivation-management/shared/errors"
	"github.com/sk-develop/motivation-management/shared/logger"
)

type ID int

func (taskValue *taskValue) NewID(i int) (*ID, error) {
	if i < minLength {
		err := errors.NewValidationError("ID must be an integer greater than or equal to 1")
		logger.Warn(err)

		return nil, err
	}
	id := ID(i)

	return &id, nil
}
