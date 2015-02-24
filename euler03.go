// Largest prime factor
// Problem 3
//
// The prime factors of 13195 are 5, 7, 13 and 29.
// 
// What is the largest prime factor of the number 600851475143 ?
// (that's over 600 billion)
//
// cf.:
// http://en.wikipedia.org/wiki/Fundamental_theorem_of_arithmetic
//
// First approach: much brute force here.
package main

import(
	"fmt"
	"math"
)

func prime(n int64) bool {
	if n == 1 { return false }
	// some ad hoc "optimization"
	if n > 2 && n&1 == 0 { return false }
	if n > 3 && n%3 == 0 { return false }
	if n > 5 && n%5 == 0 { return false }
	if n > 7 && n%7 == 0 { return false }
	if n > 11 && n%11 == 0 { return false }
	if n > 13 && n%13 == 0 { return false }
	if n > 17 && n%17 == 0 { return false }

	var top int64 = int64(math.Ceil(math.Sqrt(float64(n))))
	for i := int64(2); i <= top; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func main() {
	//var n int64 = 13195
	var n int64 = 600851475143
	var top int64 = int64(math.Ceil(math.Sqrt(float64(n))))

	for i := int64(1); i <= top; i++ {
		if (n % i == 0 && prime(i)) {
			fmt.Println(i)
		}
	}
}

// for checking things out
func printPrimes() {
	for i := 1; i < 100; i++ {
		if prime(int64(i)) {
			fmt.Println(i)
		}
	}
}

/*
Another approach (actually factor the number...)

n = number we want PFs of
primes = empty list
top = ceil(sqrt(n))
i = 2
while i <= top
  if i divides n && i is prime
    push i onto primes
    n = n / i
    top = ceil(sqrt(n))
    i = 2
  else
    i++ (or skip evens)
  end  (is this ruby??)
end
print primes

BUT does this algorithm finish??
(I think yes -- you check to see if the updated n is prime, then done)
*/
