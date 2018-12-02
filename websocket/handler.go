package websocket


import(
"golang.org/x/net/websocket"
"fmt"
"net/http"
)

type WSServer struct {
	ListenAddr string
}

//接受握手消息并判断客户端类型
//若为 producer 则为人脸识别客户端 若为 consumer 则为浏览器
func (this *WSServer)Handler(conn *websocket.Conn){
	fmt.Printf("a new ws conn: %s->%s\n", conn.RemoteAddr().String(), conn.LocalAddr().String())
	var err error
	/*for {
		var reply string
		err = websocket.Message.Receive(conn, &reply)
		if err != nil {
			fmt.Println("receive err:",err.Error())
			break
		}
		fmt.Println("Received from client: " + reply)
		SendMessage(conn,"1")
		if err = websocket.Message.Send(conn, "ok"); err != nil {
			fmt.Println("send err:", err.Error())
			break
		}*/
		var reply string
		err = websocket.Message.Receive(conn,reply)
		if reply=="producer"{

	}
}

func SendMessage(conn *websocket.Conn,message string) bool{
	fmt.Println("SENDING MESSAGE TO: ",conn.RemoteAddr().String()," MESSAGE：",message)
	var err error
	err = websocket.Message.Send(conn,message)
	if err  != nil {
		panic(err)
		return false
	}
	return true
}
func (this *WSServer)Start()(error){
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

/*func main(){
	addr  := flag.String("a", "127.0.0.1:1234", "websocket server listen address")
	flag.Parse()
	wsServer := &WSServer{
		ListenAddr : *addr,
	}
	wsServer.Start()
	fmt.Println("------end-------")
}*/

