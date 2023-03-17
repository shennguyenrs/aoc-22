package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type RULE struct {
	Judge   [3]int
	Choosen int
}

var rules = map[string]RULE{
	"X": {
		Judge:   [3]int{3, 0, 6},
		Choosen: 1,
	},
	"Y": {
		Judge:   [3]int{6, 3, 0},
		Choosen: 2,
	},
	"Z": {
		Judge:   [3]int{0, 6, 3},
		Choosen: 3,
	},
}

var opponentPos = map[string]int{
	"A": 0,
	"B": 1,
	"C": 2,
}

var secondStrategyScore = map[string]int{
	"X": 0,
	"Y": 3,
	"Z": 6,
}

var secondStrategyGuide = map[string]map[string]string{
	"A": {
		"X": "Z",
		"Y": "X",
		"Z": "Y",
	},
	"B": {
		"X": "X",
		"Y": "Y",
		"Z": "Z",
	},
	"C": {
		"X": "Y",
		"Y": "Z",
		"Z": "X",
	},
}

func first(file *os.File) int {
	var result int = 0

	// Read file
	sc := bufio.NewScanner(file)

	// Go through line by line and sum the result
	for sc.Scan() {
		if line := sc.Text(); line != "" {
			parts := strings.Fields(line)
			me := parts[1]
			opp := parts[0]
			result += rules[me].Choosen + rules[me].Judge[opponentPos[opp]]
		}
	}

	return result
}

func second(file *os.File) int {
	var result int = 0

	// Read file
	sc := bufio.NewScanner(file)

	// Go through line by line and sum the result
	for sc.Scan() {
		if line := sc.Text(); line != "" {
			parts := strings.Fields(line)
			me := parts[1]
			opp := parts[0]
			result += secondStrategyScore[me] + rules[secondStrategyGuide[opp][me]].Choosen
		}
	}

	return result
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	// val := first(file)
	val := second(file)
	fmt.Println(val)
}
