package channel

type ResponseJson struct {
	Response Response
}
type Response struct {
	Audio string
	SessionId string
	RequestId string
}
var QueueChannel chan string

var Channel chan ResponseJson