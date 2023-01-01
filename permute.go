package gocipher

import "fmt"

func Permutations(n, k int8, handler func([]int8)) error {
	if n < 0 || n > 126 {
		return fmt.Errorf("n must be between 0 and 126, inclusively: n=%d k=%d", n, k)
	}
	if k < 0 || k > n {
		return fmt.Errorf("k must be between 0 and n, inclusively: n=%d k=%d", n, k)
	}

	indices := make([]int8, n)
	cycles := make([]int8, k)

	for i := int8(0); i < n; i++ {
		indices[i] = i
		if i < k {
			cycles[i] = n - i
		}
	}

	handler(indices[:k])

	var buffer int8
OUTER:
	for {
		for i := k - 1; i >= 0; i-- {
			cycles[i] -= 1
			if cycles[i] == 0 {
				buffer = indices[i]
				copy(indices[i:], indices[i+1:])
				indices[n-1] = buffer
				cycles[i] = n - i
			} else {
				indices[i], indices[n-cycles[i]] = indices[n-cycles[i]], indices[i]
				handler(indices[:k])
				continue OUTER
			}
		}
		break
	}

	return nil
}
