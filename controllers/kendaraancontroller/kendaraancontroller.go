package kendaraancontroller

import (
	"encoding/json"
	"net/http"

	"github.com/balada-raja/GET/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type output struct {
	Id             int64   `json:"id"`
	Nama           string  `json:"nama"`
	JenisKendaraan string  `json:"jenis_kendaraan"`
	NomorPolisi    string  `json:"nomor_polisi"`
	NomorMesin     string  `json:"nomor_mesin"`
	Deskripsi      string  `json:"deskripsi"`
	Status         string  `json:"status"`
	Harga          float64 `json:"harga"`
	IdPenyediaJasa int64   `json:"id_penyedia_jasa"`
}

func Create(c *gin.Context) {
	var input struct {
		Nama           string  `json:"nama"`
		JenisKendaraan string  `json:"jenis_kendaraan"`
		NomorPolisi    string  `json:"nomor_polisi"`
		NomorMesin     string  `json:"nomor_mesin"`
		Deskripsi      string  `json:"deskripsi"`
		Status         string  `json:"status"`
		Harga          float64 `json:"harga"`
		IdPenyediaJasa int64   `json:"id_penyedia_jasa"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var PenyediaJasa models.PenyediaJasa
	if err := models.DB.First(&PenyediaJasa, input.IdPenyediaJasa).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "penyedia jasa not found"})
		return
	}

	Kendaraan := models.Kendaraan{
		Nama:           input.Nama,
		JenisKendaraan: models.JenisKendaraan(input.JenisKendaraan),
		NomorPolisi:    input.NomorPolisi,
		NomorMesin:     input.NomorMesin,
		Deskripsi:      input.Deskripsi,
		Status:         models.Status(input.Status),
		Harga:          input.Harga,
		IdPenyediaJasa: input.IdPenyediaJasa,
	}
	models.DB.Create(&Kendaraan)
	c.JSON(http.StatusOK, gin.H{"Message": "Kendaraan berhasil ditambahkan"})
}

func Index(c *gin.Context) {
	var Kendaraan []models.Kendaraan

	models.DB.Select("id, nama, jenis_kendaraan, nomor_polisi, nomor_mesin, deskripsi, status, harga, id_penyedia_jasa").Find(&Kendaraan)

	// Membuat slice baru untuk menyimpan hasil yang akan dikirimkan sebagai respons JSON
	var kendaraanResponse []output

	// Mengisi slice baru dengan data yang sesuai
	for _, k := range Kendaraan {
		kendaraanResponse = append(kendaraanResponse, output{
			Id:             k.Id,
			Nama:           k.Nama,
			JenisKendaraan: string(k.JenisKendaraan),
			NomorPolisi:    k.NomorPolisi,
			NomorMesin:     k.NomorMesin,
			Deskripsi:      k.Deskripsi,
			Status:         string(k.Status),
			Harga:          k.Harga,
			IdPenyediaJasa: k.IdPenyediaJasa,
		})
	}

	c.JSON(http.StatusOK, gin.H{"Message": kendaraanResponse})
}

func Show(c *gin.Context) {
	var Kendaraan models.Kendaraan
	id := c.Param("id")

	if err := models.DB.Select("id, nama, jenis_kendaraan").First(&Kendaraan, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	// Konversi data ke dalam KendaraanResponse
	kendaraanResponse := output{
		Id:             Kendaraan.Id,
		Nama:           Kendaraan.Nama,
		JenisKendaraan: string(Kendaraan.JenisKendaraan),
		NomorPolisi:    Kendaraan.NomorPolisi,
		NomorMesin:     Kendaraan.NomorMesin,
		Deskripsi:      Kendaraan.Deskripsi,
		Status:         string(Kendaraan.Status),
		Harga:          Kendaraan.Harga,
		IdPenyediaJasa: Kendaraan.IdPenyediaJasa,
	}

	c.JSON(http.StatusOK, gin.H{"Kendaraan": kendaraanResponse})
}

func Update(c *gin.Context) {
	id := c.Param("id")

	var input struct {
		Nama           *string  `json:"nama"`
		JenisKendaraan *string  `json:"jenis_kendaraan"`
		NomorPolisi    *string  `json:"nomor_polisi"`
		NomorMesin     *string  `json:"nomor_mesin"`
		Deskripsi      *string  `json:"deskripsi"`
		Status         *string  `json:"status"`
		Harga          *float64 `json:"harga"`
		IdPenyediaJasa *int64   `json:"id_penyedia_jasa"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Buat map untuk menyimpan nilai-nilai yang ingin diperbarui
	updateValues := make(map[string]interface{})

	// Tambahkan hanya nilai-nilai yang tidak nil ke dalam map
	if input.Nama != nil {
		updateValues["nama"] = *input.Nama
	}
	if input.JenisKendaraan != nil {
		updateValues["jenis_kendaraan"] = *input.JenisKendaraan
	}
	if input.NomorPolisi != nil {
		updateValues["nomor_polisi"] = *input.NomorPolisi
	}
	if input.NomorMesin != nil {
		updateValues["nomor_mesin"] = *input.NomorMesin
	}
	if input.Status != nil {
		updateValues["status"] = *input.Status
	}
	if input.Harga != nil {
		updateValues["harga"] = *input.Harga
	}
	if input.IdPenyediaJasa != nil {
		updateValues["id_penyedia_jasa"] = *input.IdPenyediaJasa
	}
	updateValues["deskripsi"] = *input.Deskripsi

	// Perbarui hanya nilai-nilai yang telah ditetapkan
	if err := models.DB.Model(&models.Kendaraan{}).Where("id = ?", id).Updates(updateValues).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal menyimpan perubahan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})
}

func Delete(c *gin.Context) {
	var Kendaraan models.Kendaraan

	var input struct {
		Id json.Number
	}

	//input := map[string]string{"Id": "0"}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&Kendaraan, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus Kendaraan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
