package main

import (
	"./queue"
	"./websocket"
	"fmt"
	"flag"
)



func main() {

	//新建队列
	q := queue.Queue{1}
	q.Push(1)
	q.Push(3)
	//fmt.Println(q.Pop())
	//fmt.Println(q.Pop())
	//fmt.Println(q.IsEmpty())
	//fmt.Println(q.Pop())
	//fmt.Println(q.IsEmpty())
	fmt.Println(q.FindThingsAlreadyInside(1))

	//新建websocket服务端
	addr  := flag.String("a", "127.0.0.1:1234", "websocket server listen address")
	flag.Parse()
	wsServer := &websocket.WSServer{
		ListenAddr : *addr,
	}
	//开始监听
	wsServer.Start()

	}
