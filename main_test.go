package main

import "testing"

var testCases = []struct {
	in   string
	want int
}{
	{in: "", want: 1},
}

func TestKintai(t *testing.T) {
	for _, tc := range testCases {
		result := kintai(tc.in)
		if result != tc.want {
			t.Fatal("failed in arg validation")
		}
	}
}
