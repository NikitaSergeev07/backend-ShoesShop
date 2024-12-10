package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) CreatePayment(c *gin.Context) {
	type PaymentRequest struct {
		Amount      string `json:"amount" binding:"required"`
		Description string `json:"description" binding:"required"`
	}

	var request PaymentRequest
	if err := c.BindJSON(&request); err != nil {
		logrus.Errorf("Invalid input: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input: " + err.Error()})
		return
	}

	// Логируем запрос
	logrus.Infof("Payment request received: %+v", request)

	payment, err := h.services.Payment.CreatePayment(request.Amount, request.Description)
	if err != nil {
		logrus.Errorf("Failed to create payment: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create payment: " + err.Error()})
		return
	}

	// Извлекаем confirmation_url из структуры Confirmation
	var confirmationURL string
	if confirmation, ok := payment.Confirmation.(map[string]interface{}); ok {
		if url, exists := confirmation["confirmation_url"].(string); exists {
			confirmationURL = url
		}
	}

	// Логируем ответ от YooKassa
	logrus.Infof("Payment response: %+v", payment)

	// Логируем подготовленный ответ клиенту
	response := gin.H{
		"id":               payment.ID,
		"status":           payment.Status,
		"confirmation_url": confirmationURL,
		"return_url":       "http://localhost:8080/#/home",
		"amount": gin.H{
			"value":    payment.Amount.Value,
			"currency": payment.Amount.Currency,
		},
		"description": payment.Description,
	}
	logrus.Infof("Response to client: %+v", response)

	c.JSON(http.StatusOK, response)
}
