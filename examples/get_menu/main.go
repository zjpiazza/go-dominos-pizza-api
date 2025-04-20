package main

import (
	"fmt"
	"log"

	dominos "github.com/zjpiazza/go-dominos-pizza-api"
	"github.com/zjpiazza/go-dominos-pizza-api/pkg/models"
)

func main() {
	// Set USA as the region - should be the default anyway
	dominos.UseInternational(dominos.USA)

	// Use a known store ID from the previous find_stores example
	storeID := "8180" // Beverly Hills store
	fmt.Printf("Getting menu for Domino's Pizza store ID: %s\n", storeID)

	// Create a new store with the ID
	store, err := models.NewStore(storeID)
	if err != nil {
		log.Fatalf("Error creating store: %v", err)
	}

	// Get the menu for this store
	fmt.Println("Retrieving menu...")
	menu, err := store.GetMenu("en")
	if err != nil {
		log.Fatalf("Error getting menu: %v", err)
	}

	// Print some product counts
	fmt.Printf("\nFound %d total products on the menu\n", len(menu.GetRawProducts()))
	fmt.Println("======================================")

	// List pizzas
	pizzas := menu.GetPizzas()
	fmt.Printf("\nPIZZAS (%d):\n", len(pizzas))
	fmt.Println("--------------------------------------")
	pizzaCounter := 0
	for code, pizza := range pizzas {
		name, _ := pizza["Name"].(string)
		fmt.Printf("  - %s: %s\n", code, name)

		pizzaCounter++
		if pizzaCounter >= 5 {
			fmt.Printf("  ... and %d more pizzas\n", len(pizzas)-5)
			break
		}
	}

	// List sides
	sides := menu.GetSides()
	fmt.Printf("\nSIDES AND OTHER FOOD (%d):\n", len(sides))
	fmt.Println("--------------------------------------")
	sideCounter := 0
	for code, side := range sides {
		name, _ := side["Name"].(string)
		fmt.Printf("  - %s: %s\n", code, name)

		sideCounter++
		if sideCounter >= 5 {
			fmt.Printf("  ... and %d more sides\n", len(sides)-5)
			break
		}
	}

	// List drinks
	drinks := menu.GetDrinks()
	fmt.Printf("\nDRINKS (%d):\n", len(drinks))
	fmt.Println("--------------------------------------")
	if len(drinks) > 0 {
		drinkCounter := 0
		for code, drink := range drinks {
			name, _ := drink["Name"].(string)
			fmt.Printf("  - %s: %s\n", code, name)

			drinkCounter++
			if drinkCounter >= 5 {
				fmt.Printf("  ... and %d more drinks\n", len(drinks)-5)
				break
			}
		}
	} else {
		fmt.Println("  No drinks found on the menu")
	}

	fmt.Println("\nMenu successfully retrieved!")
	fmt.Println("======================================")
	fmt.Println("Example: To get details about a specific product, use menu.GetProduct(productCode)")
}
