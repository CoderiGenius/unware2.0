package producer





import(
"golang.org/x/net/websocket"
"fmt"
)

type WSServer struct {
	ListenAddr string
}


//Deal 发来的语音struct
func  Deal(content string)  {
	//开始接受producer内容
	//当前业务为姓名



	}


}
func GetMessage(conn *websocket.Conn) string{

	var err error
	var reply string
	err = websocket.Message.Receive(conn,reply)
	if err  != nil {
		panic(err)
		return ""
	}
	fmt.Println("GET MESSAGE FROM: ",conn.RemoteAddr().String()," MESSAGE：",reply)
	return reply
}