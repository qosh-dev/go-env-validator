package main

import (
	"config/config"
	"fmt"
)

func main() {

	configInst := config.Config{}
	configInst.Initialize()

	fmt.Println(configInst.Env)

}
