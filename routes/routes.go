package routes

import (
	"customers-management/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	SetupUserRoutes(r)
	SetupCardRoutes(r)
	SetupTermRoutes(r)
	SetupBlockRoutes(r)
}

func SetupUserRoutes(r *gin.Engine) {
	userGroup := r.Group("/users")
	{
		userGroup.GET("/", controllers.UsersIndex)
		userGroup.GET("/:id", controllers.UsersShow)
		userGroup.POST("/", controllers.UsersCreate)
		userGroup.PUT("/:id", controllers.UsersUpdate)
		userGroup.DELETE("/:id", controllers.UsersDelete)
	}
}

func SetupCardRoutes(r *gin.Engine) {
	cardGroup := r.Group("/cards")
	{
		cardGroup.GET("/", controllers.CardsIndex)
		cardGroup.GET("/:id", controllers.CardsShow)
		cardGroup.GET("/user/:user_id", controllers.CardsIndexByUser)
		cardGroup.POST("/", controllers.CardsCreate)
		cardGroup.DELETE("/:id", controllers.CardsDelete)
	}
}

func SetupTermRoutes(r *gin.Engine) {
	termGroup := r.Group("/terms")
	{
		termGroup.GET("/", controllers.TermsIndex)
		termGroup.GET("/:version", controllers.TermsShow)
		termGroup.POST("/", controllers.TermsCreate)
		termGroup.POST("/accept", controllers.TermsAccept)
	}
}

func SetupBlockRoutes(r *gin.Engine) {
	blockGroup := r.Group("/block")
	{
		blockGroup.PUT("/card/:id", controllers.CardsBlock)
		blockGroup.PUT("/user/:id", controllers.UsersBlock)
	}
}