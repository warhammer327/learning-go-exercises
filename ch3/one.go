package main

import "fmt"

var x = []string{"Hello", "Hola", "नमस्कार", "こんにち", "Привіт"}

func one() {
	first_slice := x[:2]
	fmt.Println(first_slice)
	second_slice := x[2:5]
	fmt.Println(second_slice)
	third_slice := x[3:5]
	fmt.Println(third_slice)
}
