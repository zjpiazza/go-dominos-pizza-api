package models

import (
	"strings"

	"github.com/zjpiazza/go-dominos-pizza-api/pkg/utils"
)

// Address represents a delivery or pickup address
type Address struct {
	DominosFormat
	Street       string `json:"street"`
	StreetNumber string `json:"streetNumber"`
	StreetName   string `json:"streetName"`
	City         string `json:"city"`
	Region       string `json:"region"`
	PostalCode   string `json:"postalCode"`
	Type         string `json:"type"` // House, Apartment, Business, etc.
	UnitType     string `json:"unitType"`
	UnitNumber   string `json:"unitNumber"`
}

// NewAddress creates a new address from a string or address object
func NewAddress(addr interface{}) (*Address, error) {
	address := &Address{
		Type: "House", // Default type
	}

	switch a := addr.(type) {
	case string:
		// Parse the address string
		return parseAddressString(a)
	case map[string]interface{}:
		// Process address object
		address.SetFormatted(a)
	}

	return address, nil
}

// parseAddressString parses an address string into an Address object
func parseAddressString(addrStr string) (*Address, error) {
	address := &Address{
		Type: "House", // Default type
	}

	// Split the address string by commas
	parts := strings.Split(addrStr, ",")
	if len(parts) < 3 {
		return nil, utils.NewDominosAddressError("Address string must have at least 3 parts separated by commas")
	}

	// Process the parts
	address.Street = strings.TrimSpace(parts[0])
	address.City = strings.TrimSpace(parts[1])

	// Region and postal code
	if len(parts) >= 3 {
		regionParts := strings.Split(strings.TrimSpace(parts[2]), " ")
		if len(regionParts) > 0 {
			address.Region = strings.TrimSpace(regionParts[0])
		}
		if len(regionParts) > 1 {
			address.PostalCode = strings.Join(regionParts[1:], " ")
		}
	}

	// Extract street number and name
	streetParts := strings.SplitN(address.Street, " ", 2)
	if len(streetParts) >= 2 {
		address.StreetNumber = streetParts[0]
		address.StreetName = streetParts[1]
	} else {
		address.StreetName = address.Street
	}

	return address, nil
}

// GetDefaultLineOne returns the default value for line 1 of the address
func (a *Address) GetDefaultLineOne() string {
	if a.Street != "" {
		return a.Street
	}

	// If street isn't set but street number and name are, put them together
	if a.StreetNumber != "" && a.StreetName != "" {
		return a.StreetNumber + " " + a.StreetName
	}

	return ""
}

// GetDefaultLineTwo returns the default value for line 2 of the address
func (a *Address) GetDefaultLineTwo() string {
	if a.City != "" && a.Region != "" && a.PostalCode != "" {
		return a.City + ", " + a.Region + " " + a.PostalCode
	}
	return ""
}
