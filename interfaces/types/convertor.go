package types

import (
	te "github.com/rigoncs/TodoList/domain/task/entity"
	"github.com/rigoncs/TodoList/domain/user/entity"
)

func UserReq2Entity(user *UserReq) *entity.User {
	return &entity.User{
		Username: user.UserName,
		Password: user.Password,
	}
}

func Entity2TaskResp(task *te.Task) *TaskResp {
	return &TaskResp{
		ID:        task.Id,
		Title:     task.Title,
		Content:   task.Content,
		Status:    task.Status,
		StartTime: task.StartTime,
		EndTime:   task.EndTime,
		View:      0,
	}
}
