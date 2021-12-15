package controller

import (
	"github.com/gin-gonic/gin"
	"host-container-ms/model"
	"host-container-ms/service"
	"net/http"
	"strconv"
)

func Container(c *gin.Context) {
	hostId := c.Query("host_id")
	if hostId == "" {
		service.Container(c)
	}
	n, err := strconv.Atoi(hostId)
	if err != nil || n < 0 {
		c.IndentedJSON(http.StatusBadRequest, hostId)
		return
	}
	service.ContainerByHostId(c, hostId)
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
	container := model.ContainerCreateReq{}
	if err := c.ShouldBindJSON(&container); err != nil {
		c.IndentedJSON(http.StatusUnauthorized, err)
		return
	}
	service.CreateContainer(c, container.ImageName, container.HostId, container.Name)
}
