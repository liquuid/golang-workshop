package main

import "fmt"

func main() {
	username := "Sir_King_Über"
	runes := []rune(username)
	fmt.Println("Bytes:", len(username))
	fmt.Println("Runes:", len([]rune(username)))
	for i := 0; i < len(runes); i++ {
		fmt.Print(string(runes[i]), " ")
	}
}
