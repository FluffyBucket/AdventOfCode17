package code

import (
	"os"
	"bufio"
	"strings"
	"fmt"
)

func Day4() {
	sum := 0
	f, _ := os.Open("day4.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		valid := 1

		m := make(map[string]int)
		for i := range parts {
			//fmt.Println(parts[i])
			if _, ok := m[parts[i]]; ok {
				valid = 0
				break
			} else {
				m[parts[i]] = 1
			}
		}
		sum += valid
	}

	fmt.Println(sum)
}

func Day4part2() {
	sum := 0
	f, _ := os.Open("day4.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		valid := 1

		m := make(map[string]int)
		for i := range parts {
			//fmt.Println(parts[i])
			sortedString := SortString(parts[i])
			if _, ok := m[sortedString]; ok {
				valid = 0
				break
			} else {
				//fmt.Println(sortedString)
				m[sortedString] += 1
			}
		}
		sum += valid
	}
	fmt.Println(sum)
}

func SortString(s string) string {
	var items = make([]int, len(s))
	for i, v := range s {
		items[i] = int(v)
	}
	items = Mergesort(items)
	rets := ""
	for _, v := range items {
		rets += string(rune(v))
	}
	return rets
}

func merge(a []int, b []int) []int {

	var r = make([]int, len(a)+len(b))
	var i = 0
	var j = 0

	for i < len(a) && j < len(b) {

		if a[i] <= b[j] {
			r[i+j] = a[i]
			i++
		} else {
			r[i+j] = b[j]
			j++
		}

	}

	for i < len(a) {
		r[i+j] = a[i];
		i++
	}
	for j < len(b) {
		r[i+j] = b[j];
		j++
	}

	return r

}

func Mergesort(items []int) []int {

	if len(items) < 2 {
		return items

	}

	var middle = len(items) / 2
	var a = Mergesort(items[:middle])
	var b = Mergesort(items[middle:])
	return merge(a, b)

}
