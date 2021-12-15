package controller

import (
	"github.com/gin-gonic/gin"
	"host-container-ms/dataBase"
	"host-container-ms/model"
	"net/http"
	"strconv"
)

func Container(c *gin.Context) {
	containerId := c.Param("id")
	n, err := strconv.Atoi(containerId)
	if containerId != "" && err != nil || n < 0 {
		c.IndentedJSON(http.StatusBadRequest, containerId)
		return
	}
	var containers []model.ContainerRec
	query := ""
	if containerId == "" {
		query = "SELECT * FROM containers"
	} else {
		query = "SELECT * FROM containers where id=" + containerId
	}
	db := dataBase.GetDataBase()
	if db == nil {
		c.IndentedJSON(http.StatusInternalServerError, containerId)
		return
	}
	rows, _ := db.Query(query)
	var id int
	var host_id int
	var name string
	var image_name int
	for rows.Next() {
		rows.Scan(&id, &host_id, &name, &image_name)
		containers = append(containers, model.ContainerRec{
			ID:        id,
			HostId:    host_id,
			Name:      name,
			ImageName: image_name,
		})
	}
	defer rows.Close()
	if len(containers) == 0 {
		c.IndentedJSON(http.StatusNoContent, containerId)
	} else {
		c.IndentedJSON(http.StatusOK, containers)
	}
}

func ContainerByHost(c *gin.Context) {
	return
}
