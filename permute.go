package gocipher

import "fmt"

func Permutations(n, k int8, handler func([]int8)) error {
	if n < 0 || n > 126 {
		return fmt.Errorf("n must be between 0 and 126, inclusively: n=%d k=%d", n, k)
	}
	if k < 0 || k > n {
		return fmt.Errorf("k must be between 0 and n, inclusively: n=%d k=%d", n, k)
	}

	if n == k {
		return permutationsQuickPermCountdown(n, handler)
	}

	return permutationsPython(n, k, handler)
}

func permutationsPython(n, k int8, handler func([]int8)) error {
	// https://docs.python.org/3/library/itertools.html#itertools.permutations

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

func permutationsQuickPermCountdown(n int8, handler func([]int8)) error {
	// https://www.quickperm.org/

	var i, j int8
	indices := make([]int8, n)
	p := make([]int8, n+1)

	for i = 0; i < n; i++ {
		indices[i] = i
		p[i] = i
	}
	p[n] = n

	handler(indices)

	for i = 1; i < n; {
		p[i]--

		if i%2 == 1 {
			j = p[i]
		} else {
			j = 0
		}

		indices[i], indices[j] = indices[j], indices[i]
		handler(indices)

		for i = 1; p[i] == 0; i++ {
			p[i] = i
		}
	}
	return nil
}
