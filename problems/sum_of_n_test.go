package problems_test

import (
	"fmt"
	"testing"

	"github.com/kunalrsagar/go-dsa/problems"
)

func TestSumOfN(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{2, 3},
		{10, 55},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("test [%d]: input %d", i, tt.input)
		t.Run(testname, func(t *testing.T) {
			ans := problems.SumOfN(tt.input)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func TestSumOfNV2(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{2, 3},
		{10, 55},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("test [%d]: input %d", i, tt.input)
		t.Run(testname, func(t *testing.T) {
			ans := problems.SumOfNV2(tt.input)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func BenchmarkSumOfN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		problems.SumOfN(191919)
	}
}

func BenchmarkSumOfNV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		problems.SumOfNV2(191919)
	}
}
