package container

import (
	taskApp "github.com/rigoncs/TodoList/application/task"
	userApp "github.com/rigoncs/TodoList/application/user"
	taskSvc "github.com/rigoncs/TodoList/domain/task/service"
	userSvc "github.com/rigoncs/TodoList/domain/user/service"
	"github.com/rigoncs/TodoList/infrastructure/auth"
	"github.com/rigoncs/TodoList/infrastructure/encrypt"
	"github.com/rigoncs/TodoList/infrastructure/persistence"
	"github.com/rigoncs/TodoList/infrastructure/persistence/dbs"
)

func LoadingDomain() {
	repos := persistence.NewRepositories(dbs.DB)
	jwtService := auth.NewJWTTokenService()
	pwdEncryptService := encrypt.NewPwdEncryptService()

	userDomain := userSvc.NewUserDomainImpl(repos.User, pwdEncryptService)
	userApp.GetServiceImpl(userDomain, jwtService)

	taskDomain := taskSvc.NewTaskDomainImpl(repos.Task)
	taskApp.GetServiceImpl(taskDomain)
}
