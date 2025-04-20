// Package dominos provides a Go wrapper for the Domino's Pizza API
package dominos

import (
	"github.com/zjpiazza/go-dominos-pizza-api/pkg/models"
	"github.com/zjpiazza/go-dominos-pizza-api/pkg/utils"
)

// Export models
type (
	Address      = models.Address
	Customer     = models.Customer
	Item         = models.Item
	Menu         = models.Menu
	NearbyStores = models.NearbyStores
	Order        = models.Order
	Payment      = models.Payment
	Store        = models.Store
	Tracking     = models.Tracking
)

// Export constructors
var (
	NewAddress      = models.NewAddress
	NewCustomer     = models.NewCustomer
	NewItem         = models.NewItem
	NewNearbyStores = models.NewNearbyStores
	NewOrder        = models.NewOrder
	NewPayment      = models.NewPayment
	NewStore        = models.NewStore
	NewTracking     = models.NewTracking
)

// Export utility functions and values
var (
	// URL configurations
	URLs             = utils.URLs
	USA              = utils.USA
	Canada           = utils.Canada
	UseInternational = utils.UseInternational

	// HTTP utilities
	Get         = utils.Get
	Post        = utils.Post
	GetTracking = utils.GetTracking
)

// Export error constructors
var (
	NewDominosValidationError = utils.NewDominosValidationError
	NewDominosPriceError      = utils.NewDominosPriceError
	NewDominosPlaceOrderError = utils.NewDominosPlaceOrderError
	NewDominosTrackingError   = utils.NewDominosTrackingError
	NewDominosAddressError    = utils.NewDominosAddressError
	NewDominosDateError       = utils.NewDominosDateError
	NewDominosStoreError      = utils.NewDominosStoreError
	NewDominosProductsError   = utils.NewDominosProductsError
)
