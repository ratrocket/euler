// Largest palindrome product
// Problem 4
//
// A palindromic number reads the same both ways. The largest palindrome
// made from the product of two 2-digit numbers is 9009 = 91 × 99.
//
// Find the largest palindrome made from the product of two 3-digit
// numbers.
//
// (without using strings!)
package main

import "fmt"

// think: reverse recursively
// think: use a stack (explicitly)
func rev(n int) int {
	var r int

	for n > 0 {
		r = r * 10 + n % 10
		n /= 10
	}
	return r
}

func pal(n int) bool {
	return rev(n) == n
}

// Is there a way to write this with a function that returns a function
// (a closure)?  Just for fun, of course.
func main() {
	fmt.Println(try1())
}

// two digit numbers (so answer is 9009, see above)
func try1() int {
	var lg, prod int
	for i := 10; i < 100; i++ { // simplest, slowest way; SLOW
		for j := 10; j < 100; j++ {
			prod = i * j
			if pal(prod) && prod > lg {
				lg = prod
			}
		}
	}
	return lg
}

func test() { // ad hoc test... better than nothing?
	if rev(1234) != 4321     { fmt.Println("error") }
	if rev(9009) != 9009     { fmt.Println("error") }
	if rev(100244) != 442001 { fmt.Println("error") }
	if rev(12) != 21         { fmt.Println("error") }
	if rev(1) != 1           { fmt.Println("error") }
	if rev(10) != 1          { fmt.Println("error") }
	/*
	fmt.Println(rev(1234))
	fmt.Println(rev(9009))
	fmt.Println(rev(100244))
	fmt.Println(rev(12))
	fmt.Println(rev(1))
	fmt.Println(rev(10))
	*/
}
