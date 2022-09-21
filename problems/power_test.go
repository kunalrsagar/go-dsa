package problems_test

import (
	"fmt"
	"testing"

	"github.com/kunalrsagar/go-dsa/problems"
)

func TestPower(t *testing.T) {
	tests := []struct {
		x    int
		n    int
		want int
	}{
		{2, 2, 4},
		{2, 4, 16},
		{5, 5, 3125},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("test [%d]: input x: %d n: %d", i, tt.x, tt.n)
		t.Run(testname, func(t *testing.T) {
			got := problems.Power(tt.x, tt.n)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
