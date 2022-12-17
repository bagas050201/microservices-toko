package main

import (
	_"fmt"
	"github.com/bagas050201/microservices-toko/controllers/tokocontroller"
	"github.com/bagas050201/microservices-toko/models"
	"github.com/gin-gonic/gin"
	_"net/http"
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
	r.GET("/api/merchants", tokocontroller.Index)
	r.GET("/api/merchants/:id", tokocontroller.Show)
	r.POST("/api/merchants", tokocontroller.Create)
	r.PUT("/api/merchants/:id", tokocontroller.Update)
	r.DELETE("/api/merchants", tokocontroller.Delete)

	r.Run(":9000")
}