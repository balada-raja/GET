package vehiclecontroller

import (
	"encoding/json"
	"net/http"

	"github.com/balada-raja/GET/initializers"
	"github.com/balada-raja/GET/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type output struct {
	Id            int64    `json:"id"`
	Name          string   `json:"name"`
	VehicleType   string   `json:"vehicle_type"`
	PoliceNumber  string   `json:"police_number"`
	MachineNumber string   `json:"machine_number"`
	Description   string   `json:"description"`
	Status        string   `json:"status"`
	Price         float64  `json:"price"`
	Transmission  string   `json:"transmission"`
	Spesification []string `json:"spesification"`
	IdVendor      int64    `json:"id_vendor"`
}

func Create(c *gin.Context) {
	var input struct {
		Name          string   `json:"name" binding:"required"`
		VehicleType   string   `json:"vehicle_type" binding:"required"`
		PoliceNumber  string   `json:"police_number" binding:"required"`
		MachineNumber string   `json:"machine_number" binding:"required"`
		Description   string   `json:"description"`
		Status        string   `json:"status" binding:"required"`
		Price         float64  `json:"price" binding:"required"`
		Transmission  string   `json:"transmission"`
		Spesification []string `json:"spesification"`
		IdVendor      int64    `json:"id_vendor" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var vendor models.Vendor
	if err := initializers.DB.First(&vendor, input.IdVendor).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "penyedia jasa not found"})
		return
	}

	vehicle := models.Vehicle{
		Name:          input.Name,
		VehicleType:   models.VehicleType(input.VehicleType),
		PoliceNumber:  input.PoliceNumber,
		MachineNumber: input.MachineNumber,
		Description:   input.Description,
		Status:        models.Status(input.Status),
		Price:         input.Price,
		Transmission:  models.Transmission(input.Transmission),
		Spesification: input.Spesification,
		IdVendor:      input.IdVendor,
	}
	initializers.DB.Create(&vehicle)
	c.JSON(http.StatusOK, gin.H{"Message": vehicle})
}

func Index(c *gin.Context) {
	var vehicle []models.Vehicle

	initializers.DB.Select("id, name, vehicle_type, police_number, machine_number, description, status, price, transmission, id_vendor").Find(&vehicle)

	// Membuat slice baru untuk menyimpan hasil yang akan dikirimkan sebagai respons JSON
	var VehicleResponse []output

	// Mengisi slice baru dengan data yang sesuai
	for _, k := range vehicle {
		VehicleResponse = append(VehicleResponse, output{
			Id:            k.Id,
			Name:          k.Name,
			VehicleType:   string(k.VehicleType),
			PoliceNumber:  k.PoliceNumber,
			MachineNumber: k.MachineNumber,
			Description:   k.Description,
			Status:        string(k.Status),
			Price:         k.Price,
			Transmission:  string(k.Transmission),
			Spesification: k.Spesification,
			IdVendor:      k.IdVendor,
		})
	}

	c.JSON(http.StatusOK, gin.H{"Message": VehicleResponse})
}

func Show(c *gin.Context) {
	var vehicle models.Vehicle
	id := c.Param("id")

	if err := initializers.DB.Select("id, name, vehicle_type, police_number, machine_number, description, status, price, transmission, spesification, id_vendor").First(&vehicle, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	// Konversi data ke dalam vehicleResponse
	vehicleResponse := output{
		Id:            vehicle.Id,
		Name:          vehicle.Name,
		VehicleType:   string(vehicle.VehicleType),
		PoliceNumber:  vehicle.PoliceNumber,
		MachineNumber: vehicle.MachineNumber,
		Description:   vehicle.Description,
		Status:        string(vehicle.Status),
		Price:         vehicle.Price,
		Transmission:  string(vehicle.Transmission),
		Spesification: vehicle.Spesification,
		IdVendor:      vehicle.IdVendor,
	}

	c.JSON(http.StatusOK, gin.H{"vehicle": vehicleResponse})
}

func Update(c *gin.Context) {
	id := c.Param("id")

	var input struct {
		Name          *string   `json:"name"`
		VehicleType   *string   `json:"vehicle_type"`
		PoliceNumber  *string   `json:"police_number"`
		MachineNumber *string   `json:"machine_number"`
		Description   *string   `json:"description"`
		Status        *string   `json:"status"`
		Price         *float64  `json:"price"`
		Transmission  *string   `json:"transmission"`
		Spesification *[]string `json:"spesification"`
		IdVendor      *int64    `json:"id_vendor"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Buat map untuk menyimpan nilai-nilai yang ingin diperbarui
	updateValues := make(map[string]interface{})

	// Tambahkan hanya nilai-nilai yang tidak nil ke dalam map
	if input.Name != nil {
		updateValues["name"] = *input.Name
	}
	if input.VehicleType != nil {
		updateValues["vehicle_type"] = *input.VehicleType
	}
	if input.PoliceNumber != nil {
		updateValues["police_number"] = *input.PoliceNumber
	}
	if input.MachineNumber != nil {
		updateValues["machine_number"] = *input.MachineNumber
	}
	if input.Status != nil {
		updateValues["status"] = *input.Status
	}
	if input.Price != nil {
		updateValues["harga"] = *input.Price
	}
	if input.IdVendor != nil {
		updateValues["id_vendor"] = *input.IdVendor
	}
	updateValues["transmission"] = *input.Transmission
	updateValues["description"] = *input.Description

	// Perbarui hanya nilai-nilai yang telah ditetapkan
	if err := initializers.DB.Model(&models.Vehicle{}).Where("id = ?", id).Updates(updateValues).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal menyimpan perubahan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})
}

