package problems_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/kunalrsagar/go-dsa/problems"
)

func TestSieveOfEratosthenes(t *testing.T) {
	tests := []struct {
		input int
		want  []int
	}{
		{10, []int{2, 3, 5, 7}},
		{23, []int{2, 3, 5, 7, 11, 13, 17, 19, 23}},
		{50, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47}},
		{100, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("test [%d]: input %d", i, tt.input)
		t.Run(testname, func(t *testing.T) {
			got := problems.SieveOfEratosthenes(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
