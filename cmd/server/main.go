package main

import (
	"fmt"
	"github.com/lidongyooo/lightsocks-R/pkg/config"
	"github.com/lidongyooo/lightsocks-R/pkg/password"
	"github.com/lidongyooo/lightsocks-R/pkg/server"
	"github.com/phayes/freeport"
	"log"
	"net"
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
		ListenAddr: fmt.Sprintf(":%d", port),
		Password: password.RandPassword(),
	}

	_config.ReadConfig()
	_config.SaveConfig()

	_server, err :=  server.New(_config.Password, _config.ListenAddr)
	if err != nil {
		log.Fatalln(err)
	}

	log.Fatalln(_server.Listen(func(listenAddr net.Addr) {
		log.Println(fmt.Sprintf(`
lightsocks-server: 启动成功，配置如下：
服务监听地址：
%s
密码：
%s`, listenAddr, _config.Password))
	}))
}

