package channel

type ResponseJson struct {
	Response Response
}
type Response struct {
	Audio string
	SessionId string
	RequestId string
}
var Json ResponseJson



type Content struct{
	Name string
	Class string
	Photo string
}
var QueueChannel chan Content

var Channel chan ResponseJson

var SendMessageToConsumer chan int

var Count int