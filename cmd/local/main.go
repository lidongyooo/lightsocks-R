package main

import (
	"fmt"
	"github.com/lidongyooo/lightsocks-R/pkg/config"
	"github.com/lidongyooo/lightsocks-R/pkg/local"
	"log"
	"net"
)


func main()  {
	log.SetFlags(log.Lshortfile)

	_config := &config.Config{}

	_config.ReadConfig()
	_config.SaveConfig()

	_local, err := local.New(_config.Password, _config.ListenAddr, _config.RemoteAddr)
	if err != nil {
		log.Fatalln(err)
	}

	log.Fatalln(_local.Listen(func(listenAddr net.Addr) {
		log.Println(fmt.Sprintf(`
lightsocks-local: 启动成功，配置如下：
本地监听地址：
%s
远程服务地址：
%s
密码：
%s`, listenAddr, _config.RemoteAddr, _config.Password))
	}))
}