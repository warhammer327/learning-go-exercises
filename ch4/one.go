package main

import "fmt"

func one() {
	number := []int{}
	for i := 0; i <= 100; i++ {
		number = append(number, i)
	}
	fmt.Println(number)
}
