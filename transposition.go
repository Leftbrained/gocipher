package gocipher

import (
	"sort"
)

type Transposition struct {
	keyLen   int
	original []int
	sorted   []int
}

type transpositionKeyColumn struct {
	letter   byte
	original int
}

type transpositionMessageColumn struct {
	size    int
	encrypt int
	decrypt int
}

func NewTransposition(key []byte) (*Transposition, error) {

	size := len(key)
	c := Transposition{
		keyLen:   size,
		original: make([]int, size),
		sorted:   make([]int, size),
	}

	cols := make([]transpositionKeyColumn, c.keyLen)

	for i, k := range key {
		cols[i] = transpositionKeyColumn{
			letter:   k,
			original: i,
		}
	}

	sort.SliceStable(cols, func(i, j int) bool {
		return cols[i].letter < cols[j].letter
	})

	for sorted, col := range cols {
		c.original[col.original] = sorted
		c.sorted[sorted] = col.original
	}

	return &c, nil
}

func (c *Transposition) Encrypt(fromtext []byte) []byte {
	size := len(fromtext)
	totext := make([]byte, size)

	offsets := c.getOffsets(size)

	for i := 0; i < size; i++ {
		totext[offsets[i%c.keyLen]+(i/c.keyLen)] = fromtext[i]
	}

	return totext
}

func (c *Transposition) Decrypt(fromtext []byte) []byte {
	size := len(fromtext)
	totext := make([]byte, size)

	offsets := c.getOffsets(size)

	for i := 0; i < size; i++ {
		totext[i] = fromtext[offsets[i%c.keyLen]+(i/c.keyLen)]
	}

	return totext
}

func (c *Transposition) getOffsets(size int) []int {
	remainder := size % c.keyLen
	quotient := size / c.keyLen
	offsets := make([]int, c.keyLen)

	for original, sorted := range c.original {
		for s := 0; s < sorted; s++ {
			offsets[original] += quotient
			if c.sorted[s] < remainder {
				offsets[original]++
			}
		}
	}

	return offsets
}
