package code

import (
	"fmt"
	"strconv"
)

func Day1part2() {
	stringData := loadFile("day1.txt")
	sum := 0
	steps := len(stringData)/2
	for i := 0; i< len(stringData); i++ {
		n1, _ := strconv.Atoi(string(stringData[i]))
		n2, _ := strconv.Atoi(string(stringData[(i+steps )% len(stringData)]))
		if n1 == n2 {
			sum += n1
		}
	}

	fmt.Println(sum)
}