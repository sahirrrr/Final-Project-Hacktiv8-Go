package middlewares

import (
	"final-project/database"
	"final-project/helpers"
	"final-project/models"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		headerToken := c.Request.Header.Get("Authorization")
		bearer := strings.HasPrefix(headerToken, "Bearer")

		if !bearer {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"error":  "UNAUTHORIZED",
			})
			return
		}

		token := strings.Split(headerToken, " ")[1]

		id, email, err := helpers.ValidateToken(token)
		if err != nil {
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"error":  "UNAUTHORIZED",
				"msg":    err.Error(),
			})
			return
		}

		c.Set("id", id)
		c.Set("email", email)
		c.Next()
	}
}

func UserAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		id := c.MustGet("id").(uint)
		userID := c.Param("userID")

		var user models.User

		err := db.Debug().Where("id = ?", userID).Take(&user).Error
		if err != nil {
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status": http.StatusNotFound,
				"error":  "USER NOT FOUND",
				"msg":    err.Error(),
			})
			return
		}

		if user.ID != id {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"status": http.StatusForbidden,
				"error":  "FORBIDDEN",
				"msg":    "you not have right access to this resources",
			})
			return
		}

		c.Next()
	}
}

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		id := c.MustGet("id").(uint)
		photoID := c.Param("photoID")

		var photo models.Photo

		err := db.Debug().Where("id = ?", photoID).Take(&photo).Error
		if err != nil {
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status": http.StatusNotFound,
				"error":  "PHOTO NOT FOUND",
				"msg":    err.Error(),
			})
			return
		}

		if photo.UserID != id {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"status": http.StatusForbidden,
				"error":  "FORBIDDEN",
				"msg":    "you not have right access to this resources",
			})
			return
		}

		c.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		id := c.MustGet("id").(uint)
		commentID := c.Param("commentID")

		var comment models.Comment

		err := db.Debug().Where("id = ?", commentID).Take(&comment).Error
		if err != nil {
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status": http.StatusNotFound,
				"error":  "PHOTO NOT FOUND",
				"msg":    err.Error(),
			})
			return
		}

		if comment.UserID != id {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"status": http.StatusForbidden,
				"error":  "FORBIDDEN",
				"msg":    "you not have right access to this resources",
			})
			return
		}

		c.Next()
	}
}

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		id := c.MustGet("id").(uint)
		socialMediaID := c.Param("socialMediaID")

		var socialMedia models.SocialMedia

		err := db.Debug().Where("id = ?", socialMediaID).Take(&socialMedia).Error
		if err != nil {
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status": http.StatusNotFound,
				"error":  "PHOTO NOT FOUND",
				"msg":    err.Error(),
			})
			return
		}

		if socialMedia.UserID != id {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"status": http.StatusForbidden,
				"error":  "FORBIDDEN",
				"msg":    "you not have right access to this resources",
			})
			return
		}

		c.Next()
	}
}
