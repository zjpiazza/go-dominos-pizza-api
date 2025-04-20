# Domino's Pizza Menu API Example

This example demonstrates how to use the Domino's Pizza API to retrieve and work with the menu data from a specific store.

## Features

- Retrieves the complete menu from a Domino's Pizza store
- Categorizes menu items into pizzas, sides, and drinks
- Shows how to access menu items by their product codes
- Demonstrates the various types of products available

## How to Use

Run the example with:

```bash
go run main.go
```

By default, the example uses a store ID from Beverly Hills, CA. You can modify the `storeID` variable in the code to use a different store.

## Menu Structure

The menu is organized by product types with the following prefixes:

- `S_*`: Specialty Pizzas and other pizza-related items
- `F_*`: Sides, Drinks, and other food items
- `B_*`: (Rarely used) Some beverage items

## Working with Menu Items

To get details about a specific menu item, you can use:

```go
// Get a specific product by its code
product, found := menu.GetProduct("S_PIZPV")  // Pacific Veggie Pizza
if found {
    name := product["Name"].(string)
    // Work with the product details
}

// Get all pizzas
pizzas := menu.GetPizzas()

// Get all sides (excludes drinks)
sides := menu.GetSides()

// Get all drinks
drinks := menu.GetDrinks()
```

## Example Output

The example prints the first 5 items from each category (pizzas, sides, and drinks) along with the total count of items in each category. 