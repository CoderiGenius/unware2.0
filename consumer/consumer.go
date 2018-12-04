package consumer

import ("../queue"
	"fmt"
	"../speaker"
)
var Channel chan speaker.Response
//存入最终显示队列
func GetFinalQueueFromChannel(){
	Channel = make(chan speaker.Response,10)
	//新建最终显示队列
	Show := queue.Queue{}
	for{
		temp:=<-Channel
		if !Show.FindThingsAlreadyInside(temp){
			Show.Push(temp)
		}

	}
}

func DealWithFinalQueue(show *queue.Queue)  speaker.Response{

	for{
		if !show.IsEmpty(){
			s:=show.Pop()
			fmt.Println(s)
			return s.(speaker.Response)
		}
	}

}