// Sum of even Fibonacci numbers whose value doesn't exceed 4 million.
// 4613732
package main

import "fmt"

func main() {
	var prev, fib uint64 = 0, 1
	var hi uint64 = 4000000
	var tmp uint64
	var sum uint64

	for fib < hi {
		if fib & 1 != 0 {
			sum += fib
		}

		tmp  = prev
		prev = fib
		fib  = tmp + fib
	}
	fmt.Println(sum)
}

/*
// (This is from the gotour)
// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	curr, prev := 0, 0

	return func() int {
		fib := curr
		if curr == 0 && prev == 0 {
			curr = 1
		} else {
			curr = curr + prev
			prev = fib
		}
		return fib
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 15; i++ {
		fmt.Println(f())
	}
}
*/
