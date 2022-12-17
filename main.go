package main

import (
	"fmt"
	"github.com/bagas050201/microservices-toko/controllers/tokocontroller"
	"github.com/bagas050201/microservices-toko/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/gin-contrib/cors"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"POST", "GET", "DELETE"},
		AllowHeaders:  []string{"Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
	}))
	//atau allow all dengan
	//r.Use(cors.Default())

	r.GET("/", tokocontroller.Index)
	r.GET("/api/product", tokocontroller.Index)
	r.GET("/api/product/:id", tokocontroller.Show)
	r.POST("/api/product", tokocontroller.Create)
	r.PUT("/api/product/:id", tokocontroller.Update)
	r.DELETE("/api/product", tokocontroller.Delete)

	fmt.Println("server toko backend berjalan")
	r.Run(":9000")
}