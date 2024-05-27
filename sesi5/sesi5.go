package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

type User struct {
	ID       int
	Username string
}

var users = []*User{
	{ID: 1, Username: "user1"},
}

func main() {
	// standard package untuk membuat
	// web server
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println(r)
		}
	}()
	// handler
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		// if r.Method == http.MethodGet {
		// 	// get users
		// 	// mengeluarkan semua data users

		// 	// set header
		// 	w.Header().Set("Content-Type", "application/json")
		// 	// return body
		// 	err := json.NewEncoder(w).Encode(users)
		// 	if err != nil {
		// 		http.Error(w, err.Error(), http.StatusInternalServerError)
		// 		return
		// 	}
		// }

		// if r.Method == http.MethodGet {
		// 	// get users
		// 	// mengeluarkan semua data users
		// 	tpl, err := template.ParseFiles("./template.html")
		// 	if err != nil {
		// 		http.Error(w, err.Error(), http.StatusInternalServerError)
		// 		return
		// 	}
		// 	err = tpl.Execute(w, users)
		// 	if err != nil {
		// 		http.Error(w, err.Error(), http.StatusInternalServerError)
		// 		return
		// 	}
		// }

		if r.Method == http.MethodGet {
			// get users
			// mengeluarkan semua data users
			tpl, err := template.ParseFiles("./template_list.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = tpl.Execute(w, users)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		if r.Method == http.MethodPost {
			// akan menambahkan data dari users
			username := r.FormValue("username")
			if username == "" {
				http.Error(w, "username must not empty", http.StatusBadRequest)
				return
			}
			user := &User{Username: username, ID: len(users) + 1}
			users = append(users, user)
			err := json.NewEncoder(w).Encode(user)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	})

	// run server
	err := http.ListenAndServe(":5001", nil)
	if err != nil {
		panic(err)
	}
}

type OrderHandler struct{}

func (o *OrderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	return
}

func deferExit() {
	// defer fmt.Println("from defer function")
	defer fmt.Println("very end of program")
	printName("calman")
	fmt.Println("hello world")

	errSample()

	// memberhentikan aplikasi / program
	// dan menghapus semua buffer dan stack
	// (goroutine, defer stack, dll)
	// sehingga defer yang sudah masuk stack, tidak akan
	// terpanggil
	os.Exit(1)
}

func errSample() {
	// apa itu error?
	// adanya behavior yang tidak terduga
	// salah input / output tidak benar /
	// hilang koneksi
	// . . .

	// recover from defer function
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("error happened", r)
			return
		}
		fmt.Println("program exit gracefully")
	}()
	fmt.Println("PROCESS A")
	// strconv -> package string converter
	// Atoi -> string to integer
	out, err := strconv.Atoi("1000")
	if err != nil {
		fmt.Println("something went wrong", err)
		return
	}
	fmt.Println("PROCESS B")
	// err.Error()
	// UNIT TEST
	fmt.Println(out)
}

func printName(name string) {
	// block printName
	defer fmt.Println("this is from print name1")

	if name == "calman" {
		return
	}

	defer fmt.Println("this is from print name2")
}
