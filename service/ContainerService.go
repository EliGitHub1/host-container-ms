package service

import (
	"github.com/gin-gonic/gin"
	"host-container-ms/model"
	"host-container-ms/utils"
	"net/http"
)

func Container(c *gin.Context) {
	query := "SELECT * FROM containers"
	utils.ExecuteQueryContainer(c, query)
}

func ContainerById(c *gin.Context, hostId string) {
	query := "SELECT * FROM containers where id=" + hostId
	utils.ExecuteQueryContainer(c, query)
}

func CreateContainer(c *gin.Context, container model.ContainerRec) {
	c.IndentedJSON(http.StatusOK, container)
}
