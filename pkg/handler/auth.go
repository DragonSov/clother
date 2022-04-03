package handler

import (
	"clother"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) signUp(c *gin.Context) {
	var user clother.User

	user.Login, user.Password = c.PostForm("login"), c.PostForm("password")
	if len(user.Login) < 3 || len(user.Password) < 6 {
		newErrorResponse(c, 400, "Login must be more than 3 chars and the password more than 6 chars")
		return
	}

	selectedUser, err := h.services.Authorization.GetUserByLogin(user.Login)
	if err != nil && err != sql.ErrNoRows {
		newErrorResponse(c, 500, err.Error())
		return
	} else if selectedUser.ID != uuid.Nil {
		newErrorResponse(c, 409, "User already exists")
		return
	}

	id, err := h.services.Authorization.CreateUser(user)
	if err != nil {
		newErrorResponse(c, 500, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	login, password := c.PostForm("login"), c.PostForm("password")
	if len(login) < 3 || len(password) < 6 {
		newErrorResponse(c, 400, "Login must be more than 3 chars and password more than 6 chars")
		return
	}

	token, err := h.services.Authorization.GenerateToken(login, password)
	if err != nil {
		newErrorResponse(c, 500, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}
