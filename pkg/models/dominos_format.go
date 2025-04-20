package models

import (
	"encoding/json"

	"github.com/zjpiazza/go-dominos-pizza-api/pkg/utils"
)

// DominosFormat is the base struct that provides common formatting functionality
type DominosFormat struct {
	dominosAPIResponse map[string]interface{}
}

// GetDominosAPIResponse returns the raw API response
func (df *DominosFormat) GetDominosAPIResponse() map[string]interface{} {
	return df.dominosAPIResponse
}

// SetDominosAPIResponse sets the raw API response
func (df *DominosFormat) SetDominosAPIResponse(response map[string]interface{}) {
	df.dominosAPIResponse = response
}

// GetFormatted returns the struct as a map with PascalCase keys
func (df *DominosFormat) GetFormatted() map[string]interface{} {
	// Convert the struct to a map
	data, _ := json.Marshal(df)
	var objMap map[string]interface{}
	json.Unmarshal(data, &objMap)

	// Remove unexported fields like dominosAPIResponse
	delete(objMap, "dominosAPIResponse")

	// Convert to Pascal case
	pascalMap, _ := utils.PascalObjectKeys(objMap).(map[string]interface{})
	return pascalMap
}

// SetFormatted updates the struct from a map with keys in any format
func (df *DominosFormat) SetFormatted(data map[string]interface{}) {
	// Convert the keys to camelCase
	camelMap, _ := utils.CamelObjectKeys(data).(map[string]interface{})

	// Convert map to JSON
	jsonData, _ := json.Marshal(camelMap)

	// Unmarshal JSON into the struct
	json.Unmarshal(jsonData, df)
}
