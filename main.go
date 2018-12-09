package main

import (
	"./WebsocketHandler"
	"./queue"
	"./consumer"
	"flag"

)


func main() {

	//新建队列
	go queue.StartQueue()
	go consumer.GetFinalQueueFromChannel()
	go newPost()
	newWs()
}
func newWs()  {
	//新建websocket服务端
	addr  := flag.String("a", "0.0.0.0:1234", "WebsocketHandler server listen address")
	flag.Parse()
	wsServer := &WebsocketHandler.Server{
		ListenAddr : *addr,
	}
	//开始监听
	wsServer.Start()
}
func newPost()  {
		addrPost  := flag.String("post", "0.0.0.0:8000", "WebsocketHandler server listen address")
		PostServer := &WebsocketHandler.Server{
			ListenAddr : *addrPost,
		}
		PostServer.StartPostListener()

}