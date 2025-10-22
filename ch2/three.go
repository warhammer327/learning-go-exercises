package main

import (
	"fmt"
	"math"
)

func three() {
	fmt.Println("starting three")
	var b byte = math.MaxUint8
	var small int32 = math.MaxInt32
	var bigI uint64 = math.MaxUint64

	b = b + 1
	small = small + 1
	bigI = bigI + 1
	fmt.Println(b)
	fmt.Println(small)
	fmt.Println(bigI)
}
