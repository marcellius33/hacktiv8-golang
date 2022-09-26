package main

import "practice-rest-api/httpserver"

func main() {
	app := httpserver.CreateRouter()
	app.Run(":8080")
}
