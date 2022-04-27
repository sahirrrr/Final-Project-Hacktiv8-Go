package controllers

import (
	"final-project/database"
	"final-project/helpers"
	"final-project/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	id := c.MustGet("id")

	photo := models.Photo{}
	if AppJson == helpers.GetContentType(c) {
		c.ShouldBindJSON(&photo)
	} else {
		c.ShouldBind(&photo)
	}

	photo.UserID = id.(uint)

	err := db.Debug().Create(&photo).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         photo.ID,
		"title":      photo.Title,
		"caption":    photo.Caption,
		"photo_url":  photo.PhotoUrl,
		"user_id":    photo.UserID,
		"created_at": photo.CreatedAt,
	})
}

func GetPhotos(c *gin.Context) {
	db := database.GetDB()

	var (
		photo []models.Photo
	)

	err := db.Debug().Preload("User", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("email", "username", "id")
	}).Find(&photo).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	if len(photo) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"msg": "There is no data in database",
		})
	} else {
		c.JSON(http.StatusOK, photo)
	}
}

func UpdatePhotoByID(c *gin.Context) {
	db := database.GetDB()
	photoID := c.Param("photoID")

	var photo models.Photo

	if AppJson == helpers.GetContentType(c) {
		c.ShouldBindJSON(&photo)
	} else {
		c.ShouldBind(&photo)
	}

	err := db.Debug().Model(&photo).Where("id = ?", photoID).Updates(models.Photo{
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoUrl: photo.PhotoUrl,
	}).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	err = db.First(&photo).Where("id = ?", photoID).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         photo.ID,
		"title":      photo.Title,
		"caption":    photo.Caption,
		"photo_url":  photo.PhotoUrl,
		"user_id":    photo.UserID,
		"updated_at": photo.UpdatedAt,
	})
}

func DeletePhotoByID(c *gin.Context) {
	db := database.GetDB()
	photoID := c.Param("photoID")

	var photo models.Photo

	err := db.Debug().Model(&photo).Where("id = ?", photoID).Delete(photo).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
