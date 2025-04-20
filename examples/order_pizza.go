package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/zjpiazza/go-dominos-pizza-api"
)

func main() {
	// Create a customer
	customerData := map[string]interface{}{
		"address":   "2 Portola Plaza, Monterey, Ca, 93940",
		"firstName": "John",
		"lastName":  "Doe",
		"phone":     "555-555-5555",
		"email":     "test@example.com",
	}

	customer, err := dominos.NewCustomer(customerData)
	if err != nil {
		log.Fatalf("Error creating customer: %v", err)
	}

	// Create a pizza item
	itemData := map[string]interface{}{
		"code": "14SCREEN", // 14-inch hand-tossed pizza
		"options": map[string]interface{}{
			"X": map[string]interface{}{"1/1": "1"},   // Regular sauce
			"C": map[string]interface{}{"1/1": "1.5"}, // Extra cheese
			"P": map[string]interface{}{"1/2": "1.5"}, // Double pepperoni on half
		},
	}

	pizza, err := dominos.NewItem(itemData)
	if err != nil {
		log.Fatalf("Error creating pizza: %v", err)
	}

	// Find nearby stores
	nearbyStores, err := dominos.NewNearbyStores(customer.Address)
	if err != nil {
		log.Fatalf("Error finding nearby stores: %v", err)
	}

	fmt.Printf("Found %d stores near the address\n", len(nearbyStores.Stores))

	// Find the closest delivery store
	closestStore := nearbyStores.FindClosestStore("Delivery", true)
	if closestStore == nil {
		log.Fatal("No delivery stores found")
	}

	fmt.Printf("Selected store: ID=%s, Distance=%.2f miles\n",
		closestStore.StoreID, closestStore.MinDistance)

	// Create an order
	order := dominos.NewOrder(customer)
	order.StoreID = closestStore.StoreID

	// Add the pizza to the order
	order.AddItem(pizza)

	// Validate the order
	err = order.Validate()
	if err != nil {
		log.Fatalf("Order validation failed: %v", err)
	}

	fmt.Println("Order validated successfully")

	// Price the order
	err = order.Price()
	if err != nil {
		log.Fatalf("Order pricing failed: %v", err)
	}

	// Display the order total
	if amountStr, ok := order.AmountsBreakdown["Customer"].(string); ok {
		amount, _ := strconv.ParseFloat(amountStr, 64)
		fmt.Printf("Order total: $%.2f\n", amount)
	}

	// Create a payment (this is a fake credit card)
	paymentData := map[string]interface{}{
		"number":       "4100-1234-2234-3234",
		"expiration":   "01/35",
		"securityCode": "123",
		"postalCode":   "93940",
		"tipAmount":    3.00,
	}

	// Get the amount from the order
	if amountStr, ok := order.AmountsBreakdown["Customer"].(string); ok {
		amount, _ := strconv.ParseFloat(amountStr, 64)
		paymentData["amount"] = amount
	}

	payment, err := dominos.NewPayment(paymentData)
	if err != nil {
		log.Fatalf("Error creating payment: %v", err)
	}

	order.Payments = append(order.Payments, payment)

	// Normally, you would place the order here
	fmt.Println("\nTo place the order, uncomment the following code:")
	fmt.Println("// err = order.Place()")
	fmt.Println("// if err != nil {")
	fmt.Println("//     log.Fatalf(\"Order placement failed: %v\", err)")
	fmt.Println("// }")
	fmt.Println("// fmt.Println(\"Order placed successfully!\")")

	// This is not called because it would submit a real order with a fake credit card
	// err = order.Place()
	// if err != nil {
	//     log.Fatalf("Order placement failed: %v", err)
	// }
	// fmt.Println("Order placed successfully!")
}
