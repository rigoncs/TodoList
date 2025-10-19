package task

import (
	"github.com/rigoncs/TodoList/domain/task/entity"
	"github.com/rigoncs/TodoList/infrastructure/persistence/user"
)

func Entity2PO(task *entity.Task) *Task {
	return &Task{
		Uid:       task.Uid,
		Title:     task.Title,
		Status:    task.Status,
		Content:   task.Content,
		StartTime: task.StartTime,
		EndTime:   task.EndTime,
	}
}

func PO2Entity(task *Task, u *user.User) *entity.Task {
	return &entity.Task{
		Uid:       task.Uid,
		UserName:  u.UserName,
		Title:     task.Title,
		Status:    task.Status,
		Content:   task.Content,
		StartTime: task.StartTime,
		EndTime:   task.EndTime,
	}
}
