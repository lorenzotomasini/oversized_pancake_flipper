package main

import (
	"bufio"

	"os"
	"strconv"
	"strings"
	"fmt"
)

func main() {
	file, err := os.Open("A-large-practice.in")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewReader(file)

	lines, _, err := scanner.ReadLine()
	ls, err := strconv.Atoi(string(lines))
	if err != nil {
		// handle error
		panic(err)
	}
	for x := 1; x <= ls; x++ {
		line, _, err := scanner.ReadLine();
		if err != nil {
			// handle error
			panic(err)
		}
		split := strings.Split(string(line), " ")
		pancakes := strings.Split(split[0], "")
		size, err := strconv.Atoi(split[1])
		if err != nil {
			// handle error
			panic(err)
		}

		res := countFlips(pancakes, size)
		if res == -1 {
			fmt.Println(fmt.Sprintf("Case #%d: IMPOSSIBLE", x))
		} else {
			fmt.Println(fmt.Sprintf("Case #%d: %d", x, res))
		}
	}
}
func countFlips(pancakes []string, size int) int {
	flips := 0
	for i, c := range pancakes {
		if i == len(pancakes)-size+1 {
			if allPlus(pancakes[i:]) {
				return flips
			}
			return -1
		}
		if c == "-" {
			flip(pancakes[i:i+size])
			flips += 1
		}
		//fmt.Println(fmt.Sprintf("pancakes: %v, flips: %d", pancakes, flips))
	}
	return flips;
}

func flip(pancakes []string) {
	for x := 0; x < len(pancakes); x++ {
		if pancakes[x] == "+" {
			pancakes[x] = "-"
		} else {
			pancakes[x] = "+"
		}
	}
}
func allPlus(pancakeSlice []string) (bool) {
	for _, c := range pancakeSlice {
		if c == "-" {
			return false
		}
	}
	return true
}
