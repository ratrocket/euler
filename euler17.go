// Number letter counts
// Problem 17
// 
// If the numbers 1 to 5 are written out in words: one, two, three,
// four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in
// total.
// 
// If all the numbers from 1 to 1000 (one thousand) inclusive were
// written out in words, how many letters would be used?
// 
// NOTE: Do not count spaces or hyphens. For example, 342 (three hundred
// and forty-two) contains 23 letters and 115 (one hundred and fifteen)
// contains 20 letters. The use of "and" when writing out numbers is in
// compliance with British usage.
//
// A: 21124
package main

import "fmt"

func main() {
	sub20 := []int{
		// "zero"
		0,
		// "one"
		3,
		// "two"
		3,
		// "three"
		5,
		// "four"
		4,
		// "five"
		4,
		// "six"
		3,
		// "seven"
		5,
		// "eight"
		5,
		// "nine"
		4,
		// "ten"
		3,
		// "eleven"
		6,
		// "twelve"
		6,
		// "thirteen"
		8,
		// "fourteen"
		8,
		// "fifteen"
		7,
		// "sixteen"
		7,
		// "seventeen"
		9,
		// "eighteen"
		8,
		// "nineteen"
		8,
	}
	tensCol := []int{
		// null
		0,
		// null
		0,
		// "twenty"; index 2
		6,
		// "thirty"; index 3
		6,
		// "forty"; index 4
		5,
		// "fifty"; index 5
		5,
		// "sixty"; index 6
		5,
		// "seventy"; index 7
		7,
		// "eighty"; index 8
		6,
		// "ninety"; index 9
		6,
	}

	// "thousand" 8
	// "hundred" 7
	// "and" 3
	thousandLetters := 8
	hundresLetters := 7
	andLetters := 3

	var lsum, tsum int // local, total
	var thousands, hundreds, tens, rem int

	from := 1
	to := 1000

	for i := from; i <= to; i++ {
		lsum = 0

		thousands = i / 1000
		rem = i % 1000

		hundreds = rem / 100
		tens = rem % 100

		if tens < 20 {
			lsum += sub20[tens]
		} else {
			lsum += tensCol[tens/10] + sub20[tens%10]
		}

		if hundreds > 0 {
			if lsum > 0 {
				lsum += andLetters
			}

			//  eg, "one" (3)       + "hundred" (7)
			lsum += sub20[hundreds] + hundresLetters
		}

		if thousands > 0 {
			lsum += sub20[thousands] + thousandLetters
		}
		fmt.Printf("Sum for %d: %d\n", i, lsum)
		tsum += lsum
	}
	fmt.Printf("Sum %d to %d: %d\n", from, to, tsum)
}
