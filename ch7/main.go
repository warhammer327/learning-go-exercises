package main

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) ShowInformation() {
	fmt.Printf("%s %s %d\n", p.FirstName, p.LastName, p.Age)
}

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, update: %v", c.total, c.lastUpdated)
}

var c Counter

func main() {
	//p1 := Person{FirstName: "abc", LastName: "xyz", Age: 1}
	//p1.ShowInformation()

	//fmt.Println(c.String())
	//c.Increment()
	//fmt.Println(c.String())

	//fmt.Println("===running intrfce.go===")
	//intrfce()

	fmt.Println("===running exercise===")
	one()
}