func Delete(c *gin.Context) {
	var vehicle models.Vehicle

	var input struct {
		Id json.Number
	}

	//input := map[string]string{"Id": "0"}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if initializers.DB.Delete(&vehicle, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}

func FindVehicleByName(c *gin.Context) {
	var vehicles []models.Vehicle
	query := c.Query("query")

	// Melakukan pencarian kendaraan berdasarkan query
	if err := initializers.DB.Where("name = ?", query).Find(&vehicles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var VehicleResponse []output

	for _, k := range vehicles {
		VehicleResponse = append(VehicleResponse, output{
			Id:            k.Id,
			Name:          k.Name,
			VehicleType:   string(k.VehicleType),
			PoliceNumber:  k.PoliceNumber,
			MachineNumber: k.MachineNumber,
			Description:   k.Description,
			Status:        string(k.Status),
			Price:         k.Price,
			Transmission:  string(k.Transmission),
			Spesification: k.Spesification,
			IdVendor:      k.IdVendor,
		})
	}

	c.JSON(http.StatusOK, VehicleResponse)
}

func FindVehicleByVehicleType(c *gin.Context) {
	var vehicles []models.Vehicle
	query := c.Query("query")

	// Melakukan pencarian kendaraan berdasarkan query
	if err := initializers.DB.Where("vehicle_type = ?", query).Find(&vehicles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var VehicleResponse []output

	for _, k := range vehicles {
		VehicleResponse = append(VehicleResponse, output{
			Id:            k.Id,
			Name:          k.Name,
			VehicleType:   string(k.VehicleType),
			PoliceNumber:  k.PoliceNumber,
			MachineNumber: k.MachineNumber,
			Description:   k.Description,
			Status:        string(k.Status),
			Price:         k.Price,
			Transmission:  string(k.Transmission),
			Spesification: k.Spesification,
			IdVendor:      k.IdVendor,
		})
	}

	c.JSON(http.StatusOK, VehicleResponse)
}

func FindVehicleByTransmission(c *gin.Context) {
	var vehicles []models.Vehicle
	query := c.Query("query")

	// Melakukan pencarian kendaraan berdasarkan query
	if err := initializers.DB.Where("transmission = ?", query).Find(&vehicles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var VehicleResponse []output

	for _, k := range vehicles {
		VehicleResponse = append(VehicleResponse, output{
			Id:            k.Id,
			Name:          k.Name,
			VehicleType:   string(k.VehicleType),
			PoliceNumber:  k.PoliceNumber,
			MachineNumber: k.MachineNumber,
			Description:   k.Description,
			Status:        string(k.Status),
			Price:         k.Price,
			Transmission:  string(k.Transmission),
			Spesification: k.Spesification,
			IdVendor:      k.IdVendor,
		})
	}

	c.JSON(http.StatusOK, VehicleResponse)
}
