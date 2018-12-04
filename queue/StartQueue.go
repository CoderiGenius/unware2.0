package queue

import ("fmt"
"../speaker"
"../channel"
)





func StartQueue() {

	//新建channel用以储存队列
	channel.QueueChannel = make(chan string, 10)
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

		}

	}
}

//队列处理
func DealTheQueue(q *Queue)  {

	for{
		if !q.IsEmpty(){
			s:=q.Pop()
		fmt.Println(s)
		go speaker.Speak(s)
		}
	}

}