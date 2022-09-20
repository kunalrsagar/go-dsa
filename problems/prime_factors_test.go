package problems_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/kunalrsagar/go-dsa/problems"
)

func TestPrimeFactors(t *testing.T) {
	tests := []struct {
		input int
		want  []int
	}{
		{450, []int{2, 3, 3, 5, 5}},
		{84, []int{2, 2, 3, 7}},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("test [%d]: input %d", i, tt.input)
		t.Run(testname, func(t *testing.T) {
			got := problems.PrimeFactors(tt.input)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
