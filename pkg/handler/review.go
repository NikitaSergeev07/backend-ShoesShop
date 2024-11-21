package handler

import (
	"ShoesShop"
	"ShoesShop/enums"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (h *Handler) createReview(c *gin.Context) {
	var input ShoesShop.Review

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid review input: "+err.Error())
		return
	}

	if input.Category != enums.Service && input.Category != enums.Product && input.Category != enums.WebSite {
		newErrorResponse(c, http.StatusBadRequest, "invalid category")
		return
	}

	if input.Category == enums.Product && (input.ItemId == nil || *input.ItemId == 0) {
		newErrorResponse(c, http.StatusBadRequest, "item_id is required for category 'Product'")
		return
	}

	if input.Date.IsZero() {
		input.Date = time.Now()
	}

	id, err := h.services.Review.CreateReview(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (h *Handler) getAllReviews(c *gin.Context) {
	reviews, err := h.services.Review.GetAllReviews()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, reviews)
}
