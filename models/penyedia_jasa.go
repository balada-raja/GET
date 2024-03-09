package models

type PenyediaJasa struct {
	Id         int64       `gorm:"primary_key" json:"id"`
	Nama       string      `gorm:"type:varchar(255); not null" json:"nama" binding:"required"`
	Alamat     string      `gorm:"type:varchar(255); not null" json:"alamat" binding:"required"`
	Rating     float64     `gorm:"type:float" json:"rating"`
}
