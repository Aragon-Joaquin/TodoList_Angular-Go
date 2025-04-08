package GET

import (
	"goServer/db"
	e "goServer/errors"
	t "goServer/tasks/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TasksGET(c *gin.Context) {
	taskID := c.Params.ByName("id")

	if taskID == "" {
		res := tasksGetAll()
		c.JSON(200, res)
	}
	code, res := tasksGetOne(taskID)
	c.JSON(int(code), res)
}

func tasksGetAll() []t.TaskStruct {
	res, err := db.Db.Query(`
	SELECT name, description, status, photo, hex_color 
	FROM tasks 	
	`)

	e.CheckErr(err)

	var task t.TaskStruct
	tasksSlice := []t.TaskStruct{}

	for res.Next() {
		err := res.Scan(&task.Name, &task.Description, &task.Status, &task.Photo, &task.Hex_color)
		e.CheckErr(err)

		tasksSlice = append(tasksSlice, task)
	}

	return tasksSlice

}

func tasksGetOne(id string) (t.StatusCode, interface{}) {
	var task t.TaskStruct

	err := db.Db.QueryRow(`
	SELECT name, description, status, photo, hex_color 
	FROM tasks WHERE $1 = id	
	`, id).Scan(&task.Name, &task.Description, &task.Status, &task.Photo, &task.Hex_color)

	e.CheckSQLErr(err)

	if (t.TaskStruct{}) == task {
		return http.StatusBadRequest, t.ServerMsgStruct{ErrorMessage: "User not found."}
	}
	return http.StatusOK, task

}
