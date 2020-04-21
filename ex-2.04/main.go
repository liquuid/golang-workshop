/*
The Initial if Statement

It's common to need to call a function but not care too much about the returned
value. Often, you'll want to check that it executed correctly and then discard
the returned value. For example, sending an email, writing to a file, or
inserting data into a database; most of the time, if these types of operations
execute successfully, you don't need to worry about the variables they return.

Unfortunately, the variables don't go anywhere as they are still in scope.

To stop these unwanted variables from hanging around, we can use what we know
about scope rules to get rid of them. The best way to check for errors is to
use "initial" statements on if statements. The notation looks like this: if
<initial statement>; <boolean expression> { <code block> }. The initial
statement is in the same section as the Boolean expression, with ; to divide
them.
*/
package main

import (
	"errors"
	"fmt"
)

func validate(input int) error {
	if input < 0 {
		return errors.New("input can't be a negative number")
	} else if input > 100 {
		return errors.New("input can't be over 100")
	} else if input%7 == 0 {
		return errors.New("input can't be divisible by 7")
	} else {
		return nil
	}
}

func main() {
	input := 21
	if err := validate(input); err != nil {
		fmt.Println(err)
	} else if input%2 == 0 {
		fmt.Println(input, "is even")
	} else {
		fmt.Println(input, "is odd")
	}
}
