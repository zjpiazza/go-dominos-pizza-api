package utils

import (
	"reflect"
	"strings"
	"unicode"
)

// ToCamel converts a string to camelCase
func ToCamel(s string) string {
	if s == "" {
		return s
	}

	result := ""
	nextUpper := false
	for i, c := range s {
		if c == '_' || c == '-' || c == ' ' {
			nextUpper = true
			continue
		}

		if i == 0 {
			result += string(unicode.ToLower(c))
		} else if nextUpper {
			result += string(unicode.ToUpper(c))
			nextUpper = false
		} else {
			result += string(c)
		}
	}

	return result
}

// ToPascal converts a string to PascalCase
func ToPascal(s string) string {
	if s == "" {
		return s
	}

	camelCased := ToCamel(s)
	return strings.ToUpper(camelCased[0:1]) + camelCased[1:]
}

// CamelObjectKeys recursively converts all keys in a map to camelCase
func CamelObjectKeys(obj interface{}) interface{} {
	if obj == nil {
		return nil
	}

	// Process according to the type of obj
	switch objValue := obj.(type) {
	case map[string]interface{}:
		newMap := make(map[string]interface{})
		for key, value := range objValue {
			newMap[ToCamel(key)] = CamelObjectKeys(value)
		}
		return newMap

	case []interface{}:
		newArray := make([]interface{}, len(objValue))
		for i, value := range objValue {
			newArray[i] = CamelObjectKeys(value)
		}
		return newArray

	default:
		return obj
	}
}

// PascalObjectKeys recursively converts all keys in a map to PascalCase
func PascalObjectKeys(obj interface{}) interface{} {
	if obj == nil {
		return nil
	}

	// Process according to the type of obj
	switch objValue := obj.(type) {
	case map[string]interface{}:
		newMap := make(map[string]interface{})
		for key, value := range objValue {
			newMap[ToPascal(key)] = PascalObjectKeys(value)
		}
		return newMap

	case []interface{}:
		newArray := make([]interface{}, len(objValue))
		for i, value := range objValue {
			newArray[i] = PascalObjectKeys(value)
		}
		return newArray

	default:
		return obj
	}
}

// StructToPascalMap converts a struct to a map with PascalCase keys
func StructToPascalMap(obj interface{}) map[string]interface{} {
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return nil
	}

	result := make(map[string]interface{})
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Name

		// Skip unexported fields
		if !field.CanInterface() {
			continue
		}

		fieldValue := field.Interface()
		result[ToPascal(fieldName)] = fieldValue
	}

	return result
}

// StructToCamelMap converts a struct to a map with camelCase keys
func StructToCamelMap(obj interface{}) map[string]interface{} {
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return nil
	}

	result := make(map[string]interface{})
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Name

		// Skip unexported fields
		if !field.CanInterface() {
			continue
		}

		fieldValue := field.Interface()
		result[ToCamel(fieldName)] = fieldValue
	}

	return result
}
