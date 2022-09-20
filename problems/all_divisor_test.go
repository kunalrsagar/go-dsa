package problems_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/kunalrsagar/go-dsa/problems"
)

func TestAllDivisorOfNumber(t *testing.T) {
	tests := []struct {
		input int
		want  []int
	}{
		{15, []int{1, 3, 5, 15}},
		{100, []int{1, 2, 4, 5, 10, 20, 25, 50, 100}},
		{7, []int{1, 7}},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("test [%d]: input %d", i, tt.input)
		t.Run(testname, func(t *testing.T) {
			got := problems.AllDivisorsOfNumber(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
