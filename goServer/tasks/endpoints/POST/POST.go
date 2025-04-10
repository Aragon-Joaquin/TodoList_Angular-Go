package POST

import (
	"goServer/db"
	t "goServer/tasks/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mattn/go-sqlite3"
)

func TasksPOST(c *gin.Context) {
	var taskStr t.TaskStruct

	if err := c.BindJSON(&taskStr); err != nil {
		log.Println(err)
		return
	}

	var taskInserted t.TaskStruct

	err := db.Db.QueryRow(`
	INSERT INTO tasks(name,description,photo) VALUES($1, $2, $3) 
	RETURNING name, description, status, photo, hex_color
	`, taskStr.Name, taskStr.Description, taskStr.Photo).Scan(&taskInserted.Name, &taskInserted.Description, &taskInserted.Status, &taskInserted.Photo, &taskInserted.Hex_color)

	if err != nil {
		c.JSON(int(http.StatusBadRequest), t.ServerMsgStruct{ErrorMessage: err.(sqlite3.Error).Error()})
		return
	}

	c.JSON(int(http.StatusCreated), taskInserted)
}
