package WebsocketHandler

import (
	"golang.org/x/net/websocket"
	"fmt"
	"net/http"
	"encoding/json"

	"../queue"
)

type WSServer struct {
	ListenAddr string
}
type RequstJSON struct {
	From string
	Content string
}


type Server struct {
	// Config is a WebSocket configuration for new WebSocket connection.
	Config websocket.Config

	ListenAddr string
	// Handshake is an optional function in WebSocket handshake.
	// For example, you can check, or don't check Origin header.
	// Another example, you can select config.Protocol.
	Handshake func(*websocket.Config, *http.Request) error

	// Handler handles a WebSocket connection.
	//Handler func(conn *WebsocketHandler.Conn)
}



//接受握手消息并判断客户端类型
//若为 producer 则为人脸识别客户端 若为 consumer 则为浏览器
func (this *Server) Handler(conn *websocket.Conn) {
	fmt.Printf("a new ws conn: %s->%s\n", conn.RemoteAddr().String(), conn.LocalAddr().String())
	var err error

	var reply RequstJSON
	var rep string

for {
	rep = GetMessage(conn)
	if err != nil {
		panic(err)
	}

	//fmt.Println(rep)
	err2 := json.Unmarshal([]byte(rep), &reply)
	if err2 != nil {
		fmt.Println("json erro")
	}
	if reply.From == "producer" {

		queue.QueueChannel <- reply.Content
		SendMessage(conn, "accepted")

	} else if reply.From == "comsumer" {

	}
}
}


func SendMessage(conn *websocket.Conn, message string) bool {
	//fmt.Println("SENDING MESSAGE TO: ", conn.RemoteAddr().String(), " MESSAGE：", message)
	 con := conn
	var err error
	err = websocket.Message.Send(con, message)
	if err != nil {
		panic(err)
		return false
	}
	return true
}

func GetMessage(conn *websocket.Conn) string {

	var err error
	var reply string
	err = websocket.Message.Receive(conn, &reply)
	if err != nil {
		panic(err)
		return reply
	}
	fmt.Println("GET MESSAGE FROM: ", conn.RemoteAddr().String(), " MESSAGE：", reply)
	return reply
}
func (this *Server) Start() (error) {
	//配置handle
	http.Handle("/ws", websocket.Handler(this.Handler))

	fmt.Println("begin to listen")
	//开始启用handler
	err := http.ListenAndServe(this.ListenAddr, nil)

	if err != nil {
		fmt.Println("ListenAndServe:", err)
		return err
	}
	fmt.Println("start end")
	return nil
}

//func sendToChannel(queueChannel chan string)  {
//
//	<- channel
//}