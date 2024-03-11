package models

type Vendor struct {
	Id         int64       `gorm:"primary_key" json:"id"`
	Name       string      `gorm:"type:varchar(255); not null" json:"name" binding:"required"`
	Address     string      `gorm:"type:varchar(255); not null" json:"address" binding:"required"`
	Rating     float64     `gorm:"type:float" json:"rating"`
}
