package tokocontroller

import (
	"fmt"
	"encoding/json"
	"net/http"

	"github.com/bagas050201/microservices-toko/models"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	fmt.Println("get index")
	var data_toko []models.Data

	models.DB.Find(&data_toko)
	c.JSON(http.StatusOK, gin.H{"toko": data_toko})
}

func Show(c *gin.Context) {
	var data_toko models.Data
	//ambil routes id
	id := c.Param("id")

	//check jika id tidak ditemukan
	if err := models.DB.First(&data_toko, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"toko": data_toko})
}

func Create(c *gin.Context) {
	//membuat 1 variavle product dengan 1 struct Data
	var data_toko models.Data

	//ambil data json. err != nil aslinya ada if didepannya
	if err := c.ShouldBindJSON(&data_toko); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&data_toko)
	c.JSON(http.StatusOK, gin.H{"toko": data_toko})
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil ditambahkan"})
}

func Update(c *gin.Context) {
	var data_toko models.Data
	//ambil param url
	id := c.Param("id")

	if err := c.ShouldBindJSON(&data_toko); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	//check db dimana id sebagai patokan lalu updates dengan memasukan var product. jika ternyata == 0 maka data tidak ditemukan/error
	if models.DB.Model(&data_toko).Where("id = ?", id).Updates(&data_toko).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate product"})
		return
	}
	//jika product berhasil di update
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})
}

func Delete(c *gin.Context) {
	var data_toko models.Data

	//untuk menampung item id
	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	
	id, _ := input.Id.Int64()
	if models.DB.Delete(&data_toko, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}