package main

import (
	"bufio"
	"strings"
	"testing"
)

type roundTestCase struct {
	input string
	score int
}

func TestRoundScoring(t *testing.T) {
	tests := []roundTestCase{
		{input: "A Y", score: 8},
		{input: "A X", score: 4},
		{input: "B X", score: 1},
		{input: "C Z", score: 6},
	}
	for _, tc := range tests {
		got, want := ScoreRound(tc.input), tc.score
		if want != got {
			t.Fatalf("Case %v failed. Got %d\tWant:%d\n", tc.input, got, want)
		}
	}
}

func TestCorrectRoundScoring(t *testing.T) {
	input := "A Y\nB X\nC Z"

	want := 12
	s := bufio.NewScanner(strings.NewReader(input))
	got := ScoreTargetStrategy(s)
	if want != got {
		t.Fatalf("Correct scoring failed. Got: %d\tWant:%d\n", got, want)
	}
}

func TestStrategyScoring(t *testing.T) {
	input := "A Y\nB X\nC Z"

	want := 15
	s := bufio.NewScanner(strings.NewReader(input))
	got := ScoreStrategy(s)
	if want != got {
		t.Fatalf("Strategy scoring failed. Got: %d\tWant:%d\n", got, want)
	}
}
