package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func pairToInts(s string) (a int, b int) {
	parsed := strings.Split(s, "-")

	a, err := strconv.Atoi(parsed[0])
	if err != nil {
		panic(err)
	}

	b, err = strconv.Atoi(parsed[1])
	if err != nil {
		panic(err)
	}

	return
}

func first(file *os.File) (result int) {
	sc := bufio.NewScanner(file)


	for sc.Scan() {
		if line := sc.Text(); line != "" {
			parts := strings.Split(line, ",")
			a, b := pairToInts(parts[0])
			n, m := pairToInts(parts[1])

			if a <= n && b >= m || a >= n && b <= m {
				result++
			}
		}
	}

	return
}

func second(file *os.File) (result int) {
	sc := bufio.NewScanner(file)
  
	for sc.Scan() {
		if line := sc.Text(); line != "" {
			parts := strings.Split(line, ",")
			a, b := pairToInts(parts[0])
			n, m := pairToInts(parts[1])

			if b < n || a > m {
				continue
			}

			result++
		}
	}

	return
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	one := first(file)
	two := second(file)
	fmt.Println("Part one result:", one)
	fmt.Println("Part two result:", two)
}
