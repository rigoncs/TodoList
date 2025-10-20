package main

import (
	"github.com/rigoncs/TodoList/conf"
	"github.com/rigoncs/TodoList/infrastructure/common/log"
	"github.com/rigoncs/TodoList/infrastructure/container"
	"github.com/rigoncs/TodoList/infrastructure/persistence/dbs"
	"github.com/rigoncs/TodoList/interfaces/adapter/initialize"
)

func main() {
	loadingInfra()
	r := initialize.NewRouter()
	_ = r.Run(conf.Conf.Server.Port)
}

func loadingInfra() {
	conf.InitConfig()
	log.InitLog()
	dbs.MySQLInit()

	container.LoadingDomain()
}
