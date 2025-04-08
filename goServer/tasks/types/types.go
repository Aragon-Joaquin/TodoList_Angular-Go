package types

type StatusCode int

type taskState string

var StateOfTasks = map[taskState]string{
	"isDone":      "done",
	"isPending":   "pending",
	"isCancelled": "cancelled",
}

type TaskStruct struct {
	Id          int
	Name        string
	Description string
	Status      string
	Photo       string
	Hex_color   taskState
}

type ServerMsgStruct struct {
	ErrorMessage string
}
