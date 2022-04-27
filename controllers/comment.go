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

func CreateComment(c *gin.Context) {
	db := database.GetDB()
	id := c.MustGet("id")

	comment := models.Comment{}
	if AppJson == helpers.GetContentType(c) {
		c.ShouldBindJSON(&comment)
	} else {
		c.ShouldBind(&comment)
	}

	comment.UserID = id.(uint)

	err := db.Debug().Create(&comment).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         comment.ID,
		"message":    comment.Message,
		"photo_id":   comment.PhotoID,
		"user_id":    comment.UserID,
		"created_at": comment.CreatedAt,
	})
}

func GetComments(c *gin.Context) {
	db := database.GetDB()

	var (
		comment []models.Comment
	)

	err := db.Debug().Preload("User", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("id", "email", "username")
	}).Preload("Photo", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("id", "title", "caption", "photo_url", "user_id")
	}).Find(&comment).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	if len(comment) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"msg": "There is no data in database",
		})
	} else {
		c.JSON(http.StatusOK, comment)
	}
}

func UpdateCommentByID(c *gin.Context) {
	db := database.GetDB()
	commentID := c.Param("commentID")

	var comment models.Comment

	if AppJson == helpers.GetContentType(c) {
		c.ShouldBindJSON(&comment)
	} else {
		c.ShouldBind(&comment)
	}

	err := db.Debug().Model(&comment).Where("id = ?", commentID).Updates(models.Comment{
		Message: comment.Message,
	}).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	err = db.First(&comment).Where("id = ?", commentID).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         comment.ID,
		"message":    comment.Message,
		"photo_id":   comment.PhotoID,
		"user_id":    comment.UserID,
		"created_at": comment.CreatedAt,
	})
}

func DeleteCommentByID(c *gin.Context) {
	db := database.GetDB()
	commentID := c.Param("commentID")

	var comment models.Comment

	err := db.Debug().Model(&comment).Where("id = ?", commentID).Delete(comment).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}
