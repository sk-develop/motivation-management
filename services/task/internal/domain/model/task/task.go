package task

type Task struct {
	ID        ID
	UserID    UserID
	Name      Name
	Completed Completed
	DueDate   DueDate
	CreatedAt CreatedAt
	UpdatedAt UpdatedAt
}

const (
	maxLength = 255
	minLength = 1
)

type Tasks []Task

func New(userID UserID, name Name, dueDate DueDate) *Task {
	return &Task{UserID: userID, Name: name, DueDate: dueDate}
}
