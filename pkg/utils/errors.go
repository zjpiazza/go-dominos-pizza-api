package utils

import (
	"fmt"
)

// DominosError is the base error type for all Domino's API errors
type DominosError struct {
	Message string
	Details interface{}
}

func (e *DominosError) Error() string {
	return fmt.Sprintf("%s: %v", e.Message, e.Details)
}

// DominosValidationError represents an error during order validation
type DominosValidationError struct {
	DominosError
}

// NewDominosValidationError creates a new validation error
func NewDominosValidationError(details interface{}) *DominosValidationError {
	return &DominosValidationError{
		DominosError: DominosError{
			Message: "Validation failed",
			Details: details,
		},
	}
}

// DominosPriceError represents an error during order pricing
type DominosPriceError struct {
	DominosError
}

// NewDominosPriceError creates a new price error
func NewDominosPriceError(details interface{}) *DominosPriceError {
	return &DominosPriceError{
		DominosError: DominosError{
			Message: "Price calculation failed",
			Details: details,
		},
	}
}

// DominosPlaceOrderError represents an error during order placement
type DominosPlaceOrderError struct {
	DominosError
}

// NewDominosPlaceOrderError creates a new place order error
func NewDominosPlaceOrderError(details interface{}) *DominosPlaceOrderError {
	return &DominosPlaceOrderError{
		DominosError: DominosError{
			Message: "Order placement failed",
			Details: details,
		},
	}
}

// DominosTrackingError represents an error during order tracking
type DominosTrackingError struct {
	DominosError
}

// NewDominosTrackingError creates a new tracking error
func NewDominosTrackingError(details interface{}) *DominosTrackingError {
	return &DominosTrackingError{
		DominosError: DominosError{
			Message: "Order tracking failed",
			Details: details,
		},
	}
}

// DominosAddressError represents an error with an address
type DominosAddressError struct {
	DominosError
}

// NewDominosAddressError creates a new address error
func NewDominosAddressError(details interface{}) *DominosAddressError {
	return &DominosAddressError{
		DominosError: DominosError{
			Message: "Invalid address",
			Details: details,
		},
	}
}

// DominosDateError represents an error with a date
type DominosDateError struct {
	DominosError
}

// NewDominosDateError creates a new date error
func NewDominosDateError(details interface{}) *DominosDateError {
	return &DominosDateError{
		DominosError: DominosError{
			Message: "Invalid date",
			Details: details,
		},
	}
}

// DominosStoreError represents an error with a store
type DominosStoreError struct {
	DominosError
}

// NewDominosStoreError creates a new store error
func NewDominosStoreError(details interface{}) *DominosStoreError {
	return &DominosStoreError{
		DominosError: DominosError{
			Message: "Store error",
			Details: details,
		},
	}
}

// DominosProductsError represents an error with products
type DominosProductsError struct {
	DominosError
}

// NewDominosProductsError creates a new products error
func NewDominosProductsError(details interface{}) *DominosProductsError {
	return &DominosProductsError{
		DominosError: DominosError{
			Message: "Products error",
			Details: details,
		},
	}
}
