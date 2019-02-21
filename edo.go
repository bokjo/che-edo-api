package main

import "gitlab.com/bokjo/test_edo/api"

func main() {

	edoapi := api.API{}

	//TODO: Implement it as envirnmet variables and move it to the api.go
	edoapi.Init("postgres", "postgres", "edo-api", "db")
	edoapi.Run()
}
