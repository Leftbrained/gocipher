package gocipher

import (
	"fmt"
	"testing"
)

func TestPermutations(t *testing.T) {
	Permutations(4, func(i []int8) {
		fmt.Printf("  %v\n", i)
	})
}
