package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// NearbyStores represents nearby Domino's Pizza stores
type NearbyStores struct {
	DominosFormat
	Address *Address `json:"address"`
	Stores  []*Store `json:"stores"`
}

// NewNearbyStores finds stores near an address
func NewNearbyStores(address interface{}) (*NearbyStores, error) {
	// Parse the address
	addr, err := NewAddress(address)
	if err != nil {
		return nil, err
	}

	nearbyStores := &NearbyStores{
		Address: addr,
		Stores:  make([]*Store, 0),
	}

	// Get the street and city+state+zip parts separately
	street := addr.GetDefaultLineOne()
	cityStateZip := addr.GetDefaultLineTwo()

	// URL encode the parameters
	encodedStreet := url.QueryEscape(street)
	encodedCityStateZip := url.QueryEscape(cityStateZip)

	// Construct the URL directly with the encoded parameters
	urlStr := fmt.Sprintf("https://order.dominos.com/power/store-locator?s=%s&c=%s&type=Delivery",
		encodedStreet, encodedCityStateZip)

	// Make a direct HTTP request instead of using utils.Get
	// This ensures we have more control over the headers and request format
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Create request
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return nil, err
	}

	// Set headers that might be required
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Origin", "https://order.dominos.com")
	req.Header.Set("Referer", "https://order.dominos.com/")

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse JSON response
	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	// Process the response
	if storesData, ok := response["Stores"].([]interface{}); ok {
		for _, storeData := range storesData {
			if storeMap, ok := storeData.(map[string]interface{}); ok {
				store := &Store{}

				// Set the store ID
				if storeID, ok := storeMap["StoreID"].(string); ok {
					store.StoreID = storeID
				}

				// Set the phone number
				if phone, ok := storeMap["Phone"].(string); ok {
					store.Phone = phone
				}

				// Set the distance
				if minDistance, ok := storeMap["MinDistance"].(float64); ok {
					store.MinDistance = minDistance
				}

				// Set the delivery flag
				if isDeliveryStore, ok := storeMap["IsDeliveryStore"].(bool); ok {
					store.IsDeliveryStore = isDeliveryStore
				}

				// Set the open flag
				if isOpen, ok := storeMap["IsOpen"].(bool); ok {
					store.IsOpen = isOpen
				}

				// Set the online capabilities
				if isOnlineCapable, ok := storeMap["IsOnlineCapable"].(bool); ok {
					store.IsOnlineCapable = isOnlineCapable
				}

				if isOnlineNow, ok := storeMap["IsOnlineNow"].(bool); ok {
					store.IsOnlineNow = isOnlineNow
				}

				// Set store coordinates
				if coordinates, ok := storeMap["StoreCoordinates"].(map[string]interface{}); ok {
					if lat, ok := coordinates["StoreLatitude"].(string); ok {
						store.StoreCoordinates.StoreLatitude = lat
					}

					if lng, ok := coordinates["StoreLongitude"].(string); ok {
						store.StoreCoordinates.StoreLongitude = lng
					}
				}

				// Append the store to our list
				nearbyStores.Stores = append(nearbyStores.Stores, store)
			}
		}
	}

	return nearbyStores, nil
}

// FindClosestStore finds the closest store that meets the criteria
func (ns *NearbyStores) FindClosestStore(serviceMethod string, isOpen bool) *Store {
	var closestStore *Store
	minDistance := 100.0 // Default max distance is 100 miles

	for _, store := range ns.Stores {
		// Skip stores that aren't open if we require open stores
		if isOpen && !store.IsCurrentlyOpen(serviceMethod) {
			continue
		}

		// For delivery, check if it's a delivery store
		if serviceMethod == "Delivery" && !store.IsDeliveryStore {
			continue
		}

		// Check the distance
		if store.MinDistance < minDistance {
			minDistance = store.MinDistance
			closestStore = store
		}
	}

	return closestStore
}
