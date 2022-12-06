package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func findDuplicate(s string) byte {
	h1 := s[:len(s)/2]
	h2 := s[len(s)/2:]

	h1map := make(map[byte]bool)

	for _, c := range h1 {
		h1map[byte(c)] = true
	}
	for _, c := range h2 {
		if h1map[byte(c)] {
			return byte(c)
		}
	}

	return '0'
}

func scoreCharacter(c byte) int {
	if 64 < c && c <= 90 {
		return int(c - 64 + 26)
	}
	return int(c - 96)
}

func main() {
	input := "data.txt"
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)

	total := 0
	for s.Scan() {
		total += scoreCharacter(findDuplicate(s.Text()))
	}
	fmt.Println(total)
}
