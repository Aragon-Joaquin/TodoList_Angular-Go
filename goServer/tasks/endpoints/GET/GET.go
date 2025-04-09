package GET

import (
	"goServer/db"
	e "goServer/errors"
	t "goServer/tasks/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TasksGET(c *gin.Context) {
	taskID := c.Params.ByName("id")

	if taskID == "" {
		code, res := tasksGetAll()
		c.JSON(int(code), res)
		return
	}
	code, res := tasksGetOne(taskID)
	c.JSON(int(code), res)
}

func tasksGetAll() (t.StatusCode, interface{}) {
	res, err := db.Db.Query(`
	SELECT name, description, status, photo, hex_color 
	FROM tasks 	
	`)

	if err != nil {
		log.Fatalln(err)
	}

	var task t.TaskStruct
	tasksSlice := []t.TaskStruct{}

	for res.Next() {
		err := res.Scan(&task.Name, &task.Description, &task.Status, &task.Photo, &task.Hex_color)

		if err != nil {
			log.Fatalln(err)
		}

		tasksSlice = append(tasksSlice, task)
	}

	if len(tasksSlice) == 0 {
		return http.StatusOK, make([]int, 0)
	}
	return http.StatusOK, tasksSlice

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
