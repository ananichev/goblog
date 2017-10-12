package main

import (
	"fmt"
	"github.com/ananichev/goblob/accountservice/dbclient"
	"github.com/ananichev/goblob/accountservice/service"
)

var appName = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	initializeBotlClient()
	service.StartWebServer("6767")
}

func initializeBotlClient() {
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDb()
	service.DBClient.Seed()
}
