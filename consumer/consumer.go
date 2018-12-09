package consumer

import (
	"../channel"
	"../queue"
)


var Show queue.Queue

//存入最终显示队列
func GetFinalQueueFromChannel(){
	channel.Channel = make(chan channel.ResponseJson,10)
	//新建最终显示队列
	channel.SendMessageToConsumer = make(chan int,1)

	for{
		temp:=<-channel.Channel

		if !Show.FindNameAlreadyInFinalQueue(temp){
			if temp.Response.SessionId != channel.Json.Response.SessionId{
				Show.Push(temp)
				//fmt.Println("show queue pushed!",temp)
				channel.Json = temp
			}
			//channel.SendMessageToConsumer <- 1

		}


	}
}

func DealWithFinalQueue()  channel.ResponseJson{



		if !Show.IsEmpty(){
			s:=Show.Pop()
			//fmt.Println(s)
			return s.(channel.ResponseJson)
		}else {
			//<- channel.SendMessageToConsumer

		return channel.ResponseJson{}
			}

}