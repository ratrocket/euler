// Longest Collatz sequence
// Problem 14
//
// The following iterative sequence is defined for the set of positive
// integers:
// n → n/2    (n is even)
// n → 3n + 1 (n is odd)
//
// Using the rule above and starting with 13, we generate the following
// sequence:
//
// 13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
//
// It can be seen that this sequence (starting at 13 and finishing at 1)
// contains 10 terms. Although it has not been proved yet (Collatz
// Problem), it is thought that all starting numbers finish at 1.
// 
// Which starting number, under one million, produces the longest chain?
// 
// NOTE: Once the chain starts the terms are allowed to go above one
// million.
//
// A:
// Longest chain: 525
// Starting term: 837799
package main

import "fmt"

/*
struct chain {
	start int
	terms []int
	count int
}
*/

var nnext int

func next(n int) int {
	nnext++
	if n&1 == 0 { // even
		return n / 2
	} else {
		return 3 * n + 1
	}
}

func main() {
	var i, curr, count, iter, max, maxi int
	iter = 1000000
	counts := make([]int, iter)

	// To see longest:
	//for i = 837798; i < 837799; i++ {
	for i = 0; i < iter; i++ {
		count = 1  // start at 1 to account for the first value
		curr = i + 1
		fmt.Println("start", curr)
		for {
			curr = next(curr)
			count++
			fmt.Println("     ", curr)
			if curr == 1 {
				// counts[i] is count for i+1
				counts[i] = count
				break
			}
		}

	}
	//fmt.Println(counts)

	for i := range counts {
		if counts[i] > max {
			max = counts[i]
			maxi = i + 1 // account for offset of term/index
		}
	}
	fmt.Println("Longest chain:", max)
	fmt.Println("Starting term:", maxi)
	fmt.Println("next() called:", nnext)
}
