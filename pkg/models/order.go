package models

import (
	"time"

	"github.com/zjpiazza/go-dominos-pizza-api/pkg/utils"
)

// Order represents a Domino's Pizza order
type Order struct {
	DominosFormat
	Address               *Address               `json:"address"`
	Amounts               map[string]interface{} `json:"amounts"`
	AmountsBreakdown      map[string]interface{} `json:"amountsBreakdown"`
	BusinessDate          string                 `json:"businessDate"`
	Coupons               []interface{}          `json:"coupons"`
	Currency              string                 `json:"currency"`
	CustomerID            string                 `json:"customerID"`
	EstimatedWaitMinutes  string                 `json:"estimatedWaitMinutes"`
	Email                 string                 `json:"email"`
	Extension             string                 `json:"extension"`
	FirstName             string                 `json:"firstName"`
	HotspotsLite          bool                   `json:"hotspotsLite"`
	IP                    string                 `json:"iP"`
	LastName              string                 `json:"lastName"`
	LanguageCode          string                 `json:"languageCode"`
	Market                string                 `json:"market"`
	MetaData              map[string]interface{} `json:"metaData"`
	NewUser               bool                   `json:"newUser"`
	NoCombine             bool                   `json:"noCombine"`
	OrderChannel          string                 `json:"orderChannel"`
	OrderID               string                 `json:"orderID"`
	OrderInfoCollection   []interface{}          `json:"orderInfoCollection"`
	OrderMethod           string                 `json:"orderMethod"`
	OrderTaker            string                 `json:"orderTaker"`
	Partners              map[string]interface{} `json:"partners"`
	Payments              []*Payment             `json:"payments"`
	Phone                 string                 `json:"phone"`
	PhonePrefix           string                 `json:"phonePrefix"`
	PriceOrderMs          int                    `json:"priceOrderMs"`
	PriceOrderTime        string                 `json:"priceOrderTime"`
	Products              []*Item                `json:"products"`
	Promotions            map[string]interface{} `json:"promotions"`
	PulseOrderGuid        string                 `json:"pulseOrderGuid"`
	ServiceMethod         string                 `json:"serviceMethod"`
	SourceOrganizationURI string                 `json:"sourceOrganizationURI"`
	StoreID               string                 `json:"storeID"`
	Tags                  map[string]interface{} `json:"tags"`
	UserAgent             string                 `json:"userAgent"`
	Version               string                 `json:"version"`
	FutureOrderTime       string                 `json:"futureOrderTime,omitempty"`

	validationResponse map[string]interface{}
	priceResponse      map[string]interface{}
	placeResponse      map[string]interface{}
}

// NewOrder creates a new order with the given customer
func NewOrder(customer *Customer) *Order {
	order := &Order{
		Address:               customer.Address,
		Coupons:               make([]interface{}, 0),
		Email:                 customer.Email,
		Extension:             customer.Extension,
		FirstName:             customer.FirstName,
		HotspotsLite:          false,
		LastName:              customer.LastName,
		LanguageCode:          "en",
		MetaData:              map[string]interface{}{"calculateNutrition": true, "contactless": true},
		NewUser:               true,
		NoCombine:             true,
		OrderChannel:          "OLO",
		OrderInfoCollection:   make([]interface{}, 0),
		OrderMethod:           "Web",
		OrderTaker:            "go-dominos-pizza-api",
		Partners:              make(map[string]interface{}),
		Payments:              make([]*Payment, 0),
		Phone:                 customer.Phone,
		PhonePrefix:           customer.PhonePrefix,
		Products:              make([]*Item, 0),
		ServiceMethod:         "Delivery",
		SourceOrganizationURI: utils.URLs.SourceURI,
		Tags:                  make(map[string]interface{}),
		Version:               "1.0",
	}

	return order
}

// OrderInFuture sets the order for a future time
func (o *Order) OrderInFuture(futureTime time.Time) error {
	now := time.Now()
	if futureTime.Before(now) {
		return utils.NewDominosDateError("Order dates must be in the future")
	}

	// Format the time for Domino's API
	dateString := futureTime.Format("2006-01-02 15:04:05")
	o.FutureOrderTime = dateString

	return nil
}

// OrderNow removes any future order time setting
func (o *Order) OrderNow() {
	o.FutureOrderTime = ""
}

// AddCoupon adds a coupon to the order
func (o *Order) AddCoupon(coupon interface{}) *Order {
	o.Coupons = append(o.Coupons, coupon)
	return o
}

// RemoveCoupon removes a coupon from the order
func (o *Order) RemoveCoupon(coupon interface{}) *Order {
	for i, c := range o.Coupons {
		if c == coupon {
			o.Coupons = append(o.Coupons[:i], o.Coupons[i+1:]...)
			break
		}
	}
	return o
}

// AddItem adds an item to the order
func (o *Order) AddItem(item *Item) *Order {
	o.Products = append(o.Products, item)
	return o
}

// RemoveItem removes an item from the order
func (o *Order) RemoveItem(item *Item) *Order {
	for i, p := range o.Products {
		if p == item {
			o.Products = append(o.Products[:i], o.Products[i+1:]...)
			break
		}
	}
	return o
}

// GetValidationResponse returns the response from the last validation
func (o *Order) GetValidationResponse() map[string]interface{} {
	return o.validationResponse
}

