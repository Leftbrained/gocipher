package gocipher

type Cipher interface {
	Encrypt([]byte) []byte
	Decrypt([]byte) []byte
}

func GetKeyedAlphabet(key, alphabet []byte) []byte {
	output := make([]byte, 0, len(alphabet))
	lookup := make(map[byte]bool, len(alphabet))

	for _, b := range alphabet {
		lookup[b] = false
	}

	// TODO: Benchmark the append
	for _, b := range key {
		if _, ok := lookup[b]; !ok {
			// Skipping byte not present in alphabet
			continue
		}

		if lookup[b] {
			// Already added to output
			continue
		}

		lookup[b] = true
		output = append(output, b)
	}

	for _, b := range alphabet {
		if lookup[b] {
			// Already added to output
			continue
		}

		lookup[b] = true
		output = append(output, b)
	}

	return output
}
