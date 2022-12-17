package tokocontroller

import (
	"fmt"
	_"encoding/json"
	"net/http"

	"github.com/bagas050201/microservices-toko/models"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	fmt.Println("get index")
	//ambil data produk dari table Data pada mysql
	var product []models.data

	models.DB.Find(&product)
	c.JSON(http.StatusOK, gin.H{"product": product})
}

func Show(c *gin.Context) {
	var product models.data
	//ambil routes id
	id := c.Param("id")

	//check jika id tidak ditemukan
	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}

func Create(c *gin.Context) {
	fmt.Println("create index")
}

func Update(c *gin.Context) {
	fmt.Println("update index")
}

func Delete(c *gin.Context) {
	fmt.Println("delete index")
}