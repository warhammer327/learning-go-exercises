package main

import "fmt"

func two() {
	numbers := []int{}
	for i := 0; i <= 100; i++ {
		numbers = append(numbers, i)
	}
	for _, i := range numbers {
		if i%2 == 0 && i%3 == 0 {
			fmt.Println("Six")
			continue
		}
		if i%2 == 0 {
			fmt.Println("two")
			continue
		}
		if i%3 == 0 {
			fmt.Println("three")
			continue
		}
		fmt.Println("NVM")
	}
}
