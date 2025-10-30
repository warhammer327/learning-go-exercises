package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func MakePerson(firstName, lastName string, age int) person {
	//newPerson := person{firstName: firstName, lastName: lastName, age: age}
	return person{firstName: firstName, lastName: lastName, age: age}

}

func MakePersonPointer(firstName, lastName string, age int) *person {
	//newPersonPointer := person{firstName: firstName, lastName: lastName, age: age}
	return &person{firstName: firstName, lastName: lastName, age: age}
}

func one() {
	p := MakePerson("Fred", "Williamson", 25)
	fmt.Println(p)
	p2 := MakePersonPointer("Wilma", "Fredson", 32)
	fmt.Println(p2)
}
