package PATCH

import (
	"goServer/db"
	t "goServer/tasks/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mattn/go-sqlite3"
)

func TasksPatch(c *gin.Context) {
	taskID := c.Params.ByName("id")

	if taskID == "" {
		c.JSON(int(http.StatusBadRequest), t.ServerMsgStruct{ErrorMessage: "Specify the 'ID' parameter in the URL."})
		return
	}

	var taskStr t.TaskStruct

	if err := c.BindJSON(&taskStr); err != nil {
		log.Println(err)
		return
	}

	var taskUpd t.TaskStruct

	err := db.Db.QueryRow(`
	UPDATE tasks SET 
		name = COALESCE($1, name),
		description = COALESCE($2, description),
		status = COALESCE($3, status),
		photo = COALESCE($4, photo),
		hex_color = COALESCE($5, hex_color)
	WHERE $6 = id
	`, taskStr.Name, taskStr.Description, taskStr.Status, taskStr.Photo, taskStr.Hex_color, taskID).
		Scan(&taskUpd)

	if err != nil {
		c.JSON(int(http.StatusBadRequest), t.ServerMsgStruct{ErrorMessage: err.(sqlite3.Error).Error()})
		return
	}

	c.JSON(int(http.StatusAccepted), gin.H{"task": taskUpd})
}
