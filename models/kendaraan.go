package models

type JenisKendaraan string
type Status string

const (
	Mobil JenisKendaraan = "Mobil"
	Motor JenisKendaraan = "Motor"

	Ready       Status = "Ready"
	Used        Status = "Used"
	Maintenance Status = "Maintenance"
)

type Kendaraan struct {
	Id             int64          `gorm:"primary_key" json:"id"`
	Nama           string         `gorm:"type:varchar(255);not null" json:"nama" binding:"required"`
	JenisKendaraan JenisKendaraan `gorm:"type:enum('Mobil','Motor');not null" json:"jenis_kendaraan" binding:"required"`
	NomorPolisi    string         `gorm:"type:varchar(255);unique ;not null" json:"nomor_polisi" binding:"required"`
	NomorMesin     string         `gorm:"type:varchar(255);unique ;not null" json:"nomor_mesin" binding:"required"`
	Deskripsi      string         `gorm:"type:text" json:"deskripsi"`
	Status         Status         `gorm:"type:enum('Ready', 'Used', 'Maintenance'); not null; default:'Ready'" json:"status" binding:"required"`
	Harga          float64        `gorm:"type:double; not null" json:"harga" binding:"required"`
	IdPenyediaJasa int64          `json:"id_penyedia_jasa" binding:"required"`

	PenyediaJasa PenyediaJasa `gorm:"foreignKey:IdPenyediaJasa"`
}
