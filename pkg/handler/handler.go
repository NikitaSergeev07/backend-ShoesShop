package handler

import (
	"ShoesShop/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("logout", h.logout)
	}

	api := router.Group("/api")
	{
		items := api.Group("/items")
		{
			items.POST("/", h.createItem)
			items.GET("/:id", h.getItemById)
			items.GET("/", h.getAllItems)
			items.PUT("/:id", h.updateItem)
			items.DELETE("/:id", h.deleteItem)
			items.GET("/search", h.searchItems)
			items.PUT("/", h.updateItem)
		}
		reviews := api.Group("/reviews")
		{
			reviews.POST("/", h.createReview)
			reviews.GET("/", h.getAllReviews)
		}
		favorites := api.Group("/favorites")
		{
			favorites.GET("/", h.getFavoritesByUser)
			favorites.POST("/", h.addFavorite)
			favorites.DELETE("/", h.removeFavorite)
		}
		cart := api.Group("/carts")
		{
			cart.GET("/", h.getCartByUser)
			cart.POST("/", h.addToCart)
			cart.DELETE("/", h.removeFromCart)
			cart.DELETE("/all", h.removeAllFromCart)
		}
		payments := api.Group("/payments")
		{
			payments.POST("", h.CreatePayment)
		}
		feedbacks := api.Group("/feedbacks")
		{
			feedbacks.POST("", FeedbackHandlerGin)
		}
	}

	//TEST for JWT
	ping := router.Group("/ping")
	ping.Use(h.userIdentity)
	{
		ping.GET("", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "pong"})
		})
	}

	return router
}
