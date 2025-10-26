package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func fileLen(input string) (int, error) {
	f, err := os.Open(input)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	data := make([]byte, 1024)
	total := 0
	for {
		count, err := f.Read(data)
		fmt.Println(string(data[:count]))
		total += count
		if err != nil {
			if err == io.EOF {
				break
			}
			return 1, err

		}
	}
	return total, nil
}

func fileLen_two(input string) (int, error) {
	f, err := os.Open(input)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	total := 0
	for range data {
		total += 1
	}
	return total, nil
}

func two() {
	fmt.Println(fileLen("one.go"))
	fmt.Println(fileLen_two("one.go"))
}
