package order

import "fmt"

var (
	ORDER_NAME = "MY AWESOME ORDER"
)

const (
	specialPrice = 100.0
)

type Order struct {
	ID       int
	Price    float64
	Discount float64
	metadata any // tidak bisa digunakan di luar package order
}

func (o Order) calculateDiscount() float64 {
	return o.Price * o.Discount
}
func (o Order) FinalPrice() float64 {
	return o.calculateDiscount() + specialPrice
}

func init() {
	// function yang dijalankan oleh GO
	// PERTAMA KALI sebelum func MAIN
	fmt.Println("FROM ORDER")
}
