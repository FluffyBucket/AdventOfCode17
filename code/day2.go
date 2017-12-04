package code

import (
	"fmt"
	"strings"
	"os"
	"bufio"
	"strconv"
)

func Day2() {
	sum := 0
	f,_ := os.Open("day2.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		first,_ := strconv.Atoi(parts[0])
		small,large := first,first

		for i := range parts {
			fmt.Println(parts[i])
			num, _ := strconv.Atoi(parts[i])
			if num > large {
				large = num
			} else if num < small {
				small = num
			}

		}

		sum += (large - small)

		fmt.Println()
	}

	fmt.Println(sum)
}

func Day2part2() {
	sum := 0
	f,_ := os.Open("day2.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		//first,_ := strconv.Atoi(parts[0])
		//small,large := first,first

		for i := range parts {
			numi, _ := strconv.Atoi(parts[i])
			for j := i+1;j<len(parts);j++ {
				numj, _ := strconv.Atoi(parts[j])
				if numi % numj == 0 {
					sum +=	numi / numj
				} else if numj % numi == 0 {
					sum += numj /numi
				}
			}
			//fmt.Println(parts[i])



		}
	}

	fmt.Println(sum)
}
