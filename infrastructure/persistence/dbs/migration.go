package dbs

import (
	"github.com/rigoncs/TodoList/infrastructure/persistence/task"
	"github.com/rigoncs/TodoList/infrastructure/persistence/user"
)

func migration() {
	err := DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&user.User{},
			&task.Task{})
	if err != nil {
		return
	}
}
