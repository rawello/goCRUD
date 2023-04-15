package main

import (
	"github.com/rawello/goCRUD/models"
	"github.com/rawello/goCRUD/routes"
)

func main() {

	db := models.SetupDB()
	db.AutoMigrate(&models.Task{})

	r := routes.SetupRoutes(db)
	r.Run()
}
