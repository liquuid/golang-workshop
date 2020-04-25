package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	// Declare an int and initialize with math.MaxInt64, which is the highest
	// possible value for an int64 in Go, which is defined as a constant:
	intA := math.MaxInt64
	intA = intA + 1
	// Now we'll create a big int. This is a custom type and is not based on Go's
	// int type. We'll also initialize it with Go's highest possible number value:
	bigA := big.NewInt(math.MaxInt64)
	bigA.Add(bigA, big.NewInt(1))
	// Print out the max int size and the values for our Go int and our big int:
	fmt.Println("MaxInt64: ", math.MaxInt64)
	fmt.Println("Int   :", intA)
	fmt.Println("Big Int : ", bigA.String())

}
