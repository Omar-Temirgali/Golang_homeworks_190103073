// Homework 2: Object Oriented Programming
// Due February 7, 2017 at 11:59pm
package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	// Feel free to use the main function for testing your functions
	fmt.Println(Prices)
	RegisterItem(Prices, "banana", 365)
	RegisterItem(Prices, "eggs", 249)
	fmt.Println(Prices)

	var myCart Cart
	myCart.AddItem("eggs")
	myCart.AddItem("peanut butter")
	myCart.AddItem("Milk")
	myCart.AddItem("baNAna")
	myCart.AddItem("juice")
	fmt.Println(myCart.hasMilk())
	fmt.Println(myCart.HasItem("Plant"))
	fmt.Println(myCart.HasItem("EGGS"))
	fmt.Println(myCart)
	myCart.Checkout()
	fmt.Println(myCart)
}

// Price is the cost of something in US cents.
type Price int64

// String is the string representation of a Price
// These should be represented in US Dollars
// Example: 2595 cents => $25.95
func (p Price) String() string {
	return fmt.Sprintf("$%.2f", float64(p)/100)
}

// Prices is a map from an item to its price.
var Prices = map[string]Price{
	"eggs":          219,
	"bread":         199,
	"milk":          295,
	"peanut butter": 445,
	"chocolate":     150,
}

// RegisterItem adds the new item in the prices map.
// If the item is already in the prices map, a warning should be displayed to the user,
// but the value should be overwritten.
// Bonus (1pt) - Use the "log" package to print the error to the user
func RegisterItem(prices map[string]Price, item string, price Price) {
	if _, ok := prices[item]; ok == true {
		log.SetFlags(0)
		log.Printf("%s is already in the map", item)
		prices[item] = price
	} else {
		prices[item] = price
	}
}

// Cart is a struct representing a shopping cart of items.
type Cart struct {
	Items      []string
	TotalPrice Price
}

// hasMilk returns whether the shopping cart has "milk".
func (c *Cart) hasMilk() bool {
	for _, i := range c.Items {
		if strings.ToLower(i) == "milk" {
			return true
		}
	}
	return false
}

// HasItem returns whether the shopping cart has the provided item name.
func (c *Cart) HasItem(item string) bool {
	for _, i := range c.Items {
		if strings.ToLower(i) == strings.ToLower(item) {
			return true
		}
	}
	return false
}

// AddItem adds the provided item to the cart and update the cart balance.
// If item is not found in the prices map, then do not add it and print an error.
// Bonus (1pt) - Use the "log" package to print the error to the user
func (c *Cart) AddItem(item string) {
	item_lower := strings.ToLower(item)
	if _, ok := Prices[item_lower]; ok == true {
		c.Items = append(c.Items, item_lower)
		c.TotalPrice += Prices[item_lower]
	} else {
		log.Printf("%s is not found in the prices map", item_lower)
	}
}

// Checkout displays the final cart balance and clears the cart completely.
func (c *Cart) Checkout() {
	fmt.Println("Checkout:", c.TotalPrice)
	c.Items = nil
	c.TotalPrice = 0
}
