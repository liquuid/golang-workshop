/*
Loan Calculator

In this activity, we create a loan calculator for an online financial advisor
platform. Our calculator has the following rules:

A good credit score is a score of 450 or above.
For a good credit score, your interest rate is 15%.
For a less than good score, your interest rate is 20%.

For a good credit score, your monthly payment must be no more than 20% of your
monthly income.

For a less than good credit score, your monthly payment must be no more than 10%
 of your monthly income.

If a credit score, monthly income, loan amount, or loan term is less than 0,
return an error.

If the term of the loan if not divisible by 12 months, return an error.

The interest payment will be a simple calculation of loan amount * interest
rate * loan term.
After doing these calculations, display the following details to the user:

Applicant X
-----------
Credit Score : X
Income : X
Loan Amount : X
Loan Term : X
Monthly Payment : X
Rate : X
Total Cost : X
Approved : X
*/

package main

import "fmt"

func rate(creditScore int) float32 {
	if creditScore > 450 {
		return 0.15
	}
	return 0.2

}
func isAproved(income float32, mp float32) bool {
	if mp/income < 0.2 {
		return true
	}
	return false

}
func main() {
	type Applicant struct {
		creditScore int
		income      float32
		loanAmount  float32
		loanTerm    float32
	}
	listOfApplicants := []Applicant{}

	listOfApplicants = append(listOfApplicants, Applicant{500, 1000, 1000, 24})
	listOfApplicants = append(listOfApplicants, Applicant{350, 1000, 10000, 12})

	for n, item := range listOfApplicants {
		fmt.Println("Applicant ", n+1)
		fmt.Println("--------- ")
		fmt.Println("Credit Score: ", item.creditScore)
		fmt.Println("Income: ", item.income)
		fmt.Println("Loan Amount: ", item.loanAmount)
		fmt.Println("Loan Term: ", item.loanTerm)
		fmt.Println("Montly Payment: ", (item.loanAmount+item.loanAmount*rate(item.creditScore))/item.loanTerm)
		fmt.Println("Rate: ", rate(item.creditScore)*100)
		fmt.Println("Total Cost: ", item.loanAmount*rate(item.creditScore))
		fmt.Println("Aproved: ", isAproved(item.income, (item.loanAmount+item.loanAmount*rate(item.creditScore))/item.loanTerm))
		fmt.Println()
	}

}
