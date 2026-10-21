package main

import (
	"errors"
	"fmt"
)

type Status string

const (
	StatusCreated Status = "created"
	StatusPaid    Status = "paid"
	StatusFailed  Status = "failed"
)

type Item struct {
	Name       string
	PriceCents int64
	Quantity   int
}

type Order struct {
	ID       string
	Customer string
	Items    []Item
	Status   Status
}

func (o Order) Total() int64 {
	var total int64
	for _, item := range o.Items {
		n := item.PriceCents * int64(item.Quantity)
		total += n
	}
	return total
}

func FormatPrice(cents int64) string {
	return fmt.Sprintf("%d,%02d ₽", cents/100, cents%100)
}

func NewOrder(id string, customer string, items []Item) (Order, error) {
	if customer == "" {
		return Order{}, errors.New("имя не указано")
	}
	if len(items) == 0 {
		return Order{}, errors.New("товаров нет")
	}
	order := Order{ID: id, Customer: customer, Items: items, Status: StatusCreated}
	return order, nil
}

func main() {
	items := []Item{
		{Name: "Кофе", PriceCents: 25000, Quantity: 2},
		{Name: "Круассан", PriceCents: 18050, Quantity: 1},
	}
	order, err := NewOrder("1", "Николай", items)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(FormatPrice(order.Total()))
}
