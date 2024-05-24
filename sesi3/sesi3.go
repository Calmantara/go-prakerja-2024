package main

import (
	"fmt"
	"time"

	"github.com/Calmantara/go-prakerja-2024/sesi3/order"
	"github.com/Calmantara/go-prakerja-2024/sesi3/product"
)

type SomeStruct struct {
	fn func() int
}

type User struct {
	ID          uint64
	Username    string
	Email       string
	MaskedEmail string
	JoinDate    time.Time
	Address     Address
	Metadata    struct {
		Eduation string
	}
}

type Address struct {
	City    string
	ZipCode int
}

type Server struct {
	IP   string
	Port int
}

// Ada tambahan type user
// student, adalah salah satu kategori dari user
// inheritance
type Student struct {
	// embedded struct
	User
}

type Teacher struct {
	User
}

func main() {
	// init order config

	// I function
	// func Name(in type)(type){}
	age, yearNow, err := CalculateAge(2000)
	fmt.Println(age, yearNow, err)

	age, yearNow, err = CalculateAge(2025)
	fmt.Println(age, yearNow, err)

	fmt.Println(Greet("calman"))

	// hitung rata rata dari suatu inputan
	// 1. kita tidak tau berapa banyak inputannya
	// 2. output kita tau adalah float

	in := []float64{1, 1, 1}
	fmt.Println("average:", Average(in))
	// ... menandakan average2 akan menerima
	// semua element yang ada di slice in
	fmt.Println("average:", Average2(in...))
	fmt.Println("average:", Average2(1, 2, 3, 4, 5, 6, 7, 8, 9))

	// apakah bisa function returning function?
	// bisa pada fn2

	// apakah function bisa kita simpan dalam variable?
	process := fn1
	if age > 10 {
		process = fn3
	}
	process()

	// anonymous function
	fn := func(in int) float64 {
		return float64(in)
	}
	fmt.Println(fn(1))
	// NB: https://golang.cafe/blog/golang-functional-options-pattern.html

	// II Struct
	addressUser1 := Address{
		City:    "Jakarta",
		ZipCode: 123456,
	}
	user1 := User{
		ID:       1,
		Username: "user1",
		Email:    "user1@mail.com",
		JoinDate: time.Now(),
		Address:  addressUser1,
		Metadata: struct{ Eduation string }{
			Eduation: "High School",
		},
	}
	user2 := User{
		ID:       2,
		Username: "user2",
		Email:    "user2@mail.com",
		JoinDate: time.Now(),
		Address: Address{
			City:    "Surabaya",
			ZipCode: 654321,
		},
	}
	user2.Metadata.Eduation = "Junior High School"

	fmt.Printf("user1: %+v \nuser2: %+v \n", user1, user2)

	// embedded struct
	student1 := Student{}
	student1.Username = "student1"
	student1.User.Email = "student1@mail.com"
	fmt.Printf("student1: %+v", student1)

	// struct bisa juga kita gunakan
	// dalam slice / array
	users := []User{
		{ID: 1, Username: "calman", JoinDate: time.Now()},
		{ID: 2, Username: "tara", JoinDate: time.Now()},
	}
	users = append(users, User{
		ID:       3,
		Username: "calmantara",
		JoinDate: time.Now(),
	})
	fmt.Println(users)
	// anonymous
	tempStruct := struct {
		Name string
	}{}
	tempStruct.Name = "calman"

	// num1 sudah ada address dan sudah ada value
	num1 := 1
	var num2 *int
	num2 = &num1
	*num2 = 2
	fmt.Println(num1)

	// userPtr belum ada di ram
	var userPtr *User
	// usr sudah ada di ram
	var usr User
	_ = usr
	userPtr = &User{
		ID:       1,
		Username: "lalala",
	}
	// AWS SDK
	fmt.Println(userPtr)

	srv := Server{}
	srv.IP, srv.Port = DefaultServer()

	// METHOD
	// fungsi yang nempel di struct (METHOD)
	Greeting(user1)
	user1.Greeting()
	user1.MaskEmail()

	fmt.Println(user1)

	student1.Greeting()

	// create order
	order1 := order.Order{
		ID:       1,
		Price:    1000.0,
		Discount: 0.1,
	}
	fmt.Println(order1.FinalPrice(), order.ORDER_NAME)
	fmt.Println(product.ProductName)
}

func (u User) Greeting() {
	fmt.Printf("Hello %v with email %v", u.Username, u.Email)
}

func (u *User) MaskEmail() {
	u.MaskedEmail = u.Email[:3] + "****"
}

func Greeting(user User) {
	fmt.Printf("Hello %v with email %v", user.Username, user.Email)
}

func DefaultServer() (ip string, port int) {
	return "127.0.0.1", 8080
}

func fn1() int {
	return 0
}

func fn2() func() int {
	return fn1
}

func fn3() int {
	return 1
}

func Average(in []float64) float64 {
	sum := 0.0
	for _, i := range in {
		sum += i
	}
	return sum / float64(len(in))
}

func Average2(in ...float64) float64 {
	sum := 0.0
	for _, i := range in {
		sum += i
	}
	return sum / float64(len(in))
}

func CalculateAge(year int) (int, int, error) {
	// CalculateAge => function name
	// year => input parameter
	// int di input => type input parameter
	// int di paling kanan => output type

	// NB: tidak masalah nama variable dari output
	// yang penting type sama

	yearNow := time.Now().Year()
	if yearNow < year {
		return 0, yearNow, fmt.Errorf("invalid year input, year must less than %v", yearNow)
	}
	return yearNow - year, yearNow, nil
}

func Greet(name string) (greeting string) {
	greeting = fmt.Sprintf("Hello, %v!", name)
	return greeting
}

func Add(in1, in2 int) (out int) {
	out = in1 + in2
	return out
}

func Div(in1 int, in2 int) int {
	return in1 / in2
}
