package main

import (
	"encoding/json"
	"fmt"
	"github.com/rigoncs/TodoList/conf"
)

func main() {
	conf.InitConfig()
	tmp, _ := json.Marshal(conf.Conf)
	fmt.Println(string(tmp))
}
