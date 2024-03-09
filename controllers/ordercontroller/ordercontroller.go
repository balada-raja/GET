package ordercontroller

import (
	"encoding/json"
	"net/http"

	"github.com/balada-raja/GET/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// type output struct {
// 	Id             int64  `json:"id"`
// 	IdUser         int64  `json:"id_user"`
// 	IdKendaraan    int64  `json:"id_kendaraan"`
// 	IdPenyediaJasa int64  `json:"id_penyedia_jasa"`
// 	Status         string `json:"status"`
// 	IdDetailOrder  int64  `json:"id_detail_order"`
// }

func Create(c *gin.Context) {
	var input struct {
		IdUser          int64   `json:"id_user"`
		IdKendaraan     int64   `json:"id_kendaraan"`
		IdPenyediaJasa  int64   `json:"id_penyedia_jasa"`
		Status          string  `json:"status"`
		TglPeminjaman   string  `json:"tgl_peminjaman"`
		TglPengembalian string  `json:"tgl_pengembalian"`
		DurasiSewa      int     `json:"durasi_sewa"`
		Total           float64 `json:"total"`
		Jaminan         string  `json:"jaminan"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	DetailOrder := models.DetailOrder{
		TglPeminjaman:   input.TglPeminjaman,
		TglPengembalian: input.TglPengembalian,
		DurasiSewa:      input.DurasiSewa,
		Total:           input.Total,
		Jaminan:         input.Jaminan,
	}
	if err := models.DB.Create(&DetailOrder).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to create detail order"})
		return
	}

	var Users models.Users
	if err := models.DB.First(&Users, input.IdUser).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Users not found"})
		return
	}
	var Kendaraan models.Kendaraan
	if err := models.DB.First(&Kendaraan, input.IdKendaraan).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Kendaraan not found"})
		return
	}
	var PenyediaJasa models.PenyediaJasa
	if err := models.DB.First(&PenyediaJasa, input.IdPenyediaJasa).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Peneyedia jasa not found"})
		return
	}

	Order := models.Order{
		IdUser:         input.IdUser,
		IdKendaraan:    input.IdKendaraan,
		IdPenyediaJasa: input.IdPenyediaJasa,
		Status:         models.Informasi(input.Status),
		IdDetailOrder:  DetailOrder.Id,
	}
	if err := models.DB.Create(&Order).Error; err != nil {
		// jika gagal, hapus entri yang sudah dibuat dalam tabel detail_order
		models.DB.Delete(&DetailOrder, DetailOrder.Id)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to create order"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "Kendaraan berhasil ditambahkan"})
}

func Index(c *gin.Context) {
	var Order []models.Order

	models.DB.Select("id, id_user, id_kendaraan, id_penyedia_jasa, status").Find(&Order)

	if err := models.DB.Preload("DetailOrder").Find(&Order).Error; err != nil {
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

	c.JSON(http.StatusOK, gin.H{"Message": Order,})
}

func Show(c *gin.Context) {
	var Order []models.Order
	id := c.Param("id")

	if err := models.DB.Preload("DetailOrder").First(&Order, id).Error; err != nil {
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
		IdUser          *int64   `json:"id_user"`
		IdKendaraan     *int64   `json:"id_kendaraan"`
		IdPenyediaJasa  *int64   `json:"id_penyedia_jasa"`
		Status          *string  `json:"status"`
		TglPeminjaman   *string  `json:"tgl_peminjaman"`
		TglPengembalian *string  `json:"tgl_pengembalian"`
		DurasiSewa      *int     `json:"durasi_sewa"`
		Total           *float64 `json:"total"`
		Jaminan         *string  `json:"jaminan"`
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
	if input.IdKendaraan != nil {
		updateOrder["id_kendaraan"] = *input.IdKendaraan
	}
	if input.IdPenyediaJasa != nil {
		updateOrder["id_penyedia_jasa"] = *input.IdPenyediaJasa
	}
	if input.Status != nil {
		updateOrder["status"] = *input.Status
	}

	// Perbarui hanya nilai-nilai yang telah ditetapkan
	if err := models.DB.Model(&models.Order{}).Where("id = ?", id).Updates(updateOrder).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal menyimpan perubahan Order"})
		return
	}

	// Buat map untuk menyimpan nilai-nilai yang ingin diperbarui
	updateDetail := make(map[string]interface{})

	// Tambahkan hanya nilai-nilai yang tidak nil ke dalam map
	if input.TglPeminjaman != nil {
		updateDetail["tgl_peminjaman"] = *input.TglPeminjaman
	}
	if input.TglPengembalian != nil {
		updateDetail["tgl_pengambalian"] = *input.TglPengembalian
	}
	if input.DurasiSewa != nil {
		updateDetail["durasi_sewa"] = *input.DurasiSewa
	}
	if input.Total != nil {
		updateDetail["total"] = *input.Total
	}
	if input.Jaminan != nil {
		updateDetail["jaminan"] = *input.Jaminan
	}

	if err := models.DB.Model(&models.DetailOrder{}).Where("id = ?", id).Updates(updateDetail).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal menyimpan perubahan Detail Order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})
}

func Delete(c *gin.Context) {
	var input struct{
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var order models.Order
	if err := models.DB.First(&order, input.Id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Order not found"})
		return
	}

	var detailOrder models.DetailOrder
	if err := models.DB.First(&detailOrder, input.Id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Detail order not found"})
		return
	}

	if err := models.DB.Delete(&order).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete order"})
		return
	}

	if err := models.DB.Delete(&detailOrder).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete detail order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}

