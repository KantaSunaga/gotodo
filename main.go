package main

import (
	"gotodo/controllers"
	"gotodo/models"
)


func main() {
	models.DbInit()
	controllers.StartWebServer()
}