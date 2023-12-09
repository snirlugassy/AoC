package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func read_input_matrix(input_file string) ([][]byte, error) {
	file, err := os.Open(input_file)

	if err != nil {
		return nil, fmt.Errorf("error loading file %s", input_file)
	}

	defer file.Close()

	var lines [][]byte
	lines = make([][]byte, 0)

    scanner := bufio.NewScanner(file)
	var m int = -1
    for i:=0; scanner.Scan(); i++{
		line := []byte(scanner.Text())
		line = bytes.Join([][]byte{[]byte("."), line, []byte(".")}, []byte(""))
		lines = append(lines, line)

		if m == -1 {
			m = len(line)
		}

		if len(line) != m {
			return nil, fmt.Errorf("inconsistent line size at line %d\n", i)
		}
	}

	// add padding of dots, more convenient in lookup later
	padding_line := []byte(strings.Repeat(".", m))

	lines = append(lines, padding_line)
	lines = append([][]byte {padding_line}, lines...)

	return lines, nil
}

func extract_number(buf []byte) (int, error) {
  for i,c := range buf {
    if c < '0' || c > '9' {
      if i == 0 {
        return -1, fmt.Errorf("buffer doesn't start with a digit")
      }
      return strconv.Atoi(string(buf[:i]))
    }
  }

  // the entire buffer is a number
  return strconv.Atoi(string(buf))
}

func main() {
	var input_file_path string = "input.txt"

	var input [][]byte
	var err error

	input, err = read_input_matrix(input_file_path)

	if err != nil {
		panic(err)
	}

	n := len(input)
	m := len(input[0])

	// print input information
	fmt.Printf("n = %d, m = %d\n", n, m)
	for i := 0; i < n; i++ {
		fmt.Printf("line %d : %s\n", i, input[i])
		// lines[i] = []byte{".", lines[i]...}
	}

	// solve
  var part_numbers_sum int = 0
	for i := 1; i < n-1; i++ {
    // indicator to avoid parsing part of a number
    var inside_number bool = false
    var is_part_number bool = false
    var number int

  	for j := 1; j < m-1; j++ {
      char := rune(input[i][j])

      if char < '0' || char > '9' {
				// not a digit
        inside_number = false
        is_part_number = false
				continue
			}

      if inside_number && is_part_number {
        // avoid checking part number that we already know
        // it is a part number
        continue
      }

      if !inside_number {
        // first time we parse this number
        number, err = extract_number(input[i][j:])
        if err != nil {
          continue
        }
      }

      inside_number = true
      fmt.Printf("found a digit at (%d, %d) = %c | number is %d\n", i, j, char, number)

      // check if adjacent to a symbol
      row_loop:
      for ii := i-1; ii<=i+1; ii++ {
        for jj := j-1; jj<=j+1; jj++ {
          if input[ii][jj] != '.' && (input[ii][jj] < '0' || input[ii][jj] > '9') {
            is_part_number = true
            fmt.Printf("%d is a part number because %c\n", number, rune(input[ii][jj]))
            break row_loop
          }
        }
      }

      if is_part_number {
        part_numbers_sum += number
      }

		}
	}

  fmt.Printf("sum of part numbers: %d\n", part_numbers_sum)

}
