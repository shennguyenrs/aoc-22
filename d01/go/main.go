package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func parseFileToSlice(file *os.File) []int64 {
	var results []int64
	var temp int64

	sc := bufio.NewScanner(file)

	// Check each lines and add numbers to temp
	// If line is empty, reset temp and append result to slice
	for sc.Scan() {
		if line := sc.Text(); line != "" {
			if n, err := strconv.ParseInt(line, 10, 64); err == nil {
				temp += n
			}
		} else {
			results = append(results, temp)
			temp = 0
		}
	}

	return results
}

func findMaximum(arr []int64) (int64, int) {
	var max int64 = 0
	var index int = 0

	max = arr[0]

	// Find the largest numbers in slice
	for i := range arr {
		if arr[i] > max {
			max = arr[i]
			index = i
		}
	}

	return max, index
}

func first(file *os.File) int64 {
	// Parse file to slice
	lines := parseFileToSlice(file)
	max, _ := findMaximum(lines)
	return max
}

func second(file *os.File) int64 {
	lines := parseFileToSlice(file)
	var n int = 3
	topThree := [3]int64{}
	var result int64 = 0

	for i := 0; i < n; i += 1 {
		max, index := findMaximum(lines)

		log.Println(max)

		topThree[i] = max
		next := index + 1
		lines = append(lines[:index], lines[next:]...)
		result += max
	}

	return result
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	// fmt.Println("The maxium calories:", first(file))
	fmt.Println("The sum of top three calories:", second(file))
}
