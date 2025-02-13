package handler

import (
	"ShoesShop"
	"ShoesShop/enums"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) createItem(c *gin.Context) {
	var input ShoesShop.Item

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Item.CreateItem(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (h *Handler) searchItems(c *gin.Context) {
	query := c.DefaultQuery("query", "")

	sortOption := c.DefaultQuery("sort", "name")
	sortOption = strings.TrimSpace(sortOption)

	log.Printf("Raw sort option from query: '%s'", c.Query("sort"))
	log.Printf("Parsed sort option: '%s'", sortOption)

	var sortEnum enums.SortOption
	switch sortOption {
	case "name":
		sortEnum = enums.SortByName
	case "price_asc":
		sortEnum = enums.SortByPriceAsc
	case "price_desc":
		sortEnum = enums.SortByPriceDesc
	case "Shoes1":
		sortEnum = enums.Shoes1
	case "Shoes2":
		sortEnum = enums.Shoes2
	default:
		newErrorResponse(c, http.StatusBadRequest, "invalid sort parameter")
		return
	}

	items, err := h.services.Item.SearchItems(query, sortEnum)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}

func (h *Handler) getItemById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	item, err := h.services.Item.GetItemById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) getAllItems(c *gin.Context) {
	items, err := h.services.Item.GetAllItems()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}

func (h *Handler) updateItem(c *gin.Context) {
	var input ShoesShop.Item

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Item.UpdateItem(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) deleteItem(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	err = h.services.Item.DeleteItem(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
