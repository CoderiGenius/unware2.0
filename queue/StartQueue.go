package queue

import ("fmt"
"../speaker"
"../channel"
	"runtime"
)





func StartQueue() {

	//新建channel用以储存队列
	channel.QueueChannel = make(chan string, 1)

	//新建取queue
	getChannel := make(chan int,1)
	//新建姓名队列
	//Q := Queue{}
	Q := NewQueue()
	go DealTheQueue(Q)

	for {
		s := <-channel.QueueChannel
		if Q.FindThingsAlreadyInside(s) {
			fmt.Println("%s Already exsits ", s)
		} else {
			Q.Push(s)
			fmt.Println("q pushed!")
			getChannel <- 1

		}

	}
}

//队列处理
func DealTheQueue(q *Queue)  {
	//DealTheQueueChannel := make(chan string,1)
	//DealTheQueueChannel <-

	for{
		if !q.IsEmpty(){
			s:=q.Pop()
		fmt.Println("queue pop:",s)
		go speaker.Speak(s)
		}else {
			//fmt.Println("continue")
			runtime.Gosched()
		}
	}

}