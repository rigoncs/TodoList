package service

import (
	"context"
	"github.com/rigoncs/TodoList/domain/task/entity"
	"github.com/rigoncs/TodoList/domain/task/repository"
	"github.com/rigoncs/TodoList/interfaces/types"
)

type TaskDomain interface {
	CreateTask(ctx context.Context, in *entity.Task) (*entity.Task, error)
	FindTaskByTid(ctx context.Context, taskId, userId uint) (*entity.Task, error)
	ListTaskByUid(ctx context.Context, userId uint, p types.Pagination) ([]*entity.Task, int64, error)
	UpdateTask(ctx context.Context, in *entity.Task) error
	SearchTask(ctx context.Context, userId uint, keyword string, p types.Pagination) ([]*entity.Task, int64, error)
	DeleteTask(ctx context.Context, uid, tid uint) error
}

type TaskDomainImpl struct {
	repo repository.Task
}

func NewTaskDomainImpl(repo repository.Task) TaskDomain {
	return &TaskDomainImpl{
		repo: repo,
	}
}

func (t *TaskDomainImpl) CreateTask(ctx context.Context, in *entity.Task) (*entity.Task, error) {
	return t.repo.CreateTask(ctx, in)
}

func (t *TaskDomainImpl) FindTaskByTid(ctx context.Context, taskId, userId uint) (*entity.Task, error) {
	return t.repo.FindTaskByTid(ctx, taskId, userId)
}

func (t *TaskDomainImpl) ListTaskByUid(ctx context.Context, userId uint, p types.Pagination) ([]*entity.Task, int64, error) {
	return t.repo.ListTaskByUid(ctx, userId, p)
}

func (t *TaskDomainImpl) UpdateTask(ctx context.Context, in *entity.Task) error {
	return t.repo.UpdateTask(ctx, in)
}

func (t *TaskDomainImpl) SearchTask(ctx context.Context, userId uint, keyword string, p types.Pagination) ([]*entity.Task, int64, error) {
	return t.repo.SearchTask(ctx, userId, keyword, p)
}

func (t *TaskDomainImpl) DeleteTask(ctx context.Context, uid, tid uint) error {
	return t.repo.DeleteTask(ctx, uid, tid)
}
