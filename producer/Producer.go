package producer





import(
"golang.org/x/net/websocket"
"fmt"
"../channel"

)

type Response struct {
	Audio string
	SessionId string
	RequestId string
}

type WSServer struct {
	ListenAddr string
}


//Deal 发来的语音struct
func  Deal(response channel.ResponseJson)  {
	//开始接受producer内容
	//当前业务为姓名
	//fmt.Println("before push",response)
	channel.Channel <- response




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