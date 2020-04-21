/* Implementing FizzBuzz

When interviewing for a programming job, you'll be asked to do some coding
exercises. These questions have you writing something from scratch and will have
 several rules to follow. To give you an idea of what that looks like, we'll run
  you through a classic one, "FizzBuzz."

The rules are as follows:

Write a program that prints out the numbers from 1 to 100.
If the number is a multiple of 3, print "Fizz."
If the number is a multiple of 5, print "Buzz."
If the number is a multiple of 3 and 5, print "FizzBuzz."
Here are some tips:

You can convert a number to a string using strconv.Itoa().
The first number to evaluate must be 1 and the last number to evaluate must be
100.
These steps will help you to complete the activity:

Create a loop that does 100 iterations.
Have a variable that keeps count of the number of loops so far.
In the loop, use that count and check whether it's divisible by 3 or 5 using %.
Think carefully about how you'll deal with the "FizzBuzz" case.
package main
*/

package main

import (
	"fmt"
	"strconv"
)

func multiple3(number int) bool {
	if number%3 == 0 {
		return (true)
	}
	return (false)
}
func multiple5(number int) bool {
	if number%5 == 0 {
		return (true)
	}
	return (false)
}
func evaluateFizzBuzz(number int) string {
	if multiple3(number) && multiple5(number) {
		return ("FizzBuzz")
	}
	if multiple3(number) {
		return ("Fizz")
	}
	if multiple5(number) {
		return ("Buzz")
	}
	return (strconv.Itoa(number))
}

func main() {

	for i := 1; i <= 100; i++ {
		fmt.Println(evaluateFizzBuzz(i))

	}
}
