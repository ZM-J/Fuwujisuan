package main

import (
	"os"

	"github.com/ZM-J/Fuwujisuan/Homework3/service"
	flag "github.com/spf13/pflag"
)

const (
	// PORT : 缺省的监听端口
	PORT string = "8080"
)

func main() {
	// 首先获得环境变量$PORT
	port := os.Getenv("PORT")
	// 如果没有环境变量$PORT，那么将PORT设置成一个常数
	if len(port) == 0 {
		port = PORT
	}

	// 如果在CLI中运行服务器的时候，指定了-p参数的话，那么将监听端口号设置成-p参数
	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	// 开启服务器
	server := service.NewServer()
	server.Run(":" + port)
}