package ordercontroller

import (
	"encoding/json"
	"net/http"

	"github.com/balada-raja/GET/initializers"
	"github.com/balada-raja/GET/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Create(c *gin.Context) {
	var input struct {
		IdUser         int64   `json:"id_user"`
		IdVehicle      int64   `json:"id_vehicle"`
		IdVendor       int64   `json:"id_vendor"`
		Status         string  `json:"status"`
		BorrowDate     string  `json:"borrow_date"`
		ReturnDate     string  `json:"return_date"`
		BorrowDuration int     `json:"borrow_duration"`
		Total          float64 `json:"total"`
		Guarantee      string  `json:"guarantee"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	DetailOrder := models.DetailOrder{
		BorrowDate:     input.BorrowDate,
		ReturnDate:     input.ReturnDate,
		BorrowDuration: input.BorrowDuration,
		Total:          input.Total,
		Guarantee:      input.Guarantee,
	}
	if err := initializers.DB.Create(&DetailOrder).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to create detail order"})
		return
	}

	var Users models.Users
	if err := initializers.DB.First(&Users, input.IdUser).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Users not found"})
		return
	}
	var vehicle models.Vehicle
	if err := initializers.DB.First(&vehicle, input.IdVehicle).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "vehicle not found"})
		return
	}
	var vendor models.Vendor
	if err := initializers.DB.First(&vendor, input.IdVendor).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "vendor not found"})
		return
	}

	Order := models.Order{
		IdUser:        input.IdUser,
		IdVehicle:     input.IdVehicle,
		IdVendor:      input.IdVendor,
		Status:        models.Informasi(input.Status),
		IdDetailOrder: DetailOrder.Id,
	}
	if err := initializers.DB.Create(&Order).Error; err != nil {
		// jika gagal, hapus entri yang sudah dibuat dalam tabel detail_order
		initializers.DB.Delete(&DetailOrder, DetailOrder.Id)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to create order"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "Kendaraan berhasil ditambahkan"})
}

func Index(c *gin.Context) {
	var Order []models.Order

	initializers.DB.Select("id, id_user, id_vehicle, id_vendor, status").Find(&Order)

	if err := initializers.DB.Preload("DetailOrder").Find(&Order).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch orders"})
		return
	}

	// // Membuat slice baru untuk menyimpan hasil yang akan dikirimkan sebagai respons JSON
	// var orderresponse []output

	// // Mengisi slice baru dengan data yang sesuai
	// for _, o := range Order {
	// 	orderresponse = append(orderresponse, output{
	// 		Id:             o.Id,
	// 		IdUser:         o.IdUser,
	// 		IdKendaraan:    o.IdKendaraan,
	// 		IdPenyediaJasa: o.IdPenyediaJasa,
	// 		Status:         string(o.Status),
	// 		IdDetailOrder:  o.IdDetailOrder,
	// 	})
	// }

	c.JSON(http.StatusOK, gin.H{"Message": Order})
}

func Show(c *gin.Context) {
	var Order []models.Order
	id := c.Param("id")

	if err := initializers.DB.Preload("DetailOrder").First(&Order, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"Order": Order})
}

func Update(c *gin.Context) {
	id := c.Param("id")

	var input struct {
		IdUser         *int64   `json:"id_user"`
		IdVehicle      *int64   `json:"id_vehicle"`
		IdVendor       *int64   `json:"id_vendor"`
		Status         *string  `json:"status"`
		BorrowDate     *string  `json:"borrow_date"`
		ReturnDate     *string  `json:"return_date"`
		BorrowDuration *int     `json:"borrow_duration"`
		Total          *float64 `json:"total"`
		Guarantee      *string  `json:"guarantee"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Buat map untuk menyimpan nilai-nilai yang ingin diperbarui
	updateOrder := make(map[string]interface{})

	// Tambahkan hanya nilai-nilai yang tidak nil ke dalam map
	if input.IdUser != nil {
		updateOrder["id_user"] = *input.IdUser
	}
	if input.IdVehicle != nil {
		updateOrder["id_vehicle"] = *input.IdVehicle
	}
	if input.IdVendor != nil {
		updateOrder["id_vendor"] = *input.IdVendor
	}
	if input.Status != nil {
		updateOrder["status"] = *input.Status
	}

	// Perbarui hanya nilai-nilai yang telah ditetapkan
	if err := initializers.DB.Model(&models.Order{}).Where("id = ?", id).Updates(updateOrder).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal menyimpan perubahan Order"})
		return
	}

	// Buat map untuk menyimpan nilai-nilai yang ingin diperbarui
	updateDetail := make(map[string]interface{})

	// Tambahkan hanya nilai-nilai yang tidak nil ke dalam map
	if input.BorrowDate != nil {
		updateDetail["borrow_date"] = *input.BorrowDate
	}
	if input.ReturnDate != nil {
		updateDetail["return_date"] = *input.ReturnDate
	}
	if input.BorrowDuration != nil {
		updateDetail["borrow_duration"] = *input.BorrowDuration
	}
	if input.Total != nil {
		updateDetail["total"] = *input.Total
	}
	if input.Guarantee != nil {
		updateDetail["guaranetee"] = *input.Guarantee
	}

	if err := initializers.DB.Model(&models.DetailOrder{}).Where("id = ?", id).Updates(updateDetail).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal menyimpan perubahan Detail Order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})
}

func Delete(c *gin.Context) {
	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var order models.Order
	if err := initializers.DB.First(&order, input.Id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Order not found"})
		return
	}

	var detailOrder models.DetailOrder
	if err := initializers.DB.First(&detailOrder, input.Id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Detail order not found"})
		return
	}

	if err := initializers.DB.Delete(&order).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete order"})
		return
	}

	if err := initializers.DB.Delete(&detailOrder).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete detail order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
