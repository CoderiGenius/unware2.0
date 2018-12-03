package main

import (
	"./WebsocketHandler"
	"./queue"
	"flag"
)


func main() {


	//新建队列
	go queue.StartQueue()

	//新建websocket服务端
	addr  := flag.String("a", "127.0.0.1:1234", "WebsocketHandler server listen address")
	flag.Parse()
	wsServer := &WebsocketHandler.Server{
		ListenAddr : *addr,
	}
	//开始监听
	wsServer.Start()

	}
