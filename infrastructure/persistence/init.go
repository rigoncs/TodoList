package persistence

import (
	uRepo "github.com/rigoncs/TodoList/domain/user/repository"
	"github.com/rigoncs/TodoList/infrastructure/persistence/user"
	"gorm.io/gorm"
)

type Repositories struct {
	User uRepo.User
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		User: user.NewRepositoryImpl(db),
	}
}
