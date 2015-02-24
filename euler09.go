// A Pythagorean triplet is a set of three natural numbers, a < b < c,
// for which,
// 
// a^2 + b^2 = c^2
// For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.
// 
// There exists exactly one Pythagorean triplet for which
// a + b + c = 1000.
//
// Find the product abc.
//
// A:
// a,b,c 200 375 425
// a+b+c 1000
// a*b*c 31875000

package main

import "fmt"

// Clearly this is a dumb way to find perfect squares!
//
// if c ("candidate") is a square, returns sqrt(c), true
// else returns 0, false
func issq(c int) (int, bool) {
	var p int
	for i := 0; ; i++ {
		p = i * i
		if p == c {
			return i, true
		} else if p > c {
			return 0, false
		}
	}
}

func main() {
	var a, b, c, c2 int
	var succ bool

	for a = 1; a < 1000; a++ {
		for b = a + 1; b < 1000; b++ {
			if a+b > 999 { // no chance
				continue
			}

			c2 = a*a + b*b
			c, succ = issq(c2)

			if succ && a+b+c == 1000 {
				fmt.Println("a,b,c", a, b, c)
				fmt.Println("a+b+c", a+b+c)
				fmt.Println("a*b*c", a*b*c)
				return
			}
		}
	}
}

/*
if x = a^2 + b^2; issquare(x); c = sqrt(x)
	if a + b + c == 1000
		print a*b*c
*/
