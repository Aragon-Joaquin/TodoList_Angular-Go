package POST

import (
	"goServer/db"
	e "goServer/errors"
	t "goServer/tasks/types"
	"log"

	"github.com/gin-gonic/gin"
)

func TasksPOST(c *gin.Context) {
	var taskStr t.TaskStruct

	if err := c.BindJSON(&taskStr); err != nil {
		log.Println(err)
	}

	var taskInserted t.TaskStruct

	err := db.Db.QueryRow(`
	INSERT INTO tasks(name,description,photo) VALUES($1, $2, $3) 
	RETURNING name, description, status, photo, hex_color
	`).Scan(&taskInserted.Name, &taskInserted.Description, &taskInserted.Status, &taskInserted.Photo, &taskInserted.Hex_color)

	e.CheckErr(err)

	c.JSON(201, taskInserted)
}
