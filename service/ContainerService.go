package service

import (
	"github.com/gin-gonic/gin"
	"host-container-ms/DAO"
)

func Container(c *gin.Context) {
	query := "SELECT * FROM containers"
	DAO.ExecuteQueryContainer(c, query)
}

func ContainerByHostId(c *gin.Context, hostId string) {

	query := "SELECT A.id,A.host_id,A.name,A.image_name " +
		"FROM containers A " +
		"INNER JOIN hosts B " +
		"on A.host_id = B.id " +
		"where " + hostId + " = B.id"
	DAO.ExecuteQueryContainer(c, query)
}

func ContainerById(c *gin.Context, hostId string) {
	query := "SELECT * FROM containers where id=" + hostId
	DAO.ExecuteQueryContainer(c, query)
}

func CreateContainer(c *gin.Context, image_name int, host_id int, name string) {
	query := `INSERT INTO containers(image_name, host_id, name) VALUES (?, ?, ?)`
	DAO.InsertContainer(c, query, image_name, host_id, name)
}
