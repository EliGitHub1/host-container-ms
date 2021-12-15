package service

import (
	"github.com/gin-gonic/gin"
	"host-container-ms/DAO"
)

func Host(c *gin.Context) {
	query := "SELECT * FROM hosts"
	DAO.ExecuteQueryHost(c, query)
}

func HostById(c *gin.Context, hostId string) {
	query := "SELECT * FROM hosts where id=" + hostId
	DAO.ExecuteQueryHost(c, query)
}
