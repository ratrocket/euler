// Highly divisible triangular number
// Problem 12
//
// The sequence of triangle numbers is generated by adding the natural
// numbers. So the 7th triangle number would be 1 + 2 + 3 + 4 + 5 + 6 +
// 7 = 28. The first ten terms would be:
//
// 1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...
//
// Let us list the factors of the first seven triangle numbers:
//  1:  1
//  3:  1,3
//  6:  1,2,3,6
//  10: 1,2,5,10
//  15: 1,3,5,15
//  21: 1,3,7,21
//  28: 1,2,4,7,14,28
//
// We can see that 28 is the first triangle number to have over five
// divisors.
//
// What is the value of the first triangle number to have over five
// hundred divisors?
//
// A: 76576500 (12375 th triangle number)
package main

import (
	"fmt"
	"sort"
)

func factors(n int) []int {
	fac := make([]int, 0)
	var lo, hi, last_hi int // hi is lo's complement (lo*hi==n)
	last_hi = -1

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			lo = i
			hi = n / i

			if lo == hi { // square
				fac = append(fac, lo)
				break
			} else if lo == last_hi { // crossed over
				break
			}

			fac = append(fac, lo, hi)
			last_hi = hi
		}
	}

	sort.Ints(fac)
	return fac
}

func main() {
	var triangle, l int
	nseen := 0
	for i := 1; ; i++ {
		triangle += i
		l = len(factors(triangle))
		//fmt.Println(i, triangle, l)
		//fmt.Println(factors(triangle))
		if l > 500 {
			nseen++
			fmt.Println(i, triangle, l)
			//fmt.Println("Here ^^^^^")
			if nseen > 4 {
				break
			}
		}
	}
}


/*
// For looking at properties of the number of factors of the triangle
// numbers.
func main() {
	var triangle, l int
	for i := 1; i < 12500; i++ {
		triangle += i
		l = len(factors(triangle))
		fmt.Println(i, triangle, l)
	}
}
*/
