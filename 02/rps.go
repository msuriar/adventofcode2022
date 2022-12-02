package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input := "data.txt"
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	/*	s := bufio.NewScanner(file)
		total := ScoreStrategy(s)
		fmt.Println(total)*/

	s := bufio.NewScanner(file)
	total := ScoreTargetStrategy(s)
	fmt.Println(total)
}

func ScoreStrategy(s *bufio.Scanner) int {
	score := 0
	for s.Scan() {
		score += ScoreRound(s.Text())
	}
	return score
}

func ScoreTargetStrategy(s *bufio.Scanner) int {
	score := 0
	for s.Scan() {
		score += ScoreTargetRound(s.Text())
	}
	return score
}

func ScoreRound(r string) int {
	return scoreOutcome(r) + scoreShape(r)
}

func ScoreTargetRound(r string) int {
	return scoreTargetOutcome(r) + scoreTargetShape(r)
}

func scoreOutcome(r string) int {
	switch r {
	case "A Y":
		return 6
	case "B Z":
		return 6
	case "C X":
		return 6
	case "A X":
		return 3
	case "B Y":
		return 3
	case "C Z":
		return 3
	}

	return 0
}

func scoreTargetOutcome(r string) int {
	switch r[2] {
	case 'X':
		return 0
	case 'Y':
		return 3
	default:
		return 6
	}
}

func scoreShape(r string) int {
	switch r[2] {
	case 'X':
		return 1
	case 'Y':
		return 2
	default:
		return 3
	}
}

func scoreTargetShape(r string) int {
	switch r[0] {
	case 'A':
		return scoreRock(r[2])
	case 'B':
		return scorePaper(r[2])
	default:
		return scoreScissors(r[2])
	}
}

func scoreRock(o byte) int {
	switch o {
	case 'X':
		return 3
	case 'Y':
		return 1
	default:
		return 2
	}
}

func scorePaper(o byte) int {
	switch o {
	case 'X':
		return 1
	case 'Y':
		return 2
	default:
		return 3
	}
}

func scoreScissors(o byte) int {
	switch o {
	case 'X':
		return 2
	case 'Y':
		return 3
	default:
		return 1
	}
}
