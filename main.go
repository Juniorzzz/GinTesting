package main

import (
	"GinWeb/model"
	"GinWeb/server"
)

func main() {

	model.InitDB()

	gin := server.NewRounter()
	gin.Run(":3000")
}
