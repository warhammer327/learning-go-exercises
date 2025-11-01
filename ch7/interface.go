package main

import (
	"fmt"
	"log"
	"strconv"
)

type Book struct {
	Title  string
	Author string
}

func (b Book) String() string {
	return fmt.Sprintf("Book: %s - %s", b.Title, b.Author)
}

type Count int

func (c Count) String() string {
	return strconv.Itoa(int(c))
}

func WriteLog(s fmt.Stringer) {
	log.Println(s.String())
}

func intrfce() {
	book := Book{"Alice in Wonderland", "Lewis Carrol"}
	WriteLog(book)
	count := Count(3)
	WriteLog(count)
}
