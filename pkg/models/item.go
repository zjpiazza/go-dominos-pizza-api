package models

// Item represents a product item in an order
type Item struct {
	DominosFormat
	Code         string                 `json:"code"`
	ID           int                    `json:"id"`
	Qty          int                    `json:"qty"`
	CategoryCode string                 `json:"categoryCode"`
	FlavorCode   string                 `json:"flavorCode"`
	Options      map[string]interface{} `json:"options"`
	IsNew        bool                   `json:"isNew"`
}

// NewItem creates a new item from item data
func NewItem(itemData map[string]interface{}) (*Item, error) {
	item := &Item{
		Qty:   1,    // Default quantity
		IsNew: true, // Default is new
	}

	// Set item fields from itemData
	item.SetFormatted(itemData)

	// Default missing values
	if item.Code == "" {
		if code, ok := itemData["code"]; ok {
			item.Code = code.(string)
		}
	}

	// Ensure options is initialized
	if item.Options == nil {
		item.Options = make(map[string]interface{})
	}

	return item, nil
}
