package problems_test

import (
	"fmt"
	"testing"

	"github.com/kunalrsagar/go-dsa/problems"
)

func TestLCM(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{3, 7, 21},
		{4, 6, 2},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("test [%d]: input a %d; b %d", i, tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			got := problems.LCM(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
