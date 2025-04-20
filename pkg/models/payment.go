package models

import (
	"strings"
)

// Payment represents a payment method for an order
type Payment struct {
	DominosFormat
	Type         string  `json:"type"`
	Amount       float64 `json:"amount"`
	Number       string  `json:"number"`
	Expiration   string  `json:"expiration"` // MM/YY format
	SecurityCode string  `json:"securityCode"`
	PostalCode   string  `json:"postalCode"`
	TipAmount    float64 `json:"tipAmount"`
}

// NewPayment creates a new payment from payment data
func NewPayment(paymentData map[string]interface{}) (*Payment, error) {
	payment := &Payment{
		Type: "CreditCard", // Default payment type
	}

	// Set payment fields from paymentData
	payment.SetFormatted(paymentData)

	// Sanitize credit card number (remove dashes, spaces)
	payment.Number = strings.ReplaceAll(payment.Number, "-", "")
	payment.Number = strings.ReplaceAll(payment.Number, " ", "")

	// Sanitize expiration date (remove slashes)
	payment.Expiration = strings.ReplaceAll(payment.Expiration, "/", "")

	return payment, nil
}
