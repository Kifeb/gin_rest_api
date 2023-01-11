package productcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kifeb/gin_rest_api/models"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var prodcts []models.ProducModel

	models.DB.Find(&prodcts)

	c.JSON(http.StatusOK, gin.H{
		"products": prodcts,
	})
}

func Show(c *gin.Context) {
	var product models.ProducModel
	id := c.Param("id")

	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"msg": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"product": product})
}

func Create(c *gin.Context) {
	var product models.ProducModel

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	models.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"product": product})
}

func Update(c *gin.Context) {
	var product models.ProducModel
	id := c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	if models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "Data tidak bisa di update"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "data berhasil diperbaharui"})
}
