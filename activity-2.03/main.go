/*
In this activity, we'll sort a given slice of numbers by swapping the values.
This technique of sorting is known as the bubble sort technique. Go has built-in
 sorting algorithms in the sort package but I don't want you to use them; we
 want you to use the logic and loops you've just learned.

Steps:

Define a slice with unsorted numbers in it.
Print this slice to the console.
Sort the values using swapping.
Once done, print the now sorted numbers to the console.
Tips:

You can do an in-place swap in Go using:
nums[i], nums[i-1] = nums[i-1], nums[i]
You can create a new slice using:
var nums2 []int
You can add to the end of a slice using:
nums2 = append(nums2, 1)
The following is the expected output:

Before: [5, 8, 2, 4, 0, 1, 3, 7, 9, 6]
After : [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]

*/

package main

import (
	"fmt"
)

func main() {

	input := []int{5, 8, 2, 4, 0, 1, 3, 7, 9, 6}
	fmt.Println(input)
	count := 1
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input)-count; j++ {
			if j < len(input)-count {
				if input[j] > input[j+1] {
					input[j], input[j+1] = input[j+1], input[j]
				}
			}
		}
		count++
	}
	fmt.Println(input)
}
