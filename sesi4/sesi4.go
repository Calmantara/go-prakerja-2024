package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

type Address struct {
	City    string
	ZipCode int
}

type A interface {
	Get()
	Count()
}

type Order struct{}

func (o Order) Get()   {}
func (o Order) Count() {}

type Product struct{}

func (p *Product) Get()   {}
func (p *Product) Count() {}

type Database interface {
	GetAll()
}

type postgres struct{}

func (pg postgres) GetAll() {
	fmt.Println("form postgres")
}

type mongo struct{}

func (mg mongo) GetAll() {
	fmt.Println("form mongo")
}

func CalculateOrderFromDatabase(db Database) {
	db.GetAll()
}

func main() {
	// I Interface
	var iface any
	iface = 10
	fmt.Println("interface to int", iface)

	iface = "string"
	fmt.Println("interface to string", iface)
	// slice and map
	// interface as value of map => merepresentasikan json
	data := map[string]any{
		"name":  "calman",
		"age":   10,
		"email": "calman@gmail.com",
		"address": Address{
			City:    "Jakarta",
			ZipCode: 123456,
		},
		"metadata": map[string]any{
			"amount": 3,
			"name":   "milk",
		},
		"fn": func() {},
	}
	fmt.Println(data["name"], data["age"])

	slc := []any{
		1,
		"calman",
	}
	for _, val := range slc {
		fmt.Println(val)
	}
	// validate type
	var iface2 any
	iface2 = 100
	fmt.Printf("%T\n", iface2)
	// type casting
	num := iface2.(int)
	fmt.Println(num)

	// reflect type
	iface2 = "string"
	num2, ok := iface2.(int)
	if ok {
		fmt.Println(num2)
	} else {
		fmt.Println("not int")
	}
	str, ok := iface2.(string)
	if ok {
		fmt.Println(str)
	}

	// method interface
	// var stc A
	ord := NewOrder()
	ordStc := ord.(Order)
	// prd := stc.(*Product)
	fmt.Printf("%T\n", ordStc)

	// Depedency Injection
	pg := postgres{}
	// mg := mongo{}
	CalculateOrderFromDatabase(pg)

	// II Go Routine
	fmt.Println(runtime.NumGoroutine())
	// go PrintNumbers()
	fmt.Println(runtime.NumGoroutine())
	// go PrintNumbers()
	fmt.Println(runtime.NumGoroutine())

	// []users{UserA, UserB}
	// kalkulasi nilai akhir user A dan user B
	// tidak ada hubungan antara user A dan user B

	// go calculate(UserA)
	// go calculate(UserB)
	// time.Sleep(50000 * time.Nanosecond)
	// III Wait Group | Error Group
	wg := sync.WaitGroup{}
	wg.Add(2)
	tn := time.Now()
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("from process 1: finished")
		wg.Done()
	}()

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("from process 2: finished")
		wg.Done()
	}()
	wg.Wait()

	wg2 := sync.WaitGroup{}
	intSlice := []int{1, 2, 3, 4, 5, 12, 3, 1, 2, 5, 1, 2, 3, 1, 2, 5, 6, 1, 2, 3}
	for i := 0; i < len(intSlice); i++ {
		wg2.Add(1)
		go func() {
			fmt.Println("from go routine", i)
			wg2.Done()
		}()
	}
	wg2.Wait()

	// apa yang terjadi kalau
	// kita kasih wg tidak sesuai dengan jumlahnya
	wg3 := sync.WaitGroup{}
	wg3.Add(2)
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("from process 1: finished")
		wg3.Done()
	}()

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("from process 2: finished")
		wg3.Done()
	}()

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("from process 3: finished")
		wg3.Done()
	}()
	wg3.Wait()
	time.Sleep(time.Second)

	eg, _ := errgroup.WithContext(context.Background())
	eg.Go(func() error {
		return nil
	})
	eg.Go(func() error {
		return nil
	})
	eg.Go(func() error {
		return nil
	})
	eg.Wait()

	fmt.Println("exit program", time.Since(tn))
}

func NewOrder() A {
	return Order{}
}

func PrintNumbers() {
	for i := 0; i < 1000; i++ {
		fmt.Println(i)
	}
}
