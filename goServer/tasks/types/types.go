package types

type StatusCode int

type taskState string

var StateOfTasks = map[taskState]string{
	"isDone":      "done",
	"isPending":   "pending",
	"isCancelled": "cancelled",
}

type TaskStruct struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Photo       string    `json:"photo"`
	Hex_color   taskState `json:"hex_color"`
}

type ServerMsgStruct struct {
	ErrorMessage string
}
