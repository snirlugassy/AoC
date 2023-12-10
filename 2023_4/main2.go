package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parse_card_winnings(card_str string) int {
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

  var num_winning_numbers = 0

  for _, x := range my_numbers {
    if slices.Contains(winning_numbers, x) {
      num_winning_numbers++
    } 
  }

  return num_winning_numbers
}

func sum(arr []int) int {
  t := 0
  for _, x := range arr {
    t += x
  }
  return t
}

func main() {
  var input_file string = "input.txt"
  
	file, err := os.Open(input_file)

	if err != nil {
    panic("error loading file")
	}

	defer file.Close()

  scanner := bufio.NewScanner(file)
  
  // cards_winnings := make([]int, 0)
  cards_winnings := make(map[int]int) // idx -> number of winning cards
  track := make(map[int]int) // idx -> number of accumulated cards
  n := 0
  for ; scanner.Scan(); n++{
	  line := scanner.Text()
    cards_winnings[n] = parse_card_winnings(line)
    track[n] = 0
	}

  total_cards := 0

  for i:=n-1; i>=0; i-- {
    num_wins := cards_winnings[i]
    t := 1 // the i-card is first
    for j:=1; j<=num_wins; j++ {
      // add all cards earned below
      t += track[i+j]
    }
    track[i] = t
    total_cards += t
    fmt.Printf("i=%d, t=%d\n", i, t)
  }

  fmt.Printf("total cards: %d\n", total_cards)
}
