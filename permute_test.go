package gocipher

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestPermutationsFull(t *testing.T) {
	testPermutations(t, 3, 3, [][]int8{
		{0, 1, 2},
		{0, 2, 1},
		{1, 0, 2},
		{1, 2, 0},
		{2, 0, 1},
		{2, 1, 0},
	})
}

func TestPermutationsPartial(t *testing.T) {
	testPermutations(t, 3, 2, [][]int8{
		{0, 1},
		{0, 2},
		{1, 0},
		{1, 2},
		{2, 0},
		{2, 1},
	})
}

func testPermutations(t *testing.T, n, k int8, expected [][]int8) {
	out := make([][]int8, 0)
	err := PermutationsPartial(n, k, func(set []int8) {
		setCopy := make([]int8, len(set))
		copy(setCopy, set)
		out = append(out, setCopy)
	})

	if err != nil {
		t.Fatalf(`returned error`)
	}

	sort.Slice(out, func(i, j int) bool {
		for k := 0; k < len(out[i]); k++ {
			if out[i][k] != out[j][k] {
				return out[i][k] < out[j][k]
			}
		}
		return true
	})

	if !reflect.DeepEqual(out, expected) {
		t.Fatalf(`incorrect output`)
	}
}

func TestPermutationsErrorKGreaterThanN(t *testing.T) {
	err := PermutationsPartial(3, 4, func(i []int8) {})

	if err == nil {
		t.Fatalf("did not fail")
	}
}

func BenchmarkPermutations(b *testing.B) {
	benchmarks := []struct {
		n int8
		k int8
	}{
		{10, 9},
		{10, 10},
	}

	for _, bm := range benchmarks {
		b.Run(fmt.Sprintf("n=%d/k=%d", bm.n, bm.k), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = PermutationsPartial(bm.n, bm.k, func(i []int8) {})
			}
		})
	}
}
