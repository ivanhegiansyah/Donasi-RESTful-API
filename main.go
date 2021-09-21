package main

import (
	"finalproject-BE/config"
	"finalproject-BE/routes"
)

func main() {
	config.InitDB()
	e := routes.NewRoute()
	e.Start(":8000")
}
