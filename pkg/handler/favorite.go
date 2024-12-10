package handler

import (
	"ShoesShop"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) addFavorite(c *gin.Context) {
	var input ShoesShop.Favorite

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid favorite input: "+err.Error())
		return
	}

	id, err := h.services.Favorite.AddFavorite(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (h *Handler) removeFavorite(c *gin.Context) {
	userIdStr := c.Query("user_id")
	itemIdStr := c.Query("item_id")

	userId, err := strconv.Atoi(userIdStr)
	if err != nil || userId <= 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid user_id")
		return
	}

	itemId, err := strconv.Atoi(itemIdStr)
	if err != nil || itemId <= 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid item_id")
		return
	}

	if err := h.services.Favorite.RemoveFavorite(userId, itemId); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "favorite removed successfully",
	})
}

func (h *Handler) getFavoritesByUser(c *gin.Context) {
	userIdStr := c.Query("user_id")

	userId, err := strconv.Atoi(userIdStr)
	if err != nil || userId <= 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid user_id")
		return
	}

	items, err := h.services.Favorite.GetFavoritesByUserId(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}
