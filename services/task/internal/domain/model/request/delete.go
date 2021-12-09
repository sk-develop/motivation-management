package request

import "github.com/sk-develop/motivation-management/services/task/internal/domain/model/task"

type DeleteTaskReq struct {
	userID task.UserID
	ids    []task.ID
}
