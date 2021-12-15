package main

import (
	"host-container-ms/dataBase"
	"host-container-ms/router"
	"log"
)

func main() {
	if dataBase.GetDataBase() != nil {
		router.SetupRouter()
	} else {
		log.Fatal("Failed to setup database")
	}
}
