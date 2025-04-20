package models

import (
	"fmt"
	"strings"

	"github.com/zjpiazza/go-dominos-pizza-api/pkg/utils"
)

// Tracking represents a Domino's Pizza order tracking
type Tracking struct {
	DominosFormat
	OrderID        string `json:"orderID"`
	Phone          string `json:"phone"`
	ServiceMethod  string `json:"serviceMethod"`
	OrderKey       string `json:"orderKey"`
	PulseOrderGUID string `json:"pulseOrderGUID"`
}

// NewTracking creates a new tracking instance
func NewTracking() *Tracking {
	return &Tracking{}
}

// ByID gets tracking information by order ID
func (t *Tracking) ByID(orderID string) (map[string]interface{}, error) {
	if orderID == "" {
		return nil, utils.NewDominosTrackingError("Order ID is required for tracking")
	}

	t.OrderID = orderID

	// Construct the tracking URL
	var url string

	// The old method (for Canada) and new method differ
	if strings.Contains(utils.URLs.Track, "orderstorage") {
		// Old method (e.g., Canada)
		url = fmt.Sprintf("%sOrderKey=%s", utils.URLs.Track, orderID)
	} else {
		// New method (e.g., USA)
		url = fmt.Sprintf("%s/%s/%s", utils.URLs.TrackRoot, utils.URLs.Track, orderID)
	}

	// Make the tracking request
	response, err := utils.GetTracking(url, "")
	if err != nil {
		return nil, err
	}

	return response, nil
}

// ByPhone gets tracking information by phone number
func (t *Tracking) ByPhone(phone string) (map[string]interface{}, error) {
	if phone == "" {
		return nil, utils.NewDominosTrackingError("Phone number is required for tracking")
	}

	// Sanitize phone number (remove non-digits)
	sanitizedPhone := ""
	for _, char := range phone {
		if char >= '0' && char <= '9' {
			sanitizedPhone += string(char)
		}
	}

	t.Phone = sanitizedPhone

	// Construct the tracking URL
	var url string

	// The old method (for Canada) and new method differ
	if strings.Contains(utils.URLs.Track, "orderstorage") {
		// Old method (e.g., Canada)
		url = fmt.Sprintf("%sPhone=%s", utils.URLs.Track, sanitizedPhone)
	} else {
		// New method (e.g., USA)
		url = fmt.Sprintf("%s/%s/phone/%s", utils.URLs.TrackRoot, utils.URLs.Track, sanitizedPhone)
	}

	// Make the tracking request
	response, err := utils.GetTracking(url, "")
	if err != nil {
		return nil, err
	}

	return response, nil
}
