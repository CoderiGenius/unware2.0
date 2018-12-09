package WebsocketHandler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"../channel"
)



func PostHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	con, _ := ioutil.ReadAll(r.Body) //获取post的数据
	fmt.Println(string(con))
	var reply RequstJSON
	err2 := json.Unmarshal(con, &reply)
	if err2 != nil {
		fmt.Println("json erro")
		w.Write([]byte("json erro"))

	}

	if reply.From == "producer" {

		channel.QueueChannel <- reply.Content
		w.Write([]byte("success"))

	}
}

func (this *Server) StartPostListener() (error) {
	//配置handle
	//http.Handle("/ws", websocket.Handler(this.Handler))

	fmt.Println("begin to listen post")
	//开始启用handler
	//err := http.ListenAndServe(this.ListenAddr, nil)
	//
	//if err != nil {
	//	fmt.Println("ListenAndServe:", err)
	//	return err
	//}

	http.HandleFunc("/", PostHandler)
	err2 :=http.ListenAndServe(this.ListenAddr, nil)
	fmt.Println("begin to listen")
	if err2 != nil {
		fmt.Println("ListenAndServe:", err2)
		return err2
	}

	fmt.Println("start end")
	return nil
}

//func sendToChannel(queueChannel chan string)  {
//
//	<- channel
//}