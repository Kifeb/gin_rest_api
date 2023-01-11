package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kifeb/gin_rest_api/controller/productcontroller"
	"github.com/kifeb/gin_rest_api/models"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/products/:id", productcontroller.Show)
	r.POST("/api/products", productcontroller.Create)
	r.PUT("/api/products/:id", productcontroller.Update)

	r.Run()
}
