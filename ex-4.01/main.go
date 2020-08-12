package main

import "fmt"

func main() {
	fmt.Printf("%#v\n", defineArray())
}
func defineArray() [10]int {
	var arr [10]int
	return arr
}
