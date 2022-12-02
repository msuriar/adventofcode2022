package main

import (
	"bufio"
	"strings"
	"testing"
)

type testCase struct {
	id    string
	input string
	want  int
}

func TestCalories(t *testing.T) {
	tests := []testCase{
		{id: "A", input: "1\n2\n3\n\n4\n5\n6", want: 15},
		{id: "B", input: "1\n2\n3\n\n3\n2", want: 6},
		{id: "C", input: "1000\n2000\n3000\n\n4000\n\n5000\n6000\n\n7000\n8000\n9000\n\n10000", want: 24000},
	}
	for _, tc := range tests {
		s := bufio.NewScanner(strings.NewReader(tc.input))
		totals := sumCalories(s)
		got := maxCalories(totals)
		if tc.want != got {
			t.Fatalf("Case %v failed. Got: %d\tWant:%d\n", tc.id, got, tc.want)
		}
	}
}
