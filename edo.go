package main

import (
	"os"

	"gitlab.com/bokjo/test_edo/api"
)

func main() {

	//TODO: Implement it as envirnmet variables or move it to a separate nonversioned config file
	username := os.Getenv("EDOAPI_USERNAME")
	password := os.Getenv("EDOAPI_PASSWORD")
	database := os.Getenv("EDOAPI_DB")
	host := os.Getenv("EDOAPI_HOST")
	port := os.Getenv("EDOAPI_PORT")

	edoapi := api.API{}
	edoapi.Init(username, password, database, host)
	edoapi.Run(port)

}
