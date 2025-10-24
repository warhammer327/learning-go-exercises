package main

import "fmt"

func three() {
	var total int
	// total := total + 1 created new var each time
	// toal = total + 1 reuses var
	for i := 0; i < 10; i++ {
		total = total + 1
		fmt.Println(total)
	}
	fmt.Println("Out side the loop", total)
}
