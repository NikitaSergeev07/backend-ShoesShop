package handler

import (
	"ShoesShop"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) addToCart(c *gin.Context) {
	var input ShoesShop.Cart

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid cart input: "+err.Error())
		return
	}

	id, err := h.services.Cart.AddToCart(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (h *Handler) removeFromCart(c *gin.Context) {
	userIdStr := c.Query("user_id")
	itemIdStr := c.Query("item_id")
	size := c.Query("size")

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

	if err := h.services.Cart.RemoveFromCart(userId, itemId, size); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "item removed from cart successfully",
	})
}

func (h *Handler) getCartByUser(c *gin.Context) {
	userIdStr := c.Query("user_id")

	userId, err := strconv.Atoi(userIdStr)
	if err != nil || userId <= 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid user_id")
		return
	}
	cartItems, err := h.services.Cart.GetCartByUserId(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, cartItems)
}

func (h *Handler) removeAllFromCart(c *gin.Context) {
	userIdStr := c.Query("user_id")

	userId, err := strconv.Atoi(userIdStr)
	if err != nil || userId <= 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid user_id")
		return
	}

	if err := h.services.Cart.RemoveAll(userId); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "all items removed from cart successfully",
	})
}