// GetPriceResponse returns the response from the last price check
func (o *Order) GetPriceResponse() map[string]interface{} {
	return o.priceResponse
}

// GetPlaceResponse returns the response from the last order placement
func (o *Order) GetPlaceResponse() map[string]interface{} {
	return o.placeResponse
}

// Validate validates the order with Domino's API
func (o *Order) Validate() error {
	if o.StoreID == "" {
		return utils.NewDominosStoreError("Store ID must be set before validating order")
	}

	// Create payload
	payload := map[string]interface{}{
		"Order": o.GetFormatted(),
	}

	// Add address
	if o.Address != nil {
		payload["Order"].(map[string]interface{})["Address"] = o.Address.GetFormatted()
	}

	// Add products
	products := make([]map[string]interface{}, len(o.Products))
	for i, item := range o.Products {
		products[i] = item.GetFormatted()
	}
	payload["Order"].(map[string]interface{})["Products"] = products

	// Add payments
	payments := make([]map[string]interface{}, len(o.Payments))
	for i, payment := range o.Payments {
		payments[i] = payment.GetFormatted()
	}
	payload["Order"].(map[string]interface{})["Payments"] = payments

	// Send validation request
	response, err := utils.Post(utils.URLs.Order.Validate, payload)
	if err != nil {
		return err
	}

	o.validationResponse = response

	// Check for errors
	orderStatus, ok := response["Order"].(map[string]interface{})["Status"]
	responseStatus, respOk := response["Status"]

	if (ok && orderStatus.(float64) == -1) || (respOk && responseStatus.(float64) == -1) {
		return utils.NewDominosValidationError(response)
	}

	// Update order with validated data
	if orderData, ok := response["Order"].(map[string]interface{}); ok {
		o.SetFormatted(orderData)
	}

	return nil
}

// Price gets the price for the order from Domino's API
func (o *Order) Price() error {
	if o.StoreID == "" {
		return utils.NewDominosStoreError("Store ID must be set before pricing an order")
	}

	if len(o.Products) == 0 {
		return utils.NewDominosProductsError("Order must contain product items before pricing")
	}

	// Create payload
	payload := map[string]interface{}{
		"Order": o.GetFormatted(),
	}

	// Add address
	if o.Address != nil {
		payload["Order"].(map[string]interface{})["Address"] = o.Address.GetFormatted()
	}

	// Add products
	products := make([]map[string]interface{}, len(o.Products))
	for i, item := range o.Products {
		products[i] = item.GetFormatted()
	}
	payload["Order"].(map[string]interface{})["Products"] = products

	// Add payments
	payments := make([]map[string]interface{}, len(o.Payments))
	for i, payment := range o.Payments {
		payments[i] = payment.GetFormatted()
	}
	payload["Order"].(map[string]interface{})["Payments"] = payments

	// Send price request
	response, err := utils.Post(utils.URLs.Order.Price, payload)
	if err != nil {
		return err
	}

	o.priceResponse = response

	// Check for errors
	orderStatus, ok := response["Order"].(map[string]interface{})["Status"]
	responseStatus, respOk := response["Status"]

	if (ok && orderStatus.(float64) == -1) || (respOk && responseStatus.(float64) == -1) {
		return utils.NewDominosPriceError(response)
	}

	// Update order with priced data
	if orderData, ok := response["Order"].(map[string]interface{}); ok {
		o.SetFormatted(orderData)
	}

	return nil
}

// Place places the order with Domino's API
func (o *Order) Place() error {
	if o.StoreID == "" {
		return utils.NewDominosStoreError("Store ID must be set before placing an order")
	}

	if len(o.Products) == 0 {
		return utils.NewDominosProductsError("Order must contain product items before placing")
	}

	if o.Address == nil || o.Address.Region == "" {
		return utils.NewDominosAddressError("Order must have a valid address before placing")
	}

	if len(o.Payments) == 0 {
		return utils.NewDominosProductsError("Order must have at least one payment method")
	}

	// Create payload
	payload := map[string]interface{}{
		"Order": o.GetFormatted(),
	}

	// Add address
	payload["Order"].(map[string]interface{})["Address"] = o.Address.GetFormatted()

	// Add products
	products := make([]map[string]interface{}, len(o.Products))
	for i, item := range o.Products {
		products[i] = item.GetFormatted()
	}
	payload["Order"].(map[string]interface{})["Products"] = products

	// Add payments
	payments := make([]map[string]interface{}, len(o.Payments))
	for i, payment := range o.Payments {
		payments[i] = payment.GetFormatted()
	}
	payload["Order"].(map[string]interface{})["Payments"] = payments

	// Send place order request
	response, err := utils.Post(utils.URLs.Order.Place, payload)
	if err != nil {
		return err
	}

	o.placeResponse = response

	// Check for errors
	orderStatus, ok := response["Order"].(map[string]interface{})["Status"]
	responseStatus, respOk := response["Status"]

	if (ok && orderStatus.(float64) == -1) || (respOk && responseStatus.(float64) == -1) {
		return utils.NewDominosPlaceOrderError(response)
	}

	// Update order with placed data
	if orderData, ok := response["Order"].(map[string]interface{}); ok {
		o.SetFormatted(orderData)
	}

	return nil
}
