package main

import (
	"fmt"
	"github.com/lidongyooo/lightsocks-R/pkg/config"
	"github.com/lidongyooo/lightsocks-R/pkg/password"
	"github.com/phayes/freeport"
	"log"
	"os"
	"strconv"
)

func main() {
	log.SetFlags(log.Lshortfile)

	//优先从环境变量中读取监听端口
	port, err := strconv.Atoi(os.Getenv("LIGHTSOCKS_SERVER_PORT"))

	if err != nil {
		port, err = freeport.GetFreePort()
	}

	if err != nil {
		port = 7448
	}

	_config := &config.Config{
		ListenAddr: fmt.Sprint(port),
		Password: password.RandPassword(),
	}

	_config.ReadConfig()
	_config.SaveConfig()
}
