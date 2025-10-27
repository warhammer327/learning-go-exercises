package main

import "fmt"

func main() {
	x := 10
	pointeToX := &x
	fmt.Println(pointeToX)
	fmt.Println(*pointeToX)
	z := *pointeToX + 5
	fmt.Println(z)
	var y *int
	fmt.Println(*y)
}
