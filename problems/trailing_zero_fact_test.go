package problems_test

import (
	"fmt"
	"testing"

	"github.com/kunalrsagar/go-dsa/problems"
)

func TestTrailingZeroFact(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{5, 1},
		{6, 1},
		{2, 0},
		{10, 2},
		{7, 1},
		{30, 7},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("test [%d]: input %d", i, tt.input)
		t.Run(testname, func(t *testing.T) {
			got := problems.TrailingZeroFact(tt.input)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
