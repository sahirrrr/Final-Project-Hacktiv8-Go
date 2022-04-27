package controllers

import (
	"final-project/database"
	"final-project/helpers"
	"final-project/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var AppJson = "application/json"

func UserRegister(c *gin.Context) {
	db := database.GetDB()

	user := models.User{}
	if AppJson == helpers.GetContentType(c) {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	err := db.Debug().Create(&user).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       user.ID,
		"email":    user.Email,
		"username": user.Username,
		"age":      user.Age,
	})
}

func UserLogin(c *gin.Context) {
	db := database.GetDB()

	user := models.User{}
	if AppJson == helpers.GetContentType(c) {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	password := user.Password

	err := db.Debug().Where("email = ?", user.Email).Take(&user).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "NOT FOUND",
			"message": err.Error(),
		})
		return
	}

	if err := helpers.ComparePass(user.Password, password); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "UNAUTHORIZED",
			"message": err.Error(),
		})
		return
	}

	token := helpers.GenerateToken(user.ID, user.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func UpdateUserByID(c *gin.Context) {
	db := database.GetDB()
	userID := c.Param("userID")

	var user models.User

	if AppJson == helpers.GetContentType(c) {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	err := db.Debug().Model(&user).Where("id = ?", userID).Updates(models.User{
		Email:    user.Email,
		Username: user.Username,
	}).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	err = db.Debug().Where("id = ?", userID).Take(&user).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         user.ID,
		"email":      user.Email,
		"username":   user.Username,
		"age":        user.Age,
		"updated_at": user.UpdatedAt,
	})
}

func DeleteUserByID(c *gin.Context) {
	db := database.GetDB()
	userID := c.Param("userID")

	var user models.User

	err := db.Debug().Model(&user).Where("id = ?", userID).Delete(user).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}
