package handler

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" || header == " " {
		newErrorResponse(c, 403, "Empty authorization header")
		return
	}

	authHeader := strings.Split(header, " ")
	if len(authHeader) != 2 || authHeader[0] != "Bearer" {
		newErrorResponse(c, 403, "Invalid authorization header")
		return
	}

	userID, err := h.services.Authorization.ParseToken(authHeader[1])
	if err != nil {
		newErrorResponse(c, 500, err.Error())
		return
	}

	user, err := h.services.Authorization.GetUserByID(userID)
	if err != nil {
		newErrorResponse(c, 500, err.Error())
		return
	}

	c.Set("user", user)
}
