# Go Domino's Pizza API

A Go wrapper for the Domino's Pizza API, converted from the Node.js version.

## Installation

```bash
go get github.com/zjpiazza/go-dominos-pizza-api
```

## Usage

### Importing

```go
import "github.com/zjpiazza/go-dominos-pizza-api/go"
```

### Basic Example

```go
package main

import (
	"fmt"
	"log"

	"github.com/zjpiazza/go-dominos-pizza-api/go"
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
	
	// Find nearby stores
	nearbyStores, err := dominos.NewNearbyStores(customer.Address)
	if err != nil {
		log.Fatalf("Error finding nearby stores: %v", err)
	}
	
	// Find the closest delivery store
	closestStore := nearbyStores.FindClosestStore("Delivery", true)
	if closestStore == nil {
		log.Fatal("No delivery stores found")
	}
	
	// Create a pizza
	itemData := map[string]interface{}{
		"code": "14SCREEN", // 14-inch hand-tossed pizza
		"options": map[string]interface{}{
			"X": map[string]interface{}{"1/1": "1"},   // Regular sauce
			"C": map[string]interface{}{"1/1": "1.5"}, // Extra cheese
		},
	}
	
	pizza, err := dominos.NewItem(itemData)
	if err != nil {
		log.Fatalf("Error creating pizza: %v", err)
	}
	
	// Create an order
	order := dominos.NewOrder(customer)
	order.StoreID = closestStore.StoreID
	order.AddItem(pizza)
	
	// Validate the order
	err = order.Validate()
	if err != nil {
		log.Fatalf("Order validation failed: %v", err)
	}
	
	// Price the order
	err = order.Price()
	if err != nil {
		log.Fatalf("Order pricing failed: %v", err)
	}
	
	// Add payment (this is a fake credit card)
	paymentData := map[string]interface{}{
		"amount":       order.AmountsBreakdown["Customer"],
		"number":       "4100-1234-2234-3234",
		"expiration":   "01/35",
		"securityCode": "123",
		"postalCode":   "93940",
		"tipAmount":    3.00,
	}
	
	payment, err := dominos.NewPayment(paymentData)
	if err != nil {
		log.Fatalf("Error creating payment: %v", err)
	}
	
	order.Payments = append(order.Payments, payment)
	
	// Place the order (commented out to prevent accidental orders)
	/*
	err = order.Place()
	if err != nil {
		log.Fatalf("Order placement failed: %v", err)
	}
	fmt.Println("Order placed successfully!")
	*/
}
```

## API Reference

### Models

- `Address` - Represents a delivery or pickup address
- `Customer` - Represents a Domino's Pizza customer
- `Item` - Represents a product item in an order
- `Menu` - Represents a Domino's Pizza menu
- `NearbyStores` - Represents nearby Domino's Pizza stores
- `Order` - Represents a Domino's Pizza order
- `Payment` - Represents a payment method for an order
- `Store` - Represents a Domino's Pizza store
- `Tracking` - Represents Domino's Pizza order tracking

### Constructors

- `NewAddress(address interface{}) (*Address, error)` - Creates a new address from a string or address object
- `NewCustomer(customerData map[string]interface{}) (*Customer, error)` - Creates a new customer
- `NewItem(itemData map[string]interface{}) (*Item, error)` - Creates a new product item
- `NewNearbyStores(address interface{}) (*NearbyStores, error)` - Finds nearby stores
- `NewOrder(customer *Customer) *Order` - Creates a new order with a customer
- `NewPayment(paymentData map[string]interface{}) (*Payment, error)` - Creates a new payment method
- `NewStore(storeID string) (*Store, error)` - Creates a new store from a store ID
- `NewTracking() *Tracking` - Creates a new tracking instance

### Error Types

- `DominosValidationError` - Validation error
- `DominosPriceError` - Price error
- `DominosPlaceOrderError` - Order placement error
- `DominosTrackingError` - Tracking error
- `DominosAddressError` - Address error
- `DominosDateError` - Date error
- `DominosStoreError` - Store error
- `DominosProductsError` - Products error

### International Support

```go
// Switch to Canadian endpoints
dominos.UseInternational(dominos.Canada)

// Switch back to USA endpoints
dominos.UseInternational(dominos.USA)
```

## License

MIT 