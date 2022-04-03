package handler

import (
	"clother"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strconv"
)

func (h *Handler) CreateItem(c *gin.Context) {
	user, _ := c.Get("user")

	if !user.(clother.User).Admin {
		newErrorResponse(c, 403, "You don't have enough permissions")
		return
	}

	var item clother.Item
	var err error
	item.CategoryID, err = uuid.Parse(c.Param("id"))
	if err != nil {
		newErrorResponse(c, 500, err.Error())
		return
	}

	category, err := h.services.Category.GetCategoryByID(item.CategoryID)
	if err != nil && err != sql.ErrNoRows {
		newErrorResponse(c, 500, err.Error())
		return
	} else if category.ID == uuid.Nil {
		newErrorResponse(c, 404, "Category not found")
		return
	}

	item.Title, item.Description = c.PostForm("title"), c.PostForm("description")
	if len(item.Title) < 3 || len(item.Description) < 6 {
		newErrorResponse(c, 400, "Title must be more than 3 chars and description more than 6 chars")
		return
	}

	item.Cost, err = strconv.Atoi(c.PostForm("cost"))
	if err != nil || item.Cost <= 0 {
		newErrorResponse(c, 400, "Invalid cost")
		return
	}

	itemID, err := h.services.Item.CreateItem(item)
	if err != nil {
		newErrorResponse(c, 500, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"id": itemID,
	})
}

func (h *Handler) GetAllCategoryItems(c *gin.Context) {
	categoryID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		newErrorResponse(c, 500, err.Error())
		return
	}

	category, err := h.services.Category.GetCategoryByID(categoryID)
	if err != nil && err != sql.ErrNoRows {
		newErrorResponse(c, 500, err.Error())
		return
	} else if category.ID == uuid.Nil {
		newErrorResponse(c, 404, "Category not found")
		return
	}

	items, err := h.services.Item.GetAllCategoryItems(category.ID)
	if err != nil {
		newErrorResponse(c, 500, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"items": items,
	})
}
