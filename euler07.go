// 10001st prime
// Problem 7
//
// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we
// can see that the 6th prime is 13.
//
// What is the 10 001st prime number?
//
// A: 104743
package main

import (
	"fmt"
	"math"
)

func main() {
	n := 10001
	primes := make([]int, 1)
	for i := 1; len(primes) < 10001+1; i++ {
		if prime(i) {
			primes = append(primes, i)
		}
	}
	fmt.Printf("prime %d: %d\n", n, primes[n])
}

func try1() int {
	return 0
}

func prime(n int) bool {
	if n == 1 { return false }
	if n == 2 { return true }
	// some ad hoc "optimization"
	if n > 2 && n&1 == 0 { return false }
	if n > 3 && n%3 == 0 { return false }
	if n > 5 && n%5 == 0 { return false }
	if n > 7 && n%7 == 0 { return false }
	if n > 11 && n%11 == 0 { return false }
	if n > 13 && n%13 == 0 { return false }
	if n > 17 && n%17 == 0 { return false }

	var top int = int(math.Ceil(math.Sqrt(float64(n))))
	for i := int(2); i <= top; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}
