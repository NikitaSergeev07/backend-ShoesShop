package service

import (
	"github.com/rvinnie/yookassa-sdk-go/yookassa"
	yoocommon "github.com/rvinnie/yookassa-sdk-go/yookassa/common"
	yoopayment "github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
)

type PaymentService struct {
	client *yookassa.Client
}

func NewPaymentService(shopID, apiKey string) *PaymentService {
	client := yookassa.NewClient(shopID, apiKey)
	return &PaymentService{client: client}
}

func (s *PaymentService) CreatePayment(amount, description string) (*yoopayment.Payment, error) {
	paymentHandler := yookassa.NewPaymentHandler(s.client)

	payment, err := paymentHandler.CreatePayment(&yoopayment.Payment{
		Amount: &yoocommon.Amount{
			Value:    amount,
			Currency: "RUB",
		},
		PaymentMethod: yoopayment.PaymentMethodType("bank_card"),
		Confirmation: &yoopayment.Redirect{
			Type:      "redirect",
			ReturnURL: "http://localhost:8080/#/home",
		},
		Description: description,
	})

	if err != nil {
		return nil, err
	}
	return payment, nil
}
