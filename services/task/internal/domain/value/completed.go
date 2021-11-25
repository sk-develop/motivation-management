package value

type Completed bool

func (taskValue *taskValue) NewCompleted(value bool) (*Completed, error) {
	completed := Completed(value)

	return &completed, nil
}
