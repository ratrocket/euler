// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
//
// Find the sum of all the primes below two million.
// A: 142913828922
package main

import (
	"fmt"
	"math"
)

// Move this into its own package.
func isprime(n int) bool {
	if n == 2 { return true }
	if n == 1 { return false }
	// some ad hoc "optimization"
	if n > 2 && n&1 == 0 { return false }

	if n > 3 && n%3 == 0 { return false }
	if n > 5 && n%5 == 0 { return false }
	if n > 7 && n%7 == 0 { return false }
	if n > 11 && n%11 == 0 { return false }
	if n > 13 && n%13 == 0 { return false }
	if n > 17 && n%17 == 0 { return false }
	if n > 19 && n%19 == 0 { return false }
	if n > 23 && n%23 == 0 { return false }
	if n > 29 && n%29 == 0 { return false }
	if n > 31 && n%31 == 0 { return false }
	if n > 37 && n%37 == 0 { return false }
	if n > 41 && n%41 == 0 { return false }
	if n > 43 && n%43 == 0 { return false }
	if n > 47 && n%47 == 0 { return false }
	if n > 53 && n%53 == 0 { return false }
	if n > 59 && n%59 == 0 { return false }
	if n > 61 && n%61 == 0 { return false }
	if n > 67 && n%67 == 0 { return false }
	if n > 71 && n%71 == 0 { return false }
	if n > 73 && n%73 == 0 { return false }
	if n > 79 && n%79 == 0 { return false }
	if n > 83 && n%83 == 0 { return false }
	if n > 89 && n%89 == 0 { return false }
	if n > 97 && n%97 == 0 { return false }
	if n > 101 && n%101 == 0 { return false }
	if n > 103 && n%103 == 0 { return false }
	if n > 107 && n%107 == 0 { return false }
	if n > 109 && n%109 == 0 { return false }
	if n > 113 && n%113 == 0 { return false }
	if n > 127 && n%127 == 0 { return false }
	if n > 131 && n%131 == 0 { return false }
	if n > 137 && n%137 == 0 { return false }
	if n > 139 && n%139 == 0 { return false }
	if n > 149 && n%149 == 0 { return false }
	if n > 151 && n%151 == 0 { return false }
	if n > 157 && n%157 == 0 { return false }
	if n > 163 && n%163 == 0 { return false }
	if n > 167 && n%167 == 0 { return false }
	if n > 173 && n%173 == 0 { return false }
	if n > 179 && n%179 == 0 { return false }
	if n > 181 && n%181 == 0 { return false }
	if n > 191 && n%191 == 0 { return false }
	if n > 193 && n%193 == 0 { return false }
	if n > 197 && n%197 == 0 { return false }
	if n > 199 && n%199 == 0 { return false }
	if n > 211 && n%211 == 0 { return false }
	if n > 223 && n%223 == 0 { return false }
	if n > 227 && n%227 == 0 { return false }
	if n > 229 && n%229 == 0 { return false }
	if n > 233 && n%233 == 0 { return false }
	if n > 239 && n%239 == 0 { return false }
	if n > 241 && n%241 == 0 { return false }
	if n > 251 && n%251 == 0 { return false }
	if n > 257 && n%257 == 0 { return false }
	if n > 263 && n%263 == 0 { return false }
	if n > 269 && n%269 == 0 { return false }
	if n > 271 && n%271 == 0 { return false }
	if n > 277 && n%277 == 0 { return false }
	if n > 281 && n%281 == 0 { return false }
	if n > 283 && n%283 == 0 { return false }
	if n > 293 && n%293 == 0 { return false }
	if n > 307 && n%307 == 0 { return false }
	if n > 311 && n%311 == 0 { return false }
	if n > 313 && n%313 == 0 { return false }
	if n > 317 && n%317 == 0 { return false }
	if n > 331 && n%331 == 0 { return false }
	if n > 337 && n%337 == 0 { return false }
	if n > 347 && n%347 == 0 { return false }
	if n > 349 && n%349 == 0 { return false }
	if n > 353 && n%353 == 0 { return false }
	if n > 359 && n%359 == 0 { return false }
	if n > 367 && n%367 == 0 { return false }
	if n > 373 && n%373 == 0 { return false }
	if n > 379 && n%379 == 0 { return false }
	if n > 383 && n%383 == 0 { return false }
	if n > 389 && n%389 == 0 { return false }
	if n > 397 && n%397 == 0 { return false }
	if n > 401 && n%401 == 0 { return false }
	if n > 409 && n%409 == 0 { return false }
	if n > 419 && n%419 == 0 { return false }
	if n > 421 && n%421 == 0 { return false }
	if n > 431 && n%431 == 0 { return false }
	if n > 433 && n%433 == 0 { return false }
	if n > 439 && n%439 == 0 { return false }
	if n > 443 && n%443 == 0 { return false }
	if n > 449 && n%449 == 0 { return false }
	if n > 457 && n%457 == 0 { return false }
	if n > 461 && n%461 == 0 { return false }
	if n > 463 && n%463 == 0 { return false }
	if n > 467 && n%467 == 0 { return false }
	if n > 479 && n%479 == 0 { return false }
	if n > 487 && n%487 == 0 { return false }
	if n > 491 && n%491 == 0 { return false }
	if n > 499 && n%499 == 0 { return false }
	if n > 503 && n%503 == 0 { return false }
	if n > 509 && n%509 == 0 { return false }
	if n > 521 && n%521 == 0 { return false }
	if n > 523 && n%523 == 0 { return false }
	if n > 541 && n%541 == 0 { return false }
	if n > 547 && n%547 == 0 { return false }
	if n > 557 && n%557 == 0 { return false }
	if n > 563 && n%563 == 0 { return false }
	if n > 569 && n%569 == 0 { return false }
	if n > 571 && n%571 == 0 { return false }
	if n > 577 && n%577 == 0 { return false }
	if n > 587 && n%587 == 0 { return false }
	if n > 593 && n%593 == 0 { return false }
	if n > 599 && n%599 == 0 { return false }
	if n > 601 && n%601 == 0 { return false }
	if n > 607 && n%607 == 0 { return false }
	if n > 613 && n%613 == 0 { return false }
	if n > 617 && n%617 == 0 { return false }
	if n > 619 && n%619 == 0 { return false }
	if n > 631 && n%631 == 0 { return false }
	if n > 641 && n%641 == 0 { return false }
	if n > 643 && n%643 == 0 { return false }
	if n > 647 && n%647 == 0 { return false }
	if n > 653 && n%653 == 0 { return false }
	if n > 659 && n%659 == 0 { return false }
	if n > 661 && n%661 == 0 { return false }
	if n > 673 && n%673 == 0 { return false }
	if n > 677 && n%677 == 0 { return false }
	if n > 683 && n%683 == 0 { return false }
	if n > 691 && n%691 == 0 { return false }
	if n > 701 && n%701 == 0 { return false }
	if n > 709 && n%709 == 0 { return false }
	if n > 719 && n%719 == 0 { return false }
	if n > 727 && n%727 == 0 { return false }
	if n > 733 && n%733 == 0 { return false }
	if n > 739 && n%739 == 0 { return false }
	if n > 743 && n%743 == 0 { return false }
	if n > 751 && n%751 == 0 { return false }
	if n > 757 && n%757 == 0 { return false }
	if n > 761 && n%761 == 0 { return false }
	if n > 769 && n%769 == 0 { return false }
	if n > 773 && n%773 == 0 { return false }
	if n > 787 && n%787 == 0 { return false }
	if n > 797 && n%797 == 0 { return false }
	if n > 809 && n%809 == 0 { return false }
	if n > 811 && n%811 == 0 { return false }
	if n > 821 && n%821 == 0 { return false }
	if n > 823 && n%823 == 0 { return false }
	if n > 827 && n%827 == 0 { return false }
	if n > 829 && n%829 == 0 { return false }
	if n > 839 && n%839 == 0 { return false }
	if n > 853 && n%853 == 0 { return false }
	if n > 857 && n%857 == 0 { return false }
	if n > 859 && n%859 == 0 { return false }
	if n > 863 && n%863 == 0 { return false }
	if n > 877 && n%877 == 0 { return false }
	if n > 881 && n%881 == 0 { return false }
	if n > 883 && n%883 == 0 { return false }
	if n > 887 && n%887 == 0 { return false }
	if n > 907 && n%907 == 0 { return false }
	if n > 911 && n%911 == 0 { return false }
	if n > 919 && n%919 == 0 { return false }
	if n > 929 && n%929 == 0 { return false }
	if n > 937 && n%937 == 0 { return false }
	if n > 941 && n%941 == 0 { return false }
	if n > 947 && n%947 == 0 { return false }
	if n > 953 && n%953 == 0 { return false }
	if n > 967 && n%967 == 0 { return false }
	if n > 971 && n%971 == 0 { return false }
	if n > 977 && n%977 == 0 { return false }
	if n > 983 && n%983 == 0 { return false }
	if n > 991 && n%991 == 0 { return false }
	if n > 997 && n%997 == 0 { return false }

	top := int(math.Ceil(math.Sqrt(float64(n))))

	for i := 2; i <= top; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func main() {
	sum := 2

	for i := 3; i < 2000000; i += 2 {
		if isprime(i) {
			sum += i
		}
	}

	fmt.Println(sum)
}
