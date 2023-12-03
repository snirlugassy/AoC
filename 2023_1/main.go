package main

import (
	"fmt"
	"os"
	"strings"
)

func is_digit(a rune) bool {
	if a >= '0' && a <= '9' {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println("AOC 2023 Q1")

	input_file := "input.txt"
	input, err := os.ReadFile(input_file)
	if err != nil {
		panic("failed to read file")
	}

	var total int = 0
	input_lines := strings.Split(string(input), "\n")
	for i, line := range input_lines {
		fmt.Printf("%d: %s\n", i, line)

		var first_digit int = -1
		var last_digit int = -1

		for _, char := range line {
			if is_digit(char) {
				if first_digit == -1 {
					// first digit in line
					first_digit = int(char - '0')
				}
				
				last_digit = int(char - '0')
			}
		}
		
		fmt.Printf("%d -> %d\n", first_digit, last_digit)

		total += first_digit*10 + last_digit
	}

	fmt.Printf("total = %d\n", total)
}