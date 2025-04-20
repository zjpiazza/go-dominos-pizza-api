package models

// Menu represents a Domino's Pizza menu
type Menu struct {
	DominosFormat
	Categories            map[string]interface{} `json:"categories"`
	Coupons               map[string]interface{} `json:"coupons"`
	Flavors               map[string]interface{} `json:"flavors"`
	Products              map[string]interface{} `json:"products"`
	Preconfigured         map[string]interface{} `json:"preconfigured"`
	PreconfiguredProducts map[string]interface{} `json:"preconfiguredProducts"`
	Sides                 map[string]interface{} `json:"sides"`
	Toppings              map[string]interface{} `json:"toppings"`
	Variants              map[string]interface{} `json:"variants"`
}

// GetCategory returns a specific category from the menu
func (m *Menu) GetCategory(categoryCode string) (map[string]interface{}, bool) {
	if m.Categories != nil {
		if category, ok := m.Categories[categoryCode]; ok {
			return category.(map[string]interface{}), true
		}
	}
	return nil, false
}

// GetProduct returns a specific product from the menu
func (m *Menu) GetProduct(productCode string) (map[string]interface{}, bool) {
	if m.Products != nil {
		if product, ok := m.Products[productCode]; ok {
			return product.(map[string]interface{}), true
		}
	}
	return nil, false
}

// GetTopping returns a specific topping from the menu
func (m *Menu) GetTopping(toppingCode string) (map[string]interface{}, bool) {
	if m.Toppings != nil {
		if topping, ok := m.Toppings[toppingCode]; ok {
			return topping.(map[string]interface{}), true
		}
	}
	return nil, false
}

// GetVariant returns a specific variant from the menu
func (m *Menu) GetVariant(variantCode string) (map[string]interface{}, bool) {
	if m.Variants != nil {
		if variant, ok := m.Variants[variantCode]; ok {
			return variant.(map[string]interface{}), true
		}
	}
	return nil, false
}
