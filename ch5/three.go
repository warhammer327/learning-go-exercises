package main

import "fmt"

func prefixer(greeting string) func(string) string {
	return func(name string) string {
		ret := greeting + " " + name
		return ret
	}
}

func three() {
	helloPrefixer := prefixer("hello")
	fmt.Println(helloPrefixer("bob"))
	fmt.Println(prefixer("hello")("martha"))
}
