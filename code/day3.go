package code

import (
	"math"
	"fmt"
)

/*
	Reference for how the if logic works
	4 5 X 6 6
	4 X X X 7
	X X X X X
	3 X X X 8
	2 2 X 1 1

 */

func Day3(num int) (int){
	numRoot := math.Sqrt(float64(num))
	root := 0
	//get size of our "square"
	if int(numRoot) % 2 == 0 {
		root = int(numRoot) + 1
	} else {
		root = int(numRoot) +2
	}
	fmt.Printf("root: %d\n",root)
	rootSquare := root*root
	maxDistance := root -1
	minDistance := (root - 1) /2
	bottom := rootSquare - (root-1)/2
	left := rootSquare - 3*((root-1)/2)
	top := rootSquare - 5*((root-1)/2)
	right := rootSquare - 7*((root-1)/2)

	//Check for edge cases
	//fmt.Printf("b: %d\tl: %d\tt: %d\tr: %d\t\n",bottom,left,top,right)
	if num == bottom || num == left || num == top || num == right {
		return minDistance
	}


	//Use the reference square above to understand better
	if num > bottom {
		// This one checks the 1
		return maxDistance - (rootSquare - num)
	} else if num >= bottom - minDistance && num < bottom {
		// This one checks the 2
		return minDistance + (bottom - num)
	} else if  num > left && num < bottom - minDistance{
		// This one checks the 3
		return maxDistance - (bottom - minDistance - num)
	} else if num >= left - minDistance && num < left {
		// This one checks the 4
		return minDistance + (left - num)
	} else if num > top && num < left - minDistance {
		// This one checks the 5
		return maxDistance - (left - minDistance - num)
	} else if num >= top - minDistance && num < top {
		// This one checks the 6
		return minDistance + (top - num)
	} else if num > right && num < top - minDistance {
		// This one checks the 7
		return maxDistance - (num - right)
	} else if num > right - minDistance && num < right {
		// This one checks the 8
		return minDistance + (right - num)
	}

	return 0

}

func Day3part2(num,size int) {
	arr := make([][]int,size)
	for i,j := size/2,size/2; i*j<size*size; {
		
	}
}