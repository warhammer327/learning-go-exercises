package main

import "fmt"

func two() {
	fmt.Println("Two starts")
	const value = 5
	var i int = value
	var f float32 = value
	fmt.Println(i)
	fmt.Println(f)
}
