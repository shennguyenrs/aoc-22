package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func convertRuneToValue(char rune) int {
	if unicode.IsLower(char) {
		return int(char - 'a' + 1)
	} else {
		return int(char - 'A' + 27)
	}
}

func checkAppearReturnValue(a []byte, b []byte) int {
	for _, i := range a {
		for _, j := range b {
			if i == j {
				char := rune(i)
				return convertRuneToValue(char)
			}
		}
	}

	return 0
}

func stringToUniqueByte(s string) []byte {
	b := make([]byte, 0)
	seen := make(map[rune]bool)

	for _, r := range s {
		if !seen[r] {
			b = append(b, byte(r))
			seen[r] = true
		}
	}

	return b
}

func first(file *os.File) int {
	var result int = 0

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		if line := sc.Text(); line != "" {
			length := len(line)
			firstPart := line[length/2:]
			secondPart := line[:length/2]
			firstPartByte := stringToUniqueByte(firstPart)
			secondPartByte := stringToUniqueByte(secondPart)

			temp := checkAppearReturnValue(firstPartByte, secondPartByte)

			if temp != 0 {
				result += temp
			}
		}
	}

	return result
}

func second(file *os.File) int {
	var result int = 0
	lines := make([]string, 3)
	seen := make(map[rune]int)

	sc := bufio.NewScanner(file)

	for i := 0; sc.Scan(); i += 1 {
		if line := sc.Text(); line != "" {
			lines[i] = line

			if i == 2 {
				seen = make(map[rune]int)

				// Iterate the first line
				for _, c := range lines[0] {
					seen[c] |= 1
				}

				// Iterate the second line
				for _, c := range lines[1] {
					seen[c] |= 2
				}

				// Iterate the third line
				// Used 4 because i after the first and second line will be 3
				for _, c := range lines[2] {
					seen[c] |= 4
				}

				// Iterate the first line and get the charater that has been seen in all the three lines
				for _, c := range lines[0] {
					if seen[c] == 7 {
						temp := convertRuneToValue(c)
						result += temp
						break
					}
				}

				// Start the iteration again
				i = -1
			}
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
	val := second(file)
	fmt.Println(val)
}
