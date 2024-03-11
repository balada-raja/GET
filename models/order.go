package models

type Informasi string

const (
	Proses  Informasi = "Proses"
	Selesai Informasi = "Selesai"
)

type Order struct {
	Id            int64     `gorm:"primary_key" json:"id"`
	IdUser        int64     `json:"id_user" binding:"required"`
	IdVehicle     int64     `json:"id_vehicle" binding:"required"`
	IdVendor      int64     `json:"id_vendor" binding:"required"`
	Status        Informasi `gorm:"type:enum('Proses', 'Selesai');default:'Proses'" json:"status"`
	IdDetailOrder int64     `json:"id_detail_order" binding:"required"`

	Users       Users       `gorm:"foreignKey:id_user"`
	Vehicle     Vehicle     `gorm:"foreignKey:id_vehicle"`
	Vendor      Vendor      `gorm:"foreignKey:id_vendor"`
	DetailOrder DetailOrder `gorm:"foreignKey:id_detail_order"`
}
