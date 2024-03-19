package helper

import (
	"github.com/balada-raja/GET/models"
)

type Output struct {
	Id            int64   `json:"id"`
	Name          string  `json:"name"`
	VehicleType   string  `json:"vehicle_type"`
	PoliceNumber  string  `json:"police_number"`
	MachineNumber string  `json:"machine_number"`
	Description   string  `json:"description"`
	Status        string  `json:"status"`
	Price         float64 `json:"price"`
	Transmission  string  `json:"transmission"`
	IdVendor      int64   `json:"id_vendor"`
}

func ResponseOutput(vehicle []models.Vehicle) []Output {
	var VehicleResponse []Output

	for _, k := range vehicle {
		VehicleResponse = append(VehicleResponse, Output{
			Id:            k.Id,
			Name:          k.Name,
			VehicleType:   string(k.VehicleType),
			PoliceNumber:  k.PoliceNumber,
			MachineNumber: k.MachineNumber,
			Description:   k.Description,
			Status:        string(k.Status),
			Price:         k.Price,
			Transmission:  string(k.Transmission),
			IdVendor:      k.IdVendor,
		})
	}

	return VehicleResponse
}

func ShowOutput(vehicle models.Vehicle) []Output {
	var VehicleResponse []Output

	VehicleResponse = append(VehicleResponse, Output{
		Id:            vehicle.Id,
		Name:          vehicle.Name,
		VehicleType:   string(vehicle.VehicleType),
		PoliceNumber:  vehicle.PoliceNumber,
		MachineNumber: vehicle.MachineNumber,
		Description:   vehicle.Description,
		Status:        string(vehicle.Status),
		Price:         vehicle.Price,
		Transmission:  string(vehicle.Transmission),
		IdVendor:      vehicle.IdVendor,
	})

	return VehicleResponse
}
