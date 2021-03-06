package handler

import (
	"clother/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		services: service,
	}
}

func (h Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	category := router.Group("/category")
	{
		category.POST("/", h.userIdentity, h.CreateCategory)
		category.POST("/:id/", h.userIdentity, h.CreateItem)
		category.GET("/", h.GetAllCategories)
		category.GET("/:id/", h.GetAllCategoryItems)
	}

	return router
}
