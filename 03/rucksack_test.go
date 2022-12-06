package main

import "testing"

type testCase struct {
	id    string
	input string
	want  byte
}

func TestFindDuplicate(t *testing.T) {
	tests := []testCase{
		{id: "A", input: "vJrwpWtwJgWrhcsFMMfFFhFp", want: 'p'},
		{id: "B", input: "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", want: 'L'},
		{id: "C", input: "PmmdzqPrVvPwwTWBwg", want: 'P'},
		{id: "D", input: "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", want: 'v'},
		{id: "E", input: "ttgJtRGJQctTZtZT", want: 't'},
		{id: "F", input: "CrZsJsPPZsGzwwsLwLmpwMDw", want: 's'},
	}

	for _, tc := range tests {
		got := findDuplicate(tc.input)
		if tc.want != got {
			t.Fatalf("Case %v failed.\tGot: %s\tWant:%s\n", tc.id, string(got), string(tc.want))
		}
	}
}

func TestScoreCharacter(t *testing.T) {
	tests := map[byte]int{
		'a': 1,
		'A': 27,
	}
	for k, v := range tests {
		got := scoreCharacter(k)
		want := v
		if want != got {
			t.Fatalf("Scoring %s failed.\tGot: %d\tWant:%d\n", string(k), got, want)
		}
	}
}
