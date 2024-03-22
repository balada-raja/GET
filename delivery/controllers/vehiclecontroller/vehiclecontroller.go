package vehiclecontroller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/balada-raja/GET/delivery/helper"
	"github.com/balada-raja/GET/models"
	"github.com/balada-raja/GET/repository/initializers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Create(c *gin.Context) {
	var input struct {
		Name          string  `json:"name" binding:"required"`
		VehicleType   string  `json:"vehicle_type" binding:"required"`
		PoliceNumber  string  `json:"police_number" binding:"required"`
		MachineNumber string  `json:"machine_number" binding:"required"`
		Description   string  `json:"description"`
		Status        string  `json:"status" binding:"required"`
		Price         float64 `json:"price" binding:"required"`
		Transmission  string  `json:"transmission"`
		// Specification json.RawMessage `json:"specifications"`
		IdVendor int64 `json:"id_vendor" binding:"required"`
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

	// // Get file from request
	// file, err := c.FormFile("image")
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to upload file"})
	// 	return
	// }

	// // Validate file type
	// if !helper.IsImageFile(file) {
	// 	c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid file type. Please upload an image file"})
	// 	return
	// }

	// // Save uploaded file
	// imagePath, err := helper.SaveUploadedFile(file)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save uploaded file"})
	// 	return
	// }

	// var specifications []string
	// if err := json.Unmarshal(input.Specification, &specifications); err != nil {
	// 	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	// 	return
	// }

	vehicle := models.Vehicle{
		Name:          input.Name,
		VehicleType:   models.VehicleType(input.VehicleType),
		PoliceNumber:  input.PoliceNumber,
		MachineNumber: input.MachineNumber,
		Description:   input.Description,
		Status:        models.Status(input.Status),
		Price:         input.Price,
		Transmission:  models.Transmission(input.Transmission),
		// Specifications: specifications,
		//Image:    imagePath,
		IdVendor: input.IdVendor,
	}
	if err := initializers.DB.Create(&vehicle).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"vehicle": "success"})
}

func Index(c *gin.Context) {
	var vehicle []models.Vehicle

	initializers.DB.Select("id, name, vehicle_type, police_number, machine_number, description, status, price, transmission, id_vendor").Find(&vehicle)

	VehicleResponse := helper.ResponseOutput(vehicle)

	c.JSON(http.StatusOK, gin.H{"Message": VehicleResponse})
}

func Show(c *gin.Context) {
	var vehicle models.Vehicle
	id := c.Param("id")

	if err := initializers.DB.Select("id, name, vehicle_type, police_number, machine_number, description, status, price, transmission, id_vendor").First(&vehicle, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	VehicleResponse := helper.ShowOutput(vehicle)

	c.JSON(http.StatusOK, gin.H{"vehicle": VehicleResponse})
}

func Update(c *gin.Context) {
	id := c.Param("id")

	var input struct {
		Name          *string  `json:"name"`
		VehicleType   *string  `json:"vehicle_type"`
		PoliceNumber  *string  `json:"police_number"`
		MachineNumber *string  `json:"machine_number"`
		Description   *string  `json:"description"`
		Status        *string  `json:"status"`
		Price         *float64 `json:"price"`
		Transmission  *string  `json:"transmission"`
		// Specification *[]string `json:"specification"`
		IdVendor *int64 `json:"id_vendor"`
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
		updateValues["price"] = *input.Price
	}
	if input.IdVendor != nil {
		updateValues["id_vendor"] = *input.IdVendor
	}
	updateValues["transmission"] = *input.Transmission
	updateValues["description"] = *input.Description

	// Perbarui hanya nilai-nilai yang telah ditetapkan
	if err := initializers.DB.Model(&models.Vehicle{}).Where("id = ?", id).Updates(updateValues).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
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
	var vehicle []models.Vehicle
	name := c.Query("name")

	// Melakukan pencarian kendaraan berdasarkan query
	if err := initializers.DB.Where("name = ?", name).Find(&vehicle).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	VehicleResponse := helper.ResponseOutput(vehicle)

	c.JSON(http.StatusOK, VehicleResponse)
}

func FindVehicleByVehicleType(c *gin.Context) {
	var vehicle []models.Vehicle
	vehicle_type := c.Query("vehicle_type")

	// Melakukan pencarian kendaraan berdasarkan query
	if err := initializers.DB.Where("vehicle_type = ?", vehicle_type).Find(&vehicle).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	VehicleResponse := helper.ResponseOutput(vehicle)

	c.JSON(http.StatusOK, VehicleResponse)
}

func FindVehicleByTransmission(c *gin.Context) {
	var vehicle []models.Vehicle
	transmission := c.Query("transmission")

	// Melakukan pencarian kendaraan berdasarkan query
	if err := initializers.DB.Where("transmission = ?", transmission).Find(&vehicle).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	VehicleResponse := helper.ResponseOutput(vehicle)

	c.JSON(http.StatusOK, VehicleResponse)
}

func FindVehicleByPriceRange(c *gin.Context) {
	var vehicle []models.Vehicle

	// Mengambil rentang harga dari query string
	minPriceStr := c.Query("min_price")
	maxPriceStr := c.Query("max_price")

	// Mengonversi rentang harga menjadi tipe data yang sesuai
	minPrice, err := strconv.ParseFloat(minPriceStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid min_price"})
		return
	}

	maxPrice, err := strconv.ParseFloat(maxPriceStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid max_price"})
		return
	}

	if err := initializers.DB.Where("price BETWEEN ? AND ?", minPrice, maxPrice).Find(&vehicle).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	VehicleResponse := helper.ResponseOutput(vehicle)

	c.JSON(http.StatusOK, VehicleResponse)
}
