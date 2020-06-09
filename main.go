package main

import (
	"flag"
	"fmt"

	"github.com/iAviPro/goRestSetup/config"
	"github.com/iAviPro/goRestSetup/server"
)

func main() {
	var env string
	//Params: address of string name, flagName, default value, message
	flag.StringVar(&env, "e", "prod", "Define Environment -e [local | prod]. Default Environment is prod")
	flag.Parse()
	fmt.Printf(":Execution Environment => %s \n ", env)
	config.SetEnvProps(env, "")
	conf := config.GetProps()
	fmt.Println(conf.DbHost)
	server.Start()
}
