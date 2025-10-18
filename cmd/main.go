package main

import (
	"github.com/rigoncs/TodoList/conf"
	"github.com/rigoncs/TodoList/interfaces/adapter/initialize"
)

func main() {
	conf.InitConfig()
	r := initialize.NewRouter()
	_ = r.Run(conf.Conf.Server.Port)
}
