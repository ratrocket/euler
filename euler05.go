// Smallest multiple
// Problem 5
// 2520 is the smallest number that can be divided by each of the
// numbers from 1 to 10 without any remainder.
// 
// What is the smallest positive number that is evenly divisible by all
// of the numbers from 1 to 20?
package main

import "fmt"

/*
A number divisible by 20 is already divisible by
1, 2, 4, 5, 10
So only look at multiples of 20 & don't test for divisibility by those
factors.
The remaining factors:
3: only every third multiple of 20 is divisble by 3 (60, 120, ...) so
   check multiples of 60.
6: hmmm, same advice as for 3 (w/o thinking too much)
7: multiples of 420? (least common for 60 & 7)
8: multiples of 840? (this might be going off the rails)
*/
func main() {
	fmt.Println(try1())
}

func try1() int {
	increment := 420
	// leave out 15 because it divides 420
	for i := increment; ; i += increment {
		if i%8 == 0 &&
			i%9 == 0 &&
			i%11 == 0 &&
			i%12 == 0 &&
			i%13 == 0 &&
			i%14 == 0 &&
			i%16 == 0 &&
			i%17 == 0 &&
			i%18 == 0 &&
			i%19 == 0 {
			return i
		}
	}
	return 0
}
// 232792560
