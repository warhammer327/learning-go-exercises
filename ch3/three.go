package main

import "fmt"

type employee struct {
	firstName string
	lastName  string
	id        int
}

func three() {
	firstEmployee := employee{"firstEmployeeFirstName", "firstEmployeeLastName", 1}
	fmt.Println(firstEmployee)

	secondEmployee := employee{firstName: "secondEmployeeFirstName", lastName: "secondEmployeeLastName", id: 2}
	fmt.Println(secondEmployee)

	var thirdEmployee employee
	thirdEmployee.firstName = "thirdEmployeeFirstName"
	thirdEmployee.lastName = "thirdEmployeeLastName"
	thirdEmployee.id = 3
	fmt.Println(thirdEmployee)
}
