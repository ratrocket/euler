package main

import "fmt"

func main() {
	fmt.Println("try1: ", try1())
	fmt.Println("try2: ", try2())
	fmt.Println("try3: ", try3())
}

/*
Break 1..1000 into chunks with 15 numbers each.

In each chunk you know the multiples of 3 and multiples of 5 are indexed
at (offsets) 2, 4, 5, 8, 9, 11, 14

01 02 03 04 05 06 07 08 09 10 11 12 13 14 15  <= values for chunk one
0  1  2  3  4  5  6  7  8  9  10 11 12 13 14  <= index
      !        !        !        !        !
            *              *
*/
func try3() int {
	a := [15]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	addend := len(a)
	var sum int

	for a[14] < 1000 {
		sum += a[2] + a[4] + a[5] + a[8] + a[9] + a[11] + a[14]
		for i := 0; i < addend; i++ { // a gets next "chunk"
			a[i] += addend
		}
	} // at end of loop, a has 1000 in it (& >1000) so ADD to sum

	for i := 0; a[i] < 1000; i++ {
		switch i {
		case 2, 4, 5, 8, 9, 11, 14:
			sum += a[i]
		}
	}
	return sum
}

// the straightforward solution
func try1() int {
	sum := 0
	for i := 1; i < 1000; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			sum += i
		}
	}
	return sum
}

// alternative (just trying something)
func try2() int {
	sum := 0

	for i := 3; i < 1000; i += 3 {
		sum += i
	}

	for i := 5; i < 1000; i += 5 {
		if i % 3 != 0 {
			sum += i
		}
	}

	return sum
}
