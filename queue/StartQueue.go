package queue

import "fmt"


var QueueChannel chan string


func StartQueue()  {

	//新建channel用以储存队列
	QueueChannel = make(chan string, 10)
	//新建队列
	Q := Queue{}
	for {
		Q.Push(<-QueueChannel)
		fmt.Println("q pushed!")
		fmt.Println(Q.Pop())
	}
	//fmt.Println(q.Pop())
	//fmt.Println(q.Pop())
	//fmt.Println(q.IsEmpty())
	//fmt.Println(q.Pop())
	//fmt.Println(q.IsEmpty())
	//fmt.Println(q.FindThingsAlreadyInside(1))
}