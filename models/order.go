package models

type Informasi string

const (
	Proses  Informasi = "Proses"
	Selesai Informasi = "Selesai"
)

type Order struct {
	Id             int64 `gorm:"primary_key" json:"id"`
	IdUser         int64
	IdKendaraan    int64
	IdPenyediaJasa int64
	Status         Informasi `gorm:"type:enum('Proses', 'Selesai');default:'Proses'" json:"status"`
	IdDetailOrder  int64

	Users        Users        `gorm:"foreignKey:IdUser"`
	Kendaraan    Kendaraan    `gorm:"foreignKey:IdKendaraan"`
	PenyediaJasa PenyediaJasa `gorm:"foreignKey:IdPenyediaJasa"`
	DetailOrder  DetailOrder  `gorm:"foreignKey:IdDetailOrder"`
}
