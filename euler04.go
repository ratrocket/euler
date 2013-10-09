// Palindromic numbers (eg, 9009)
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

func main() {
}

func try1() int {
	return 0
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
