package main

import "fmt"

type Order struct {
	ID       string
	Price    float64
	Quantity int
}

func (o Order) getTotal() float64{
	return o.Price * float64(o.Quantity)
}

func main() {
	order := Order{
		ID:       "123",
		Price:    10.0,
		Quantity: 5,
	}

	fmt.Println(order.ID,order.Price,order.Quantity,order.getTotal())
}