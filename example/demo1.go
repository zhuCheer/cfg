package main

import (
	"fmt"
	"github.com/zhuCheer/cfg"
)

func main () {

	fmt.Println("cfg start")

	cfg.InitConfFile("./config.toml")
	value:=cfg.GetInt("database.connection_max")

	fmt.Println(value)

}