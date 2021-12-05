package value

import "time"

type taskValue struct{}

func NewTaskValue() TaskValue {
	return &taskValue{}
}

const minLength = 1
const lengthLimit = 255

type TaskValue interface {
	NewID(value int) (*ID, error)
	NewUserID(value string) (*UserID, error)
	NewName(value string) (*Name, error)
	NewDueDate(value time.Time) (*DueDate, error)
}
