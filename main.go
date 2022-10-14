package main

import (
	"github.com/oivinig/api-go-gin/database"
	"github.com/oivinig/api-go-gin/routes"
)

func main() {
	database.Connect()
	routes.Handler()
}
