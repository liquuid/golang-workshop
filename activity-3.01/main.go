package main

import "fmt"

func main() {
	type Product struct {
		item string
		cost float32
		tax  float32
	}
	cart := []Product{}

	cart = append(cart, Product{"Cake", 0.99, 7.5})
	cart = append(cart, Product{"Milk", 2.75, 1.5})
	cart = append(cart, Product{"Butter", 0.87, 2})

	var totalTaxes float32

	for _, item := range cart {
		totalTaxes = totalTaxes + (item.cost * item.tax / 100.0)
	}

	fmt.Println("Sales Tax Total: ", totalTaxes)

}
