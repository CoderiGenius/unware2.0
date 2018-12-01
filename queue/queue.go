package queue

// A FIFO queue.
type Queue []interface{}

// Pushes the element into the queue.
// 		e.g. q.Push(123)
func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

// Pops element from head.
func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

// Returns if the queue is empty or not.
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) FindThingsAlreadyInside(v interface{}) bool  {

	for i:=0;i<len(*q);i++{
		temp := (*q)[i]
		if v==temp{
			return true
		}else {
			continue
		}
	}
	return false
}
