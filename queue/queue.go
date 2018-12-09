package queue

import "../channel"

func NewQueue() *Queue{
	//type QueueV []interface{}
	QueueV := new(Queue)
	lockChannel = make(chan int,1)
	return QueueV
}

//A FIFO queue.
type Queue []interface{}




var lockChannel chan int

// Pushes the element into the queue.
// 		e.g. q.Push(123)
func (q *Queue) Push(v interface{}) {

	//堵塞
	lockChannel <- 1
	*q = append(*q, v)
	//释放通道
	<- lockChannel
}

// Pops element from head.
func (q *Queue) Pop() interface{} {
	//堵塞
	lockChannel <- 1
	head := (*q)[0]
	*q = (*q)[1:]
	//释放
	<- lockChannel
	return head
}

// Returns if the queue is empty or not.
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) FindNameAlreadyInQueue(v interface{}) bool  {

	for i:=0;i<len(*q);i++{
		temp := (*q)[i]
		if v.(channel.Content).Name==temp.(channel.Content).Name{
			return true
		}else {
			continue
		}
	}
	return false
}
func (q *Queue) FindNameAlreadyInFinalQueue(v interface{}) bool  {

	for i:=0;i<len(*q);i++{
		temp := (*q)[i]
		if v.(channel.ResponseJson).Response.SessionId==temp.(channel.ResponseJson).Response.SessionId{
			return true
		}else {
			continue
		}
	}
	return false
}
