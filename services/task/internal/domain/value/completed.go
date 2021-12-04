package value

type Completed bool

func (tv *taskValue) NewCompleted(value bool) (*Completed, error) {
	completed := Completed(value)

	return &completed, nil
}
