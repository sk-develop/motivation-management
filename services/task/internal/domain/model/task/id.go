package task

import (
	"github.com/sk-develop/motivation-management/shared/errors"
)

type ID int

func NewID(i int) (*ID, error) {
	if i < minLength {
		return nil, errors.ValidationError("ID must be an integer greater than or equal to 1")
	}
	id := ID(i)

	return &id, nil
}
