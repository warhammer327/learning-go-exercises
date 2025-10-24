package main

import "fmt"

var value string = "Hi 😊 and 👋 "

func two() {
	var first_icon string = value[3:7]
	fmt.Println(first_icon)
	var second_icon string = value[12:17]
	fmt.Println(second_icon)
}
