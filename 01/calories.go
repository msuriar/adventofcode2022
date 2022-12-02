package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
	totals := sumCalories(s)
	fmt.Println(maxCalories(totals))
}

func sumCalories(s *bufio.Scanner) (out []int) {
	acc := 0
	for s.Scan() {
		if s.Text() == "" {
			out = append(out, acc)
			acc = 0
		} else {
			i, _ := strconv.Atoi(s.Text())
			acc += i
		}
		out = append(out, acc)
	}
	sort.Ints(out)
	return out
}

func maxCalories(totals []int) int {
	return totals[len(totals)-1]
}
