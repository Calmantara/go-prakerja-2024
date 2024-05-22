package main

import (
	"fmt"
	"time"
)

func main() {
	// I conditions
	// - if; else; else if
	// - switch
	// - type bool
	age := 20
	if age < 17 {
		fmt.Println("masih bocil")
	} else if age >= 17 && age < 20 {
		// OR => ||
		// AND => &&
		fmt.Println("abg")
	} else {
		// kondisi default
		fmt.Println("sudah dewasa")
	}

	// temporary variable
	name := "tara"
	if name := "calman"; len(name) >= 10 {
		fmt.Println("valid")
	} else {
		if age < 17 {
			fmt.Println("calman bocil")
		}
		// di print
		name = "calmantara"
		fmt.Println(name)
	}
	// di print
	fmt.Println(name)

	// switch
	// kalau age == 10 / 20 / 30 => kamu keren
	// kalau age 17 => anak baru gede
	age = 10
	switch age {
	case 10, 20, 30:
		fmt.Println("kamu keren")
	case 17:
		fmt.Println("anak baru gede")
	default:
		fmt.Println("tidak masuk kriteria")
	}

	switch {
	case age < 17:
		fmt.Println("masih bocil")
		fallthrough
	case age >= 17 && age < 20:
		fmt.Println("abg")
	default:
		fmt.Println("sudah dewasa")
	}

	// II Pengulangan / Looping
	// 1. menampilkan data banyak / mengakses data
	// 2. melakukan sesuatu berulang / menghindari code duplication

	// menampilkan data 1 - 10
	fmt.Println(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	// menampilkan data 1 - 1000000
	// 1. infinite loop
	threshold, counter := 100, 0
	for {
		fmt.Println(counter)
		time.Sleep(time.Second)
		if counter > threshold {
			break
		}
		// increase 10
		counter += 100
	}
	// 2. bounded loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		// break
		if i == 4 {
			break
		}
	}
	arr := []string{"calman", "tara"}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	// 3. accessing data
	for idx, name := range arr {
		fmt.Println(idx, name)
	}
	keyValue := map[string]any{"key1": "val1", "key2": "val2"}
	for key, val := range keyValue {
		fmt.Println(key, val)
	}
	// apakah bisa nested loop
	// loop di dalam loop
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println(i, j)
		}
	}

	// ARRAY vs SLICE
	arrInt := [5]int{1, 2, 3, 4, 5}
	sliceInt := []int{1, 2, 3, 4, 5}
	fmt.Println(arrInt, sliceInt)

	// apakah bisa aku menambah element
	// di dalam arr / slice
	// append untuk menambahkan element slice
	sliceInt = append(sliceInt, 6, 7, 8, 9)
	// arrInt = append(arrInt, 1) => tidak valid
	fmt.Println(arrInt, sliceInt)

	// mengubah nilai di dalam arr / slice
	for i := 0; i < len(arrInt); i++ {
		arrInt[i] = i * 100
	}
	// ketika kita mengakses element
	// dari arr / slice melebihi slot yang tersedia
	// maka program akan PANIC
	fmt.Println(arrInt)

	//
	slc := []int{1, 2, 3, 4}
	slc2 := slc
	// ubah nilai di slice 2
	slc2[0] = 100
	fmt.Println(slc, slc2)
	// copy
	slc3 := []int{0, 0, 0}
	nn := copy(slc3, slc2)
	slc3[0] = 1000
	fmt.Println(nn, slc, slc2, slc3)

	// init new slice
	newSlc := []int{}
	for i := 0; i < 2000; i++ {
		newSlc = append(newSlc, i)
	}
	newSlc2 := make([]int, 1000)
	copy(newSlc2, newSlc)
	// kegunaan slice
	// 1. input data
	// user mau menginput data nilai siswa sebanyak 1000
	// siswa menjadi 1450
	// 2. mengambil data di database

	//bagaimana caranya kita mengakses index 100 - 200
	// mengambil nilai dari 100 - 199
	slc4 := newSlc2[100:200]
	// mengambil nilai dari 200 - selesai
	slc5 := newSlc2[200:]
	// mengambil nilai dari awal sampai 199
	slc6 := newSlc2[:200]
	fmt.Println(len(slc4), len(slc5), len(slc6))

	// menghapus element
	// dengan index 175
	// harus di copy
	rmSlc := newSlc2[:175]
	rmSlc2 := newSlc2[176:]
	rmSlc = append(rmSlc, rmSlc2...)
	fmt.Println(newSlc2[175], rmSlc[175])
}
