package models

// Customer represents a Domino's Pizza customer
type Customer struct {
	DominosFormat
	Address     *Address `json:"address"`
	Email       string   `json:"email"`
	FirstName   string   `json:"firstName"`
	LastName    string   `json:"lastName"`
	Phone       string   `json:"phone"`
	PhonePrefix string   `json:"phonePrefix"`
	Extension   string   `json:"extension"`
}

// NewCustomer creates a new customer from a customer data object
func NewCustomer(customerData map[string]interface{}) (*Customer, error) {
	customer := &Customer{}

	// Set customer fields from customerData
	customer.SetFormatted(customerData)

	// Handle address specially (could be a string or object)
	if addrData, ok := customerData["address"]; ok {
		address, err := NewAddress(addrData)
		if err != nil {
			return nil, err
		}
		customer.Address = address
	}

	return customer, nil
}
