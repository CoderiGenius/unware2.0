package producer





import(
"golang.org/x/net/websocket"
"fmt"
"../speaker"
"../consumer"
)

type WSServer struct {
	ListenAddr string
}


//Deal 发来的语音struct
func  Deal(response speaker.Response)  {
	//开始接受producer内容
	//当前业务为姓名

	consumer.Channel <- response




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