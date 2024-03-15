package models

type VehicleType string
type Status string
type Transmission string

const (
	Mobil VehicleType = "Car"
	Motor VehicleType = "Motor"

	Ready       Status = "Ready"
	Used        Status = "Used"
	Maintenance Status = "Maintenance"

	Manual Transmission = "Manual"
	Matic  Transmission = "Matic"
)

type Vehicle struct {
	Id            int64        `gorm:"primary_key" json:"id"`
	Name          string       `gorm:"type:varchar(255);not null" json:"name" binding:"required"`
	VehicleType   VehicleType  `gorm:"type:enum('Car','Motor');not null" json:"vehicle_type" binding:"required"`
	PoliceNumber  string       `gorm:"type:varchar(255);unique ;not null" json:"police_number" binding:"required"`
	MachineNumber string       `gorm:"type:varchar(255);unique ;not null" json:"machine_number" binding:"required"`
	Description   string       `gorm:"type:text" json:"description"`
	Status        Status       `gorm:"type:enum('Ready', 'Used', 'Maintenance'); not null; default:'Ready'" json:"status" binding:"required"`
	Price         float64      `gorm:"type:double; not null" json:"price" binding:"required"`
	Transmission  Transmission `gorm:"type:enum('Manual','Matic');unique ;not null" json:"transmission" binding:"required"`
	Spesification []string     `gorm:"type:varchar(255)[]" json:"Spesification"`
	IdVendor      int64        `json:"id_vendor" binding:"required"`

	Vendor Vendor `gorm:"foreignKey:id_vendor"`
}
