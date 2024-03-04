package main

import (
	"main/controller"
)

func main() {

	server := controller.NewServer()
	server.Run()
}
