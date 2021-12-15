package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"host-container-ms/model"
	"host-container-ms/service"
	"net/http"
	"strconv"
)

func Container(c *gin.Context) {
	service.Container(c)
}

func ContainerById(c *gin.Context) {
	containerId := c.Param("id")
	n, err := strconv.Atoi(containerId)
	if err != nil || n < 0 {
		c.IndentedJSON(http.StatusBadRequest, containerId)
		return
	}
	service.ContainerById(c, containerId)
}

func CreateContainer(c *gin.Context) {
	container := model.ContainerRec{}
	if err := c.ShouldBindJSON(&container); err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusUnauthorized, err)
		return
	}
	service.CreateContainer(c, container)
}
