package models

import (
	//"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func ConnectDatabase() {
	//go_restapi_mysql adalah nama table
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/microservices_toko"))
	//check jika database gagal terkoneksi
	if err != nil {
		panic(err)
	}

	//auto migrate file data.go
	database.AutoMigrate(&informasi_toko{})

	DB = database
}