package vendorcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/balada-raja/GET/models"
	"github.com/balada-raja/GET/initializers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Create(c *gin.Context) {
	var vendor models.Vendor

	if err := c.ShouldBindJSON(&vendor); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	initializers.DB.Create(&vendor)
	c.JSON(http.StatusOK, gin.H{"vendor": vendor})

}

func Index(c *gin.Context) {
	var vendor []models.Vendor

	initializers.DB.Find(&vendor)
	c.JSON(http.StatusOK, gin.H{"vendor": vendor})
}

func Show(c *gin.Context) {
	var vendor []models.Vendor
	id := c.Param("id")

	if err := initializers.DB.First(&vendor, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"vendor": vendor})
}

func Update(c *gin.Context) {
	var vendor models.Vendor
	id := c.Param("id")
	if err := c.ShouldBindJSON(&vendor); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if initializers.DB.Model(&vendor).Where("id = ?", id).Updates(&vendor).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupadate data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})
}

func Delete(c *gin.Context) {
	var vendor models.Vendor

	var input struct {
		Id json.Number
	}

	//input := map[string]string{"Id": "0"}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if initializers.DB.Delete(&vendor, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
