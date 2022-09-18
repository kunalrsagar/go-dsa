package problems_test

import (
	"fmt"
	"testing"

	"github.com/kunalrsagar/go-dsa/problems"
)

func TestCountDigits(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{1919, 4},
		{19, 2},
		{19000, 5},
		{8, 1},
		{10101010, 8},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("test [%d]: input %d", i, tt.input)
		t.Run(testname, func(t *testing.T) {
			ans := problems.CountDigits(tt.input)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
