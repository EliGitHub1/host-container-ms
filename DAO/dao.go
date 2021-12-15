package DAO

import (
	"github.com/gin-gonic/gin"
	"host-container-ms/dataBase"
	"host-container-ms/model"
	"net/http"
)

func ExecuteQueryHost(c *gin.Context, query string) {
	var hosts []model.HostRec
	db := dataBase.GetDataBase()
	if db == nil {
		c.IndentedJSON(http.StatusInternalServerError, hosts)
		return
	}
	rows, err := db.Query(query)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	var id int
	var uuid string
	var name string
	var ip_address string
	for rows.Next() {
		rows.Scan(&id, &uuid, &name, &ip_address)
		hosts = append(hosts, model.HostRec{
			ID:        id,
			Uuid:      uuid,
			Name:      name,
			IpAddress: ip_address,
		})
	}
	defer rows.Close()
	if len(hosts) == 0 {
		c.IndentedJSON(http.StatusNoContent, hosts)
	} else {
		c.IndentedJSON(http.StatusOK, hosts)
	}
}

func ExecuteQueryContainer(c *gin.Context, query string) {
	var containers []model.ContainerRec
	db := dataBase.GetDataBase()
	if db == nil {
		c.IndentedJSON(http.StatusInternalServerError, containers)
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
		c.IndentedJSON(http.StatusNoContent, containers)
	} else {
		c.IndentedJSON(http.StatusOK, containers)
	}
}

func InsertContainer(c *gin.Context, query string, image_name int, host_id int, name string) {
	db := dataBase.GetDataBase()
	if db == nil {
		c.IndentedJSON(http.StatusInternalServerError, image_name)
		return
	}
	statement, err := db.Prepare(query)
	if err != nil {
		panic(err.Error())
	}
	_, err = statement.Exec(image_name, host_id, name)
	if err != nil {
		panic(err.Error())
	}
	c.IndentedJSON(http.StatusOK, image_name)
}
