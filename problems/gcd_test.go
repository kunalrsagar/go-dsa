package problems_test

import (
	"fmt"
	"testing"

	"github.com/kunalrsagar/go-dsa/problems"
)

func TestGcd(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{115, 125, 5},
		{49, 70, 7},
		{7, 11, 1},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("test [%d]: input %d:%d", i, tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			got := problems.Gcd(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
