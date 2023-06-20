package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	var numberOfIncreases int

	sc.Scan()
	prev2, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	prev1, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	prev0, _ := strconv.Atoi(sc.Text())

	for sc.Scan() {
		depth, _ := strconv.Atoi(sc.Text())
		if depth > prev2 {
			numberOfIncreases++
		}
		prev2 = prev1
		prev1 = prev0
		prev0 = depth
	}

	fmt.Println(numberOfIncreases)
}
