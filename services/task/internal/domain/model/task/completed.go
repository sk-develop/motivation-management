package task

type Completed bool

func NewCompleted(value bool) (*Completed, error) {
	completed := Completed(value)

	return &completed, nil
}
