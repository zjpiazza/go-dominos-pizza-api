package main

import (
	"fmt"
	"log"

	"github.com/zjpiazza/go-dominos-pizza-api"
	"github.com/zjpiazza/go-dominos-pizza-api/pkg/models"
)

func main() {
	// Set USA as the region - should be the default anyway
	dominos.UseInternational(dominos.USA)

	// Address to search for nearby stores
	address := "9876 Wilshire Blvd, Beverly Hills, CA 90210"
	fmt.Printf("Finding Domino's Pizza stores near: %s\n\n", address)

	// Parse the address
	addr, err := models.NewAddress(address)
	if err != nil {
		log.Fatalf("Error parsing address: %v\n", err)
	}

	fmt.Println("Address Details:")
	fmt.Printf("  Street: %s\n", addr.GetDefaultLineOne())
	fmt.Printf("  City: %s\n", addr.City)
	fmt.Printf("  Region: %s\n", addr.Region)
	fmt.Printf("  Postal Code: %s\n", addr.PostalCode)

	// Find nearby stores
	nearbyStores, err := models.NewNearbyStores(address)
	if err != nil {
		log.Fatalf("Error finding stores: %v\n", err)
	}

	fmt.Printf("Found %d stores\n\n", len(nearbyStores.Stores))

	if len(nearbyStores.Stores) == 0 {
		fmt.Println("No stores found")
		return
	}

	// Print store information
	for i, store := range nearbyStores.Stores {
		if i >= 5 {
			fmt.Printf("\n... and %d more stores\n", len(nearbyStores.Stores)-5)
			break
		}

		fmt.Printf("Store #%d:\n", i+1)
		fmt.Printf("  Store ID: %s\n", store.StoreID)
		fmt.Printf("  Phone: %s\n", store.Phone)
		fmt.Printf("  Distance: %.1f miles\n\n", store.MinDistance)
	}
}
