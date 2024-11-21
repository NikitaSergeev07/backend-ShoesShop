package handler

import (
	"ShoesShop"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *Handler) signUp(c *gin.Context) { // регистрация
	var input ShoesShop.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) { // аутентификация(авторизация)
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if h.services.Authorization.IsTokenBlacklisted(input.Email) {
		newErrorResponse(c, http.StatusUnauthorized, "Previous token was invalidated. Please log in again.")
		return
	}

	id, err := h.services.Authorization.GetUser(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"token": token,
	})
}

func (h *Handler) logout(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "No authorization header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "Invalid authorization header")
		return
	}

	token := headerParts[1]

	// Аннулируем токен, добавляя его в черный список
	err := h.services.Authorization.InvalidateToken(token)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Failed to logout")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully logged out",
	})
}
