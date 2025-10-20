package persistence

import (
	tRepo "github.com/rigoncs/TodoList/domain/task/repository"
	uRepo "github.com/rigoncs/TodoList/domain/user/repository"
	"github.com/rigoncs/TodoList/infrastructure/persistence/task"
	"github.com/rigoncs/TodoList/infrastructure/persistence/user"
	"gorm.io/gorm"
)

type Repositories struct {
	User uRepo.User
	Task tRepo.Task
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		User: user.NewRepositoryImpl(db),
		Task: task.NewRepository(db),
	}
}
