package main

import (
	"fmt"

	"github.com/flodht/Motorsport-Calendar/entity"
)

// init the main function
func main() {
	// init the grand prix
	grand_prix := entity.GrandPrix{
		Location: "Monza",
		Date: "2021-09-12",
		Name: "Italian Grand Prix",
	}

	// init the category
	category := entity.Category{
		Name: "Formula 1",
	}

	// print the grand prix
	fmt.Println(grand_prix)

	// print the category
	fmt.Println(category)
}
