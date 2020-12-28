package main

import "fmt"

type car struct {
	model       string
	createdAt   int
	from        string
	velocityMax int
}

func newCar(model string) *car {
	c := car{model: model}

	c.createdAt = 1990
	c.from = "Japan"
	c.velocityMax = 80

	return &c
}

func main() {
	fmt.Println(car{
		model:       "Mazda Rx5",
		createdAt:   2000,
		from:        "Japan",
		velocityMax: 280,
	})

	fmt.Println(newCar("Nissan"))
}
