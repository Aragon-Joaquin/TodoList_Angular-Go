package DELETE

import (
	"goServer/db"
	t "goServer/tasks/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mattn/go-sqlite3"
)

/* //! i have no idea why:
- This file throws an error if its called "DELETE.go" (No packages found for open file)
- How does goLang package structure works
- or why does "main.go" or any name fixes this
*/

func TasksDELETE(c *gin.Context) {

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

	var taskDel t.TaskStruct

	err := db.Db.QueryRow(`
	DELETE FROM tasks WHERE id = $1 RETURNING name
	`, taskID).Scan(&taskDel.Name)

	if err != nil {
		c.JSON(int(http.StatusBadRequest), t.ServerMsgStruct{ErrorMessage: err.(sqlite3.Error).Error()})
		return
	}

	c.JSON(int(http.StatusOK), taskDel)

}
