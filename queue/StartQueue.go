package queue

import ("fmt"
"../speaker"
)


var QueueChannel chan string


func StartQueue() {

	//新建channel用以储存队列
	QueueChannel = make(chan string, 10)
	//新建姓名队列
	Q := Queue{}

	go DealTheQueue(&Q)

	for {
		s := <-QueueChannel
		if Q.FindThingsAlreadyInside(s) {
			fmt.Println("%s Already exsits ", s)
		} else {
			Q.Push(s)
			fmt.Println("q pushed!")
			//fmt.Println(Q.Pop())}
		}
		//fmt.Println(q.Pop())
		//fmt.Println(q.Pop())
		//fmt.Println(q.IsEmpty())
		//fmt.Println(q.Pop())
		//fmt.Println(q.IsEmpty())
		//fmt.Println(q.FindThingsAlreadyInside(1))
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