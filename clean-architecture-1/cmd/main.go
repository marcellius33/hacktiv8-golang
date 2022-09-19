package main

import "clean-architecture-1/httpserver"

func main() {
	app := httpserver.CreateRouter()
	app.Run(":8000")
}
