package main

import (
	"goServer/db"

	"goServer/tasks/endpoints/DELETE"
	"goServer/tasks/endpoints/GET"
	"goServer/tasks/endpoints/PATCH"
	"goServer/tasks/endpoints/POST"

	u "goServer/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitializeDb()
	defer db.Db.Close()

	db.CreateTables()

	router := gin.Default()
	router.SetTrustedProxies([]string{"localhost"})
	router.Use()

	router.GET("/tasks/:id", GET.TasksGET)
	router.GET("/tasks", GET.TasksGET)

	router.POST("/tasks", POST.TasksPOST)

	router.PATCH("/tasks/:id", PATCH.TasksPatch)
	router.DELETE("/tasks/:id", DELETE.TasksDELETE)

	router.NoRoute(func(ctx *gin.Context) {
		route := ctx.Request.URL.Path
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":           "Route " + route + " is not implemented.",
			"availableRoutes": "/tasks",
		})
	})

	router.Run(u.GetEnv(u.SERVER_PORT))
}
