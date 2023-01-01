package gocipher

import "fmt"

func Permutations(n, k int8, handler func([]int8)) error {
	return permutationsQuickPerm(n, k, handler)
}

func permutationsPython(n, k int8, handler func([]int8)) error {
	// https://docs.python.org/3/library/itertools.html#itertools.permutations
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

func permutationsQuickPerm(length, k int8, handler func([]int8)) error {
	// https://www.baeldung.com/cs/array-generate-all-permutations
	if length < 0 || length > 126 {
		return fmt.Errorf("n must be between 0 and 126, inclusively: length=%d k=%d", length, k)
	}
	if k != length {
		return fmt.Errorf("k must be equal to length: length=%d k=%d", length, k)
	}

	elementsToPermute := make([]int8, length)
	p := make([]int8, length+1)

	for i := int8(0); i < length; i++ {
		p[i] = i
		elementsToPermute[i] = i
	}
	p[length] = length

	handler(elementsToPermute)

	var j int8
	for index := int8(1); index < length; {
		p[index] -= 1

		if index%2 == 1 {
			j = p[index]
		} else {
			j = 0
		}
		elementsToPermute[index], elementsToPermute[j] = elementsToPermute[j], elementsToPermute[index]
		handler(elementsToPermute)

		index = 1
		for p[index] == 0 {
			p[index] = index
			index++
		}
	}
	return nil
}
