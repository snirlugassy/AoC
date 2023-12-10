package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parse_card_points(card_str string) int {
  var card_str_parts []string = strings.Split(card_str, ":")
  var numbers_parts []string = strings.Split(card_str_parts[1], "|")
  
  if len(card_str_parts) != 2 || len(numbers_parts) != 2{
    fmt.Println(card_str)
    panic("invalid line!")
  }

  var winning_numbers_str string = numbers_parts[0]
  var my_numbers_str string = numbers_parts[1]

  // handle duplicate whitespaces
  winning_numbers_str = strings.ReplaceAll(winning_numbers_str, "  ", " ")
  my_numbers_str = strings.ReplaceAll(my_numbers_str, "  ", " ")
 
  // trim whitespace
  winning_numbers_str = strings.Trim(winning_numbers_str, " ")
  my_numbers_str = strings.Trim(my_numbers_str, " ")

  var winning_numbers []int = make([]int, 0)
  for _, num := range strings.Split(winning_numbers_str, " ") {
    x,err := strconv.Atoi(num)
    if err != nil {
      fmt.Printf("%s\n", num)
      panic("can't parse number")
    }
    winning_numbers = append(winning_numbers, x)
  }

  var my_numbers []int = make([]int, 0)
  for _, num := range strings.Split(my_numbers_str, " ") {
    x,err := strconv.Atoi(num)
    if err != nil {
      panic("can't parse number")
    }
    my_numbers = append(my_numbers, x)
  }

  var score = 0

  for _, x := range my_numbers {
    if slices.Contains(winning_numbers, x) {
      score++
    } 
  }

  return int(math.Pow(2, float64(score) - 1))
}

func main() {
  var input_file string = "input.txt"
  
	file, err := os.Open(input_file)

	if err != nil {
    panic("error loading file")
	}

	defer file.Close()

  scanner := bufio.NewScanner(file)
  
  total_points := 0
  for i:=0; scanner.Scan(); i++{
	  line := scanner.Text()
    total_points += parse_card_points(line)
	}

  fmt.Printf("total points: %d\n", total_points)
}
