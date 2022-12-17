package models

type Data struct {
	//beritahu ke gorm bahwa id adalah primary key
	//panggil ke postman pakai ailas json.cnth jumlah_produk
	//struct wajib capital huruf depan

	//di db nama tablenya : Id, nama_toko, deskripsi, jumlah_product
	Id     			int64  `gorm:"primaryKey" json:"id"`
	Nama_toko		string `gorm:"type:varchar(300)" json:"nama_toko"`
	Deskripsi  		string `gorm:"type:text(2000)" json:"Deskripsi"`
	JumlahProduct  	int64 `gorm:"type:integer(255)" json:"jumlah_produk"`
}