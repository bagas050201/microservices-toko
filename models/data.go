package models

type informasi_toko struct {
	//beritahu ke gorm bahwa id adalah primary key
	Id_toko          int64  `gorm:"primaryKey" json:"id"`
	nama_toko string `gorm:"type:varchar(300)" json:"nama_toko"`
	deskripsi  string `gorm:"type:text" json:"deskripsi"`
	jumlah_product  string `gorm:"type:text" json:"jumlah_product"`
}