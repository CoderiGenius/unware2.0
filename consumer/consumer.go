package consumer

import ("../queue"
	"fmt"
	"../channel"
)


var Show queue.Queue

//存入最终显示队列
func GetFinalQueueFromChannel(){
	channel.Channel = make(chan channel.ResponseJson,10)
	//新建最终显示队列

	for{
		temp:=<-channel.Channel
		if !Show.FindThingsAlreadyInside(temp){
			Show.Push(temp)
			fmt.Println("show queue pushed!",temp)
		}

	}
}

func DealWithFinalQueue()  channel.ResponseJson{

	for{
		if !Show.IsEmpty(){
			s:=Show.Pop()
			fmt.Println(s)
			return s.(channel.ResponseJson)
		}
	}

}