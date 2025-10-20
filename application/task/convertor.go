package task

import (
	"github.com/rigoncs/TodoList/domain/task/entity"
	"github.com/rigoncs/TodoList/interfaces/types"
	"time"
)

func ListResponse(list []*entity.Task, count int64) types.List[*entity.Task] {
	return types.List[*entity.Task]{
		Items: list,
		Count: count,
	}
}

func UpdateReq2TaskEntity(tid, uid uint, username string, req *types.UpdateTaskReq) *entity.Task {
	return &entity.Task{
		Id:        tid,
		Uid:       uid,
		UserName:  username,
		Title:     req.Title,
		Content:   req.Content,
		Status:    req.Status,
		UpdatedAt: time.Now(),
	}
}
