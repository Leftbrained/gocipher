package gocipher

import (
	"reflect"
	"sort"
	"testing"
)

func TestPermutations(t *testing.T) {
	out := make([][]int8, 6)
	var i int
	err := Permutations(3, 3, func(set []int8) {
		if i >= 6 {
			t.Fatalf("too much output")
		}
		out[i] = make([]int8, len(set))
		copy(out[i], set)
		// fmt.Printf("%v\n", set)
		i++
	})

	if err != nil {
		t.Fatalf(`returned error`)
	}

	sort.Slice(out, func(i, j int) bool {
		for k := 0; k < 3; k++ {
			if out[i][k] != out[j][k] {
				return out[i][k] < out[j][k]
			}
		}
		return true
	})

	if !reflect.DeepEqual(out, [][]int8{
		{0, 1, 2},
		{0, 2, 1},
		{1, 0, 2},
		{1, 2, 0},
		{2, 0, 1},
		{2, 1, 0},
	}) {
		t.Fatalf(`incorrect output`)
	}
}

func TestPermutationsErrorKGreaterThanN(t *testing.T) {
	err := Permutations(3, 4, func(i []int8) {})

	if err == nil {
		t.Fatalf("did not fail")
	}
}

func BenchmarkPermutations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Permutations(10, 10, func(i []int8) {})
	}
}
