package main

import "fmt"

type Status string

const (
	StatusCreated Status = "created"
	StatusPaid Status = "paid"
	StatusFailed Status = "failed"
)

type Item struct {
	Name	string
	PriceCents	int64
	Quantity	int
}

type Order struct {
	ID	string
	Customer	string
	Items	[]Item
	Status	Status
}

func (o Order) Total() int64 {
	var total int64
	for _, item := range o.Items {
		n := item.PriceCents * int64(item.Quantity)
		total += n
	}
	return total
}

func main() {
	order := Order{
		ID: "1",
		Customer: "Иван",
		Status: StatusCreated,
		Items: []Item{
			{Name: "Кофе", PriceCents: 25000, Quantity: 2},
			{Name: "Круассан", PriceCents: 18050, Quantity: 1},
		},
	}
	fmt.Println(order.Total())
}
