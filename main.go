package main

import (
	"github.com/FluffyBucket/AdventOfCode17/code"
	"time"
	"fmt"
)

func main() {
	//fmt.Printf("distance: %d",code.Day3(325489))
	start := time.Now()
	code.Day4part2()
	end := time.Since(start)

	fmt.Println(end)

}
