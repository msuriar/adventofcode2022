package main

import (
	"bufio"
	"fmt"
  "log"
  "os"
	"strconv"
)

func main() {
	input := "data.txt"
  file, err := os.Open(input)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

	s := bufio.NewScanner(file)
	fmt.Println(maxCalories(s))
}

func maxCalories(s *bufio.Scanner) int {
  max := 0
  acc := 0
  for s.Scan() {
    if s.Text() == "" {
      if acc > max {
        max, acc = acc, 0
      }
      acc = 0
    } else {
      i, _ := strconv.Atoi(s.Text())
      acc += i
    }
  }
  if acc > max {
    max = acc
  }
  return max
}
