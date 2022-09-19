package problems_test

import (
	"fmt"
	"testing"

	"github.com/kunalrsagar/go-dsa/problems"
)

func TestIsPrime(t *testing.T) {
	tests := []struct {
		input int
		want  bool
	}{
		{2, true},
		{11, true},
		{4, false},
		{7, true},
		{101, true},
		{100, false},
		{98, false},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("test [%d]: input %d", i, tt.input)
		t.Run(testname, func(t *testing.T) {
			got := problems.IsPrime(tt.input)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
