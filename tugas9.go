package main

import (
	"fmt"
	"strconv"
)

func main() {
	defer errRecover()
	inputUmur()
}

func inputUmur() {
	var input string
	fmt.Printf("Masukkan Umur Anda: ")
	fmt.Scanln(&input)

	var umur int
	var err error
	umur, err = strconv.Atoi(input)

	// Kalo di input angka maka tidak error
	if err == nil {
		fmt.Println("Umur Anda", umur)
		// error kalo di input huruf hasil balik dari input adalah integer
	} else {
		panic(err.Error())
	}
}

func errRecover() {
	if er := recover(); er != nil {
		fmt.Println("Silahkan Masukkan Umur")
	} else {
		fmt.Println("Terima Kasih")
	}
}
