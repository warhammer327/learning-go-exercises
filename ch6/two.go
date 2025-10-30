package main

import "fmt"

func UpdateSlice(list []string, value string) {
	fmt.Println("Update slice starts")
	fmt.Println(list)
	last_pos := len(list) - 1
	list[last_pos] = value
	fmt.Println(list)
	fmt.Println("Update slice ends")
}

func GrowSlice(list []string, value string) {
	fmt.Println("Grow slice starts")
	fmt.Println(list)
	list = append(list, value)
	fmt.Println(list)
	fmt.Println("Grow slice ends")
}

func GrowSliceWithPointer(list *[]string, value string) {
	fmt.Println("Grow slice pointer starts")
	fmt.Println(list)
	*list = append(*list, value)
	fmt.Println(list)
	fmt.Println("Grow slice pointer ends")
}

func two() {
	string_list := []string{"a", "b", "c"}
	UpdateSlice(string_list, "d")
	GrowSlice(string_list, "e")
	GrowSliceWithPointer(&string_list, "e")
	fmt.Println("OG LIST")
	fmt.Println(string_list)
}
