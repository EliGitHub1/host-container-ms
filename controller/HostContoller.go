package controller

import (
	"github.com/gin-gonic/gin"
	"host-container-ms/service"
	"net/http"
	"strconv"
)

func Host(c *gin.Context) {
	service.Host(c)
}

func HostById(c *gin.Context) {
	hostId := c.Param("id")
	n, err := strconv.Atoi(hostId)
	if err != nil || n < 0 {
		c.IndentedJSON(http.StatusBadRequest, hostId)
		return
	}
	service.HostById(c, hostId)
}
