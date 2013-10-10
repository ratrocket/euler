// Sum square difference
// Problem 6
// The sum of the squares of the first ten natural numbers is,
//
// 1^2 + 2^2 + ... + 10^2 = 385
// The square of the sum of the first ten natural numbers is,
//
// (1 + 2 + ... + 10)^2 = 55^2 = 3025
// Hence the difference between the sum of the squares of the first ten
// natural numbers and the square of the sum is 3025 − 385 = 2640.
//
// Find the difference between the sum of the squares of the first one
// hundred natural numbers and the square of the sum.
package main

import "fmt"

func main() {
	fmt.Println(sqOfSum(5))
	fmt.Println(sqOfSum(6))
	fmt.Println(sqOfSum(10))
	fmt.Println(sqOfSum(100))
}

func sumOfSqs(upto int) int {
	return 0
}

func sqOfSum(upto int) int {
	sum := (upto + 1) * (upto / 2)
	if upto&1 == 1 {
		sum += (upto + 1) / 2
	}
	return sum * sum
}

func try1() int {
	return 0
}
