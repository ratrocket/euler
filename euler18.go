// Maximum path sum I
// Problem 18
// 
// By starting at the top of the triangle below and moving to adjacent
// numbers on the row below, the maximum total from top to bottom is 23.
// 
// 3
// 7 4
// 2 4 6
// 8 5 9 3
// 
// That is, 3 + 7 + 4 + 9 = 23.
// 
// Find the maximum total from top to bottom of the triangle below:
// 
// 75
// 95 64
// 17 47 82
// 18 35 87 10
// 20 04 82 47 65
// 19 01 23 75 03 34
// 88 02 77 73 07 63 67
// 99 65 04 28 06 16 70 92
// 41 41 26 56 83 40 80 70 33
// 41 48 72 33 47 32 37 16 94 29
// 53 71 44 65 25 43 91 52 97 51 14
// 70 11 33 28 77 73 17 78 39 68 17 57
// 91 71 52 38 17 14 91 43 58 50 27 29 48
// 63 66 04 68 89 53 67 30 73 16 69 87 40 31
// 04 62 98 27 23 09 70 98 73 93 38 53 60 04 23
// 
// NOTE: As there are only 16384 routes, it is possible to solve this
// problem by trying every route. However, Problem 67, is the same
// challenge with a triangle containing one-hundred rows; it cannot be
// solved by brute force, and requires a clever method! ;o)
package main

import "fmt"

type node struct {
	value int
	subtreeMax int
}

// if x==y, prefer x
func max(x, y int) int {
	if y > x {
		return y
	}
	return x
}

func leftchild(i, level int) int {
	return i + level + 1
}

func rightchild(i, level int) int {
	return i + level + 2
}

// returns leftchild, rightchild
func children(i, level int) (int, int) {
	return i+level+1, i+level+2
}

func printTree(nodes []node, lsi []int) {
	for i := 0; i < len(lsi) - 1; i++ {
		fmt.Println(nodes[lsi[i]:lsi[i+1]])
	}
}

// Terminology:
// "lsi": (list of) level starter indexes: the index at which a new
// level starts.
func main() {
	//values := []int{
	//	3,
	//	7, 4,
	//	2, 4, 6,
	//	8, 5, 9, 3,
	//}
	//nlevel := 4
	values := []int{
		75,
		95, 64,
		17, 47, 82,
		18, 35, 87, 10,
		20, 4, 82, 47, 65,
		19, 1, 23, 75, 3, 34,
		88, 2, 77, 73, 7, 63, 67,
		99, 65, 4, 28, 6, 16, 70, 92,
		41, 41, 26, 56, 83, 40, 80, 70, 33,
		41, 48, 72, 33, 47, 32, 37, 16, 94, 29,
		53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14,
		70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57,
		91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48,
		63, 66, 4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31,
		4, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 4, 23,
	}
	nlevel := 15
	//maxlevel := nlevel - 1
	//depth := maxlevel

	nodes := make([]node, len(values))

	for i := range values {
		nodes[i] = node{values[i],0}
	}

	lsi := make([]int, nlevel+1)

	lsi[0] = 0
	for i := 1; i < nlevel+1; i++ {
		lsi[i] = lsi[i-1] + i
	}

	var left, right int

	for level := len(lsi) - 2; level >= 0; level-- {
		for j := lsi[level]; j < lsi[level+1]; j++ {
			//left = leftchild(j, level)
			//right = rightchild(j, level)
			left, right = children(j, level)

			if left > len(nodes)-1 {
				// no children; leaf
				nodes[j].subtreeMax = nodes[j].value
			} else {
				nodes[j].subtreeMax = nodes[j].value +
					max(nodes[left].subtreeMax,
					nodes[right].subtreeMax)
			}
		}
	}

	fmt.Println("Max sum:", nodes[0].subtreeMax)
	fmt.Println("Tree:")
	printTree(nodes, lsi)
	fmt.Println("lsi", lsi)
}
