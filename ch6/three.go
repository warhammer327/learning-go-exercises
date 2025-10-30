package main

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePersonTest(firstName, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func three() {
	var people []Person
	//people := make([]Person, 0, 10_000_000)
	for i := 0; i < 10_000_000; i++ {
		people = append(people, MakePersonTest("Fred", "Williamson", 25))
	}
}
