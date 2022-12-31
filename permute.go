package gocipher

// type Permute struct{}

// type PermuteConfig struct{}

// type PermuteOption func(*PermuteConfig)

// func NewPermute() (*Permute, error) {
// 	return &Permute{}, nil
// }

// func (p *Permute) next() {

// }

// func (p *Permute) Permutations() {

// }

// func Permutations[T interface{}](pool []T, length int, handler func([]T)) {
// 	n := len(pool)
// 	if length > n {
// 		return
// 	}

// }

// https://docs.python.org/3/library/itertools.html#itertools.permutations
func Permutations(size int8, handler func([]int8)) {
	indices := make([]int8, size) // [0, 1, 2]
	cycles := make([]int8, size)  // [3, 2, 1]

	for i := int8(0); i < size; i++ {
		indices[i] = i
		cycles[i] = size - i
	}

	handler(indices)

	var buffer int8
OUTER:
	for {
		for i := size - 1; i >= 0; i-- {
			cycles[i] -= 1
			if cycles[i] == 0 {
				buffer = indices[i]
				copy(indices[i:size-1], indices[i+1:]) // I think indices[i:size-1] can be indices[i:]
				indices[size-1] = buffer
				cycles[i] = size - i
			} else {
				indices[i], indices[size-cycles[i]] = indices[size-cycles[i]], indices[i]
				handler(indices)
				continue OUTER
			}
		}
		break
	}
}
