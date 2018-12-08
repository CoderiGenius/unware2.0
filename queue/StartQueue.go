package queue

import ("fmt"
"../speaker"
"../channel"
)




var GetChannel chan int
func StartQueue() {

	//新建channel用以储存队列
	channel.QueueChannel = make(chan channel.Content, 1)

	//新建取queue
	GetChannel = make(chan int,1)
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
			GetChannel <- 1

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
			<- GetChannel
			//fmt.Println("continue")
			//runtime.Gosched()
		}
	}

}