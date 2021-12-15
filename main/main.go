package main

import (
	"host-container-ms/Router"
	"host-container-ms/dataBase"
	"log"
)

func main() {
	if dataBase.GetDataBase() != nil {
		Router.SetupRouter()
	} else {
		log.Fatal("Failed to setup database")
	}
}
