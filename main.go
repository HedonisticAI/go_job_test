package main

import (
	"go_job_test/models"
	"go_job_test/models/handlers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	pagination "github.com/webstradev/gin-pagination"
)

func main() {
	err := godotenv.Load("api.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Print("Got data from .env file")
	db := models.Init_Database()
	if db == nil {
		return
	}
	var Handler = handlers.Handler{db}
	r := gin.Default()

	r.Use(pagination.Default())

	r.GET("get/all", Handler.GetAll)
	r.GET("get/by_name/:name", Handler.GetByName)
	r.GET("get/by_surname/:surname", Handler.GetBySurname)
	r.GET("get/by_age/:age/*condition*")
	r.DELETE("/delete/:id")
	r.POST("/create", Handler.Create)
	r.PATCH("/update/:id", Handler.Update)
	r.Run(os.Getenv("API_PORT"))
}
