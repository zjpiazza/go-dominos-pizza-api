package models

import (
	"strings"

	"github.com/zjpiazza/go-dominos-pizza-api/pkg/utils"
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

	// Get the store search url
	url := strings.Replace(utils.URLs.Store.Find, "${line1}", addr.GetDefaultLineOne(), -1)
	url = strings.Replace(url, "${line2}", addr.GetDefaultLineTwo(), -1)
	url = strings.Replace(url, "${type}", "Delivery", -1) // Default to delivery

	// Make the request
	response, err := utils.Get(url)
	if err != nil {
		return nil, err
	}

	// Process the response
	if storesData, ok := response["Stores"].([]interface{}); ok {
		for _, storeData := range storesData {
			if storeMap, ok := storeData.(map[string]interface{}); ok {
				store := &Store{}
				store.SetFormatted(storeMap)
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
