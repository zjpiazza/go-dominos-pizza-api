package models

import (
	"strings"
)

// Menu represents a Domino's Pizza menu
type Menu struct {
	DominosFormat
	Categories               map[string]interface{} `json:"categories"`
	Coupons                  map[string]interface{} `json:"coupons"`
	Flavors                  map[string]interface{} `json:"flavors"`
	Products                 map[string]interface{} `json:"products"`
	Preconfigured            map[string]interface{} `json:"preconfigured"`
	PreconfiguredProducts    map[string]interface{} `json:"preconfiguredProducts"`
	Sides                    map[string]interface{} `json:"sides"`
	Toppings                 map[string]interface{} `json:"toppings"`
	Variants                 map[string]interface{} `json:"variants"`
	Sizes                    map[string]interface{} `json:"sizes"`
	Categorization           map[string]interface{} `json:"categorization"`
	AlternativeProductNames  map[string]interface{} `json:"alternativeProductNames"`
	ShortProductDescriptions map[string]interface{} `json:"shortProductDescriptions"`
	ExcludedProducts         map[string]interface{} `json:"excludedProducts"`
	ExcludedOptions          map[string]interface{} `json:"excludedOptions"`
	CookingInstructions      map[string]interface{} `json:"cookingInstructions"`
	CookingInstructionGroups map[string]interface{} `json:"cookingInstructionGroups"`
}

// ProductCategory represents a category of products like Pizzas, Sides, Drinks, etc.
type ProductCategory struct {
	Code        string
	Name        string
	Description string
	Products    []string
}

// Common drink keywords for filtering
var drinkKeywords = []string{
	"coke", "sprite", "pepsi", "water",
	"soda", "beverage", "drink", "cola",
	"fanta", "dr pepper", "dasani",
	"orange", "diet coke", "bottle water",
}

// Foods that might match drink keywords but aren't drinks
var drinkExclusions = []string{
	"lava cake", "chocolate lava", "lava crunch",
}

// GetCategory returns a specific category from the menu
func (m *Menu) GetCategory(categoryCode string) (map[string]interface{}, bool) {
	if m.Categorization != nil {
		food, ok := m.Categorization["Food"].(map[string]interface{})
		if ok {
			if categories, ok := food["Categories"].(map[string]interface{}); ok {
				if category, ok := categories[categoryCode]; ok {
					return category.(map[string]interface{}), true
				}
			}
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

// GetMenuCategories returns all food categories from the menu
func (m *Menu) GetMenuCategories() map[string]interface{} {
	// Get the raw response
	rawResponse := m.GetDominosAPIResponse()

	if rawResponse != nil {
		if categorization, ok := rawResponse["Categorization"].(map[string]interface{}); ok {
			// If we found Categorization, loop through its keys to find Food categories
			for _, catSection := range []string{"Food", "Coupon", "Preconfigured"} {
				if food, ok := categorization[catSection].(map[string]interface{}); ok {
					if categories, ok := food["Categories"].(map[string]interface{}); ok {
						return categories
					}
				}
			}
		}
	}

	// If struct fields are populated, use them as fallback
	if m.Categorization != nil {
		if food, ok := m.Categorization["Food"].(map[string]interface{}); ok {
			if categories, ok := food["Categories"].(map[string]interface{}); ok {
				return categories
			}
		}
	}

	return make(map[string]interface{})
}

// GetRawProducts returns products directly from the raw API response
func (m *Menu) GetRawProducts() map[string]interface{} {
	rawResponse := m.GetDominosAPIResponse()
	if rawResponse != nil {
		if products, ok := rawResponse["Products"].(map[string]interface{}); ok {
			return products
		}
	}

	// Fallback to struct fields
	if m.Products != nil {
		return m.Products
	}

	return make(map[string]interface{})
}

// GetProductsByType returns all products that match a specific type prefix
// Common types:
// - S_* for Specialty Pizzas
// - P_* for Pizzas
// - F_* for Sides and Other Food Items
// - B_* for Beverages/Drinks
func (m *Menu) GetProductsByType(prefix string) map[string]map[string]interface{} {
	result := make(map[string]map[string]interface{})

	products := m.GetRawProducts()
	for code, productData := range products {
		if len(code) >= 2 && code[0:2] == prefix {
			if product, ok := productData.(map[string]interface{}); ok {
				result[code] = product
			}
		}
	}

	return result
}

// GetPizzas returns all pizza products
func (m *Menu) GetPizzas() map[string]map[string]interface{} {
	return m.GetProductsByType("S_")
}

// GetSides returns side items like wings, bread, etc.
func (m *Menu) GetSides() map[string]map[string]interface{} {
	result := make(map[string]map[string]interface{})

	// Get F_ products but exclude drinks
	products := m.GetProductsByType("F_")

	for code, product := range products {
		name, ok := product["Name"].(string)
		if !ok {
			result[code] = product // Keep it if we can't determine what it is
			continue
		}

		// Check if the product name contains drink keywords
		nameLower := strings.ToLower(name)
		isDrink := false
		for _, keyword := range drinkKeywords {
			if strings.Contains(nameLower, keyword) {
				// Check exclusions - items that might match drink keywords but aren't drinks
				isExcluded := false
				for _, exclusion := range drinkExclusions {
					if strings.Contains(nameLower, exclusion) {
						isExcluded = true
						break
					}
				}

				if !isExcluded {
					isDrink = true
					break
				}
			}
		}

		// Only add to result if it's not a drink
		if !isDrink {
			result[code] = product
		}
	}

	return result
}

// GetDrinks returns beverage products
func (m *Menu) GetDrinks() map[string]map[string]interface{} {
	result := make(map[string]map[string]interface{})

	// Drink products typically use F_ prefix but have specific keywords in their names
	products := m.GetRawProducts()

	for code, productData := range products {
		product, ok := productData.(map[string]interface{})
		if !ok {
			continue
		}

		name, ok := product["Name"].(string)
		if !ok {
			continue
		}

		// Check if the product name contains drink keywords
		nameLower := strings.ToLower(name)

		// Skip exclusions - items that might match drink keywords but aren't drinks
		isExcluded := false
		for _, exclusion := range drinkExclusions {
			if strings.Contains(nameLower, exclusion) {
				isExcluded = true
				break
			}
		}

		if isExcluded {
			continue
		}

		// Check for drink keywords
		for _, keyword := range drinkKeywords {
			if strings.Contains(nameLower, keyword) {
				result[code] = product
				break
			}
		}
	}

	return result
}
