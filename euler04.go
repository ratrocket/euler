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
	/*
	2 digit largest:  9009
	3 digit largest:  906609
	4 digit largest:  99000099
	*/
	fmt.Println("2 digit largest: ", try1(10, 100))
	fmt.Println("3 digit largest: ", try1(100, 1000))
	//fmt.Println("4 digit largest: ", try1(1000, 10000))
}

// TODO how to make this smarter & more efficient?
func try1(lo, hi int) int {
	var lg, prod int
	for i := lo; i < hi; i++ { // simplest, slowest way; SLOW
		for j := lo; j < hi; j++ {
			prod = i * j
			if pal(prod) && prod > lg {
				lg = prod
			}
		}
	}
	return lg
}

// THIS DOES NOT WORK
// go from top to bottom
func try2(lo, hi int) int {
	for i := hi - 1; i >= lo; i-- {
		for j := hi - 1; j >= lo; j-- {
			prod := i * j
			if pal(prod) {
				return prod
			}
		}
	}
	return 0 // didn't find anything
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
