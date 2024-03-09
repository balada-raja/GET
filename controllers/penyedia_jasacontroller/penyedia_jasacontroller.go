package penyedia_jasacontroller

import (
	"encoding/json"
	"net/http"

	"github.com/balada-raja/GET/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Create(c *gin.Context) {
	var PenyediaJasa models.PenyediaJasa

	if err := c.ShouldBindJSON(&PenyediaJasa); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	models.DB.Create(&PenyediaJasa)
	c.JSON(http.StatusOK, gin.H{"penyedia_jasa": PenyediaJasa})

}

func Index(c *gin.Context) {
	var PenyediaJasa []models.PenyediaJasa

	models.DB.Find(&PenyediaJasa)
	c.JSON(http.StatusOK, gin.H{"penyedia_jasa": PenyediaJasa})
}

func Show(c *gin.Context) {
	var PenyediaJasa []models.PenyediaJasa
	id := c.Param("id")

	if err := models.DB.First(&PenyediaJasa, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"penyedia_jasa": PenyediaJasa})
}

func Update(c *gin.Context) {
	var PenyediaJasa models.PenyediaJasa
	id := c.Param("id")
	if err := c.ShouldBindJSON(&PenyediaJasa); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&PenyediaJasa).Where("id = ?", id).Updates(&PenyediaJasa).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupadate penyedia jasa"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})
}

func Delete(c *gin.Context) {
	var PenyediaJasa models.PenyediaJasa

	var input struct {
		Id json.Number
	}

	//input := map[string]string{"Id": "0"}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&PenyediaJasa, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus penyedia jasa"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
