package infopelanggancontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
	"rest-api/models"
)

func Index(c *gin.Context) {
	var info_pelanggan []models.Info_pelanggan

	models.DB.Find(&info_pelanggan)

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Data ditemukan",
		"data":    info_pelanggan,
	})
}

func Show(c *gin.Context) {
	var info_pelanggan models.Info_pelanggan
	id := c.Param("id")

	if err := models.DB.First(&info_pelanggan, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.JSON(http.StatusOK, gin.H{
				"status":  false,
				"message": "Data tidak ditemukan",
			})
			return
		default:
			c.JSON(http.StatusOK, gin.H{
				"status":  false,
				"message": err.Error(),
			})
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Data ditemukan",
		"data":    info_pelanggan,
	})
}

func Create(c *gin.Context) {
	var info_pelanggan models.Info_pelanggan

	if err := c.ShouldBindJSON(&info_pelanggan); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Gagal membuat data",
			"data":    err.Error(),
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(info_pelanggan); err != nil {
		errors := err.(validator.ValidationErrors)
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Tidak dapat menambahkan data",
			"data":    errors.Error(),
		})
		return
	}

	if err := models.DB.Create(&info_pelanggan).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Gagal memasukkan data",
			"data":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Sukses menambahkan data",
	})
}

func Update(c *gin.Context) {
	var info_pelanggan models.Info_pelanggan
	id := c.Param("id")

	if err := c.ShouldBindJSON(&info_pelanggan); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Gagal melakukan update data",
			"data":    err.Error(),
		})
		return
	}

	if models.DB.Model(&info_pelanggan).Where("id_plg = ?", id).Updates(&info_pelanggan).RowsAffected == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Tidak ada perubahan",
			"data":    "Tidak ada perubahan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Sukses melakukan update data",
		"data":    info_pelanggan,
	})
}
func Delete(c *gin.Context) {
	var info_pelanggan models.Info_pelanggan

	id := c.Param("id")

	if models.DB.Delete(&info_pelanggan, id).RowsAffected == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Tidak dapat menghapus data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Sukses melakukan delete data",
	})
}
