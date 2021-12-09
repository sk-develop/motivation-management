package request

import (
	"time"

	"github.com/sk-develop/motivation-management/services/task/internal/domain/model/task"
)

type UpdateTaskReq struct {
	userID  task.UserID
	name    string
	dueDate time.Time
}
