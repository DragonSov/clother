package handler

import (
	"clother"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCategory(c *gin.Context) {
	user, _ := c.Get("user")

	if !user.(clother.User).Admin {
		newErrorResponse(c, 403, "You don't have enough permissions")
		return
	}

	var category clother.Category
	category.Title = c.PostForm("title")
	if len(category.Title) < 3 {
		newErrorResponse(c, 400, "Title must be more than 3 chars")
		return
	}

	categoryID, err := h.services.Category.CreateCategory(category)
	if err != nil {
		newErrorResponse(c, 500, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"id": categoryID,
	})
}

func (h *Handler) GetAllCategories(c *gin.Context) {
	categories, err := h.services.Category.GetAllCategories()
	if err != nil {
		newErrorResponse(c, 500, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"categories": categories,
	})
}
