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

func CreateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	id := c.MustGet("id")

	socialMedia := models.SocialMedia{}
	if AppJson == helpers.GetContentType(c) {
		c.ShouldBindJSON(&socialMedia)
	} else {
		c.ShouldBind(&socialMedia)
	}

	socialMedia.UserID = id.(uint)

	err := db.Debug().Create(&socialMedia).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":               socialMedia.ID,
		"name":             socialMedia.Name,
		"social_media_url": socialMedia.SosialMediaURL,
		"user_id":          socialMedia.UserID,
		"created_at":       socialMedia.CreatedAt,
	})
}

func GetSocialMedia(c *gin.Context) {
	db := database.GetDB()

	var (
		socialMedia []models.SocialMedia
	)

	err := db.Debug().Preload("User", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("username", "id")
	}).Find(&socialMedia).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	if len(socialMedia) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"msg": "There is no data in database",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"social_medias": socialMedia,
		})
	}
}

func UpdateSocialMediaByID(c *gin.Context) {
	db := database.GetDB()
	socialMediaID := c.Param("socialMediaID")

	var socialMedia models.SocialMedia

	if AppJson == helpers.GetContentType(c) {
		c.ShouldBindJSON(&socialMedia)
	} else {
		c.ShouldBind(&socialMedia)
	}

	err := db.Debug().Model(&socialMedia).Where("id = ?", socialMediaID).Updates(models.SocialMedia{
		Name:           socialMedia.Name,
		SosialMediaURL: socialMedia.SosialMediaURL,
	}).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	err = db.First(&socialMedia).Where("id = ?", socialMediaID).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":               socialMedia.ID,
		"name":             socialMedia.Name,
		"social_media_url": socialMedia.SosialMediaURL,
		"user_id":          socialMedia.UserID,
		"updated_at":       socialMedia.UpdatedAt,
	})
}

func DeleteSocialMediaByID(c *gin.Context) {
	db := database.GetDB()
	socialMediaID := c.Param("socialMediaID")

	var socialMedia models.SocialMedia

	err := db.Debug().Model(&socialMedia).Where("id = ?", socialMediaID).Delete(socialMedia).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your Sosial Media has been successfully deleted",
	})
}
