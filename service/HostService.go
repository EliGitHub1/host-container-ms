package service

import (
	"github.com/gin-gonic/gin"
	"host-container-ms/utils"
)

func Host(c *gin.Context) {
	query := "SELECT * FROM hosts"
	utils.ExecuteQueryHost(c, query)
}

func HostById(c *gin.Context, hostId string) {
	query := "SELECT * FROM hosts where id=" + hostId
	utils.ExecuteQueryHost(c, query)
}
