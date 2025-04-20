package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/zjpiazza/go-dominos-pizza-api/pkg/utils"
)

// Store represents a Domino's Pizza store
type Store struct {
	DominosFormat
	ID                  string                 `json:"id"`
	StoreID             string                 `json:"storeID"`
	Phone               string                 `json:"phone"`
	IsDeliveryStore     bool                   `json:"isDeliveryStore"`
	IsOpen              bool                   `json:"isOpen"`
	IsOnlineCapable     bool                   `json:"isOnlineCapable"`
	IsOnlineNow         bool                   `json:"isOnlineNow"`
	MinDistance         float64                `json:"minDistance"`
	ServiceHours        map[string]interface{} `json:"serviceHours"`
	ServiceIsOpen       map[string]interface{} `json:"serviceIsOpen"`
	AllowDeliveryOrders bool                   `json:"allowDeliveryOrders"`
	AllowCarryoutOrders bool                   `json:"allowCarryoutOrders"`
	StoreCoordinates    struct {
		StoreLatitude  string `json:"storeLatitude"`
		StoreLongitude string `json:"storeLongitude"`
	} `json:"storeCoordinates"`
}

// NewStore creates a new store from store ID
func NewStore(storeID string) (*Store, error) {
	store := &Store{
		StoreID: storeID,
	}

	// Get store info from API
	url := strings.Replace(utils.URLs.Store.Info, "${storeID}", storeID, -1)
	response, err := utils.Get(url)
	if err != nil {
		return nil, err
	}

	store.SetFormatted(response)

	return store, nil
}

// GetMenu retrieves the menu for this store
func (s *Store) GetMenu(lang string) (*Menu, error) {
	if s.StoreID == "" {
		return nil, utils.NewDominosStoreError("Store ID is required to get menu")
	}

	if lang == "" {
		lang = "en"
	}

	menu := &Menu{}

	// Get menu from API
	url := fmt.Sprintf("https://order.dominos.com/power/store/%s/menu?lang=%s&structured=true",
		s.StoreID, lang)

	// Create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Set headers - these are critical for the Dominos API to respond correctly
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.63 Safari/537.36")
	req.Header.Set("Origin", "https://order.dominos.com")
	req.Header.Set("Referer", "https://order.dominos.com/")

	// Send request
	client := &http.Client{Timeout: time.Second * 30}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse JSON response
	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	// Set formatted data in the menu struct
	menu.SetFormatted(response)

	// Also store the raw response for direct access
	menu.SetDominosAPIResponse(response)

	return menu, nil
}

// IsCurrentlyOpen checks if this store is currently open
func (s *Store) IsCurrentlyOpen(serviceMethod string) bool {
	if !s.IsOpen || !s.IsOnlineCapable || !s.IsOnlineNow {
		return false
	}

	// Check if the specific service method is open
	if serviceMethod == "Delivery" {
		if !s.IsDeliveryStore || !s.AllowDeliveryOrders {
			return false
		}
		// Check if delivery service is open
		if open, ok := s.ServiceIsOpen["Delivery"].(bool); ok {
			return open
		}
	} else if serviceMethod == "Carryout" {
		if !s.AllowCarryoutOrders {
			return false
		}
		// Check if carryout service is open
		if open, ok := s.ServiceIsOpen["Carryout"].(bool); ok {
			return open
		}
	}

	return false
}
