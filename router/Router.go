package router

import (
	"github.com/gin-gonic/gin"
	"host-container-ms/controller"
)

const PORT string = "localhost:8080"

var (
	r = gin.Default()
)

func SetupRouter() {
	r.GET("/container", controller.Container)
	r.GET("/container/:id", controller.ContainerById)
	r.POST("/container", controller.CreateContainer)
	r.GET("/host/:id", controller.HostById)
	r.GET("/host", controller.Host)
	r.Run(PORT)
}
