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
	channel.SendMessageToConsumer = make(chan int,1)

	for{
		temp:=<-channel.Channel

		if !Show.FindNameAlreadyInFinalQueue(temp){
			//channel.SendMessageToConsumer <- 1
			Show.Push(temp)
			fmt.Println("show queue pushed!",temp)
		}

	}
}

func DealWithFinalQueue()  channel.ResponseJson{



		if !Show.IsEmpty(){
			s:=Show.Pop()
			fmt.Println(s)
			return s.(channel.ResponseJson)
		}else {
			//<- channel.SendMessageToConsumer

		return channel.ResponseJson{}
			}

}