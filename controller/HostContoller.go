package controller

import (
	"github.com/gin-gonic/gin"
	"host-container-ms/dataBase"
	"host-container-ms/model"
	"net/http"
	"strconv"
)

func Host(c *gin.Context) {
	hostId := c.Param("id")
	n, err := strconv.Atoi(hostId)
	if hostId != "" && err != nil || n < 0 {
		c.IndentedJSON(http.StatusBadRequest, hostId)
		return
	}
	var hosts []model.HostRec
	query := ""
	if hostId == "" {
		query = "SELECT * FROM hosts"
	} else {
		query = "SELECT * FROM hosts where id=" + hostId
	}
	db := dataBase.GetDataBase()
	if db == nil {
		c.IndentedJSON(http.StatusInternalServerError, hostId)
		return
	}
	rows, _ := db.Query(query)
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
		c.IndentedJSON(http.StatusNoContent, hostId)
	} else {
		c.IndentedJSON(http.StatusOK, hosts)
	}
}
