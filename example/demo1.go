package main

import (
	"fmt"
	"github.com/zhuCheer/cfg"
)

func main() {

	fmt.Println("cfg start")

	cfgHandler, _ := cfg.New("./config.toml")
	value := cfgHandler.GetInt("database.connection_max")

	fmt.Println(value)

}
