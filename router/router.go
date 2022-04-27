package router

import (
	"final-project/controllers"
	"final-project/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRoute := r.Group("/users")
	{
		userRoute.POST("/register", controllers.UserRegister)
		userRoute.POST("/login", controllers.UserLogin)

		userRoute.Use(middlewares.Auth())
		userRoute.PUT("/:userID", middlewares.UserAuthorization(), controllers.UpdateUserByID)
		userRoute.DELETE("/:userID", middlewares.UserAuthorization(), controllers.DeleteUserByID)
	}

	photoRoute := r.Group("/photos")
	{
		photoRoute.Use(middlewares.Auth())
		photoRoute.POST("/", controllers.CreatePhoto)
		photoRoute.GET("/", controllers.GetPhotos)
		photoRoute.PUT("/:photoID", middlewares.PhotoAuthorization(), controllers.UpdatePhotoByID)
		photoRoute.DELETE("/:photoID", middlewares.PhotoAuthorization(), controllers.DeletePhotoByID)
	}

	commentRoute := r.Group("/comments")
	{
		commentRoute.Use(middlewares.Auth())
		commentRoute.POST("/", controllers.CreateComment)
		commentRoute.GET("/", controllers.GetComments)
		commentRoute.PUT("/:commentID", middlewares.CommentAuthorization(), controllers.UpdateCommentByID)
		commentRoute.DELETE("/:commentID", middlewares.CommentAuthorization(), controllers.DeleteCommentByID)
	}

	socialMediaRoute := r.Group("/socialmedias")
	{
		socialMediaRoute.Use(middlewares.Auth())
		socialMediaRoute.POST("/", controllers.CreateSocialMedia)
		socialMediaRoute.GET("/", controllers.GetSocialMedia)
		socialMediaRoute.PUT("/:socialMediaID", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMediaByID)
		socialMediaRoute.DELETE("/:socialMediaID", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMediaByID)
	}

	return r
}
