/*
Looping Over Map Data Using range

The steps to solve the activity are as follows:

Put the words into a map like this:
  words := map[string]int{
  "Gonna": 3,
  "You": 3,
  "Give": 2,
  "Never": 1,
  "Up":  4,
}
Create a loop and use range to capture the word and the count.
Keep track of the word with the highest count using a variable for what the
highest count is and its associated word.
Print the variables out.
The following is the expected output displaying the most popular word with its
count value:

Most popular word: Up
With a count of : 4

*/

package main

import (
	"fmt"
)

func main() {
	words := map[string]int{
		"Gonna": 3,
		"You":   3,
		"Give":  12,
		"Never": 1,
		"Up":    4,
	}
	highest := ""
	for key := range words {
		if words[key] > words[highest] {
			highest = key
		}
	}
	fmt.Println("Most popular word:", highest)
	fmt.Println("With a count of :", words[highest])
}
