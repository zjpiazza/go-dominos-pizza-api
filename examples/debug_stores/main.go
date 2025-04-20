package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	// Try simple approach with just a zip code first (often more reliable)
	fmt.Println("Trying simple approach with zip code only...")
	zipURL := "https://order.dominos.com/power/store-locator?s=90210&type=Delivery"
	fmt.Printf("Zip Code URL: %s\n\n", zipURL)

	testURL(zipURL, "Zip Code Test")

	// If needed, try with full address
	fmt.Println("\n\nTrying with full address...")
	// The address information
	street := "9876 Wilshire Blvd"
	cityStateZip := "Beverly Hills, CA 90210"

	// URL encode the parameters
	encodedStreet := url.QueryEscape(street)
	encodedCityStateZip := url.QueryEscape(cityStateZip)

	// Create the URL directly - bypassing the package for debugging
	storeFinderURL := fmt.Sprintf("https://order.dominos.com/power/store-locator?s=%s&c=%s&type=Delivery",
		encodedStreet, encodedCityStateZip)

	fmt.Printf("Full Address URL: %s\n\n", storeFinderURL)

	testURL(storeFinderURL, "Full Address Test")

	// Try alternative URL formats
	fmt.Println("\n\nTrying different URL format...")

	// Just address, no city
	streetOnlyURL := fmt.Sprintf("https://order.dominos.com/power/store-locator?s=%s&type=Delivery",
		encodedStreet)

	fmt.Printf("Street-only URL: %s\n\n", streetOnlyURL)

	testURL(streetOnlyURL, "Street-only Test")
}

func testURL(url string, testName string) {
	// Create an HTTP client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Create a request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error creating request for %s: %v", testName, err)
	}

	// Set headers that might be required
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Origin", "https://order.dominos.com")
	req.Header.Set("Referer", "https://order.dominos.com/")

	// Send the request
	fmt.Printf("Sending request for %s...\n", testName)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request for %s: %v", testName, err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response for %s: %v", testName, err)
	}

	// Check the status code
	fmt.Printf("Status code for %s: %d\n", testName, resp.StatusCode)

	// If it's JSON, pretty print it
	var prettyJSON bytes.Buffer
	contentType := resp.Header.Get("Content-Type")
	fmt.Printf("Content-Type: %s\n", contentType)

	if contentType == "application/json" || contentType == "application/json; charset=utf-8" {
		err = json.Indent(&prettyJSON, body, "", "  ")
		if err == nil {
			fmt.Printf("API Response for %s (JSON):\n", testName)
			fmt.Println(prettyJSON.String())
		} else {
			fmt.Printf("API Response for %s (Raw, couldn't parse as JSON):\n", testName)
			fmt.Println(string(body))
		}
	} else {
		fmt.Printf("API Response for %s (Raw):\n", testName)
		fmt.Println(string(body))
	}

	// Check if it contains stores data
	var jsonData map[string]interface{}
	if err := json.Unmarshal(body, &jsonData); err == nil {
		if stores, ok := jsonData["Stores"].([]interface{}); ok {
			fmt.Printf("\nFound %d stores in the response\n", len(stores))

			if len(stores) > 0 {
				// Print the first store
				storeObj := stores[0].(map[string]interface{})
				fmt.Println("\nFirst store details:")
				fmt.Printf("  StoreID: %v\n", storeObj["StoreID"])
				fmt.Printf("  Phone: %v\n", storeObj["Phone"])
				fmt.Printf("  MinDistance: %v\n", storeObj["MinDistance"])
			}
		} else {
			fmt.Println("\nNo stores found in the response")
		}
	}
}
