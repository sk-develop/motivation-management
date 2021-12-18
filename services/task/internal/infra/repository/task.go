package repository

import (
	"github.com/sk-develop/motivation-management/services/task/internal/domain/model/task"
	"github.com/sk-develop/motivation-management/services/task/internal/domain/repository"
	"github.com/sk-develop/motivation-management/shared/logger"

	"gorm.io/gorm"
)

type TaskRepository struct {
	conn *gorm.DB
}

func NewTaskRepository(conn *gorm.DB) repository.TaskRepository {
	return &TaskRepository{conn: conn}
}

func (tr *TaskRepository) ReadAll(userId task.UserID) (*task.Tasks, error) {
	var foundTask *task.Tasks

	if err := tr.conn.Where("user_id = ?", userId).Find(&foundTask).Error; err != nil {
		return nil, err
	}

	return foundTask, nil
}

func (tr *TaskRepository) Create(userID task.UserID, name task.Name, dueDate task.DueDate) (*task.Task, error) {
	createdTask := task.New(userID, name, dueDate)

	if err := tr.conn.Create(createdTask).Error; err != nil {
		return nil, err
	}

	return createdTask, nil
}

func (tr *TaskRepository) Update(id task.ID, name task.Name, dueDate task.DueDate) (*task.Task, error) {
	var task *task.Task
	if err := tr.conn.Take(&task, id).Error; err != nil {
		return nil, err
	}

	task.Name = name
	task.DueDate = dueDate
	if err := tr.conn.Save(task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (tr *TaskRepository) Delete(ids []task.ID) error {
	var tasks *task.Task
	if err := tr.conn.Delete(&tasks, ids).Error; err != nil {
		return err
	}

	return nil
}

func (tr *TaskRepository) SwitchCompleted(id task.ID) (*task.Task, error) {
	var task task.Task
	if err := tr.conn.Raw("UPDATE tasks SET completed = NOT completed WHERE id = ?", id).Scan(&task).Error; err != nil {
		return nil, err
	}
	if err := tr.conn.Where("id = ?", id).Find(&task).Error; err != nil {
		return nil, err
	}

	return &task, nil
}
