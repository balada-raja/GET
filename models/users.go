package models

type role string

const (
	Pelanggan role = "Pelanggan"
	Premium   role = "Premium"
	Admin     role = "Admin"
)

type Users struct {
	Id       int64  `gorm:"primary_key" json:"id"`
	Nama     string `gorm:"type:varchar(255); not null" json:"name" binding:"required"`
	Email    string `gorm:"type:varchar(255); not null; unique" json:"email" binding:"required"`
	Password string `gorm:"type:varchar(255); not null" json:"password" binding:"required"`
	NomorHP  string `gorm:"type:varchar(255); not null; unique" json:"nomor_hp" binding:"required"`
	Role     role   `gorm:"type:enum('Pelanggan','Premium','Admin'); default:'Pelanggan; not null'" json:"role"`
	Ktp      string `gorm:"type:varchar(255)" json:"KTP"`
	Alamat   string `gorm:"type:varchar(255)" json:"alamat"`
}
