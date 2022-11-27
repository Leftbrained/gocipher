package cipher

import (
	"fmt"
	"math"
)

type Playfair struct {
	size    int
	grid    [][]PlayfairCell
	letters map[byte]PlayfairCell
}

type PlayfairCell struct {
	letter byte
	x      int
	y      int
}

type PlayfairConfig struct {
	alphabet []byte
}

type PlayfairOption func(*PlayfairConfig)

func NewPlayfair(key []byte, opts ...PlayfairOption) (*Playfair, error) {
	cfg := &PlayfairConfig{
		alphabet: []byte("ABCDEFGHIKLMNOPQRSTUVWXYZ"),
	}

	for _, opt := range opts {
		opt(cfg)
	}

	for i, k := range key {
		if k < 65 || k > 90 {
			return nil, fmt.Errorf("invalid character in key: %s", string(k))
		}
		if k == 74 {
			key[i] = 73
		}
	}

	cfg.alphabet = GetKeyedAlphabet(key, cfg.alphabet)

	alphabetSize := len(cfg.alphabet)
	size := int(math.Sqrt(float64(alphabetSize)))

	if size*size != alphabetSize {
		return nil, fmt.Errorf("expecting alphabet size to be a perfect square, found: %d", alphabetSize)
	}

	c := Playfair{
		size:    size,
		grid:    make([][]PlayfairCell, size),
		letters: make(map[byte]PlayfairCell, alphabetSize),
	}

	for y := 0; y < c.size; y++ {
		c.grid[y] = make([]PlayfairCell, c.size)
	}

	var i int

	for _, k := range cfg.alphabet {
		cell := PlayfairCell{
			letter: k,
			x:      i % c.size,
			y:      i / c.size,
		}

		c.grid[cell.y][cell.x] = cell
		c.letters[cell.letter] = cell
		i++
	}

	return &c, nil
}

func PlayfairWithAlphabet(alphabet []byte) PlayfairOption {
	return func(cfg *PlayfairConfig) {
		cfg.alphabet = alphabet
	}
}

func (c *Playfair) crypt(text []byte, shift int) []byte {
	output := make([]byte, 0, len(text))

	for i := 0; i < len(text); {
		var cell1 *PlayfairCell
		cell1, i = c.getUnigram(text[i:], i)
		if cell1 == nil {
			// No more letters
			break
		}

		cell2, j := c.getUnigram(text[i:], i)
		if cell2 == nil || cell1.letter == cell2.letter {
			var cell PlayfairCell
			if cell1.letter == 88 {
				// Inserting a Q because cell1 is an X
				cell = c.letters[81]
			} else {
				// Inserting an X
				cell = c.letters[88]
			}
			cell2 = &cell
		} else {
			i = j
		}

		x1, y1, x2, y2 := cell1.x, cell1.y, cell2.x, cell2.y
		switch {
		case x1 == x2:
			y1, y2 = (y1+shift)%c.size, (y2+shift)%c.size
		case y1 == y2:
			x1, x2 = (x1+shift)%c.size, (x2+shift)%c.size
		default:
			x2, x1 = x1, x2
		}

		output = append(output, c.grid[y1][x1].letter, c.grid[y2][x2].letter)
	}

	return output
}

func (c *Playfair) Encrypt(text []byte) []byte {
	return c.crypt(text, 1)
}

func (c *Playfair) Decrypt(text []byte) []byte {
	return c.crypt(text, 4)
}

func (c *Playfair) getUnigram(text []byte, i int) (*PlayfairCell, int) {
	for _, from := range text {
		if from == 74 {
			// Converting J to I
			from = 73
		}
		i++
		if cell, ok := c.letters[from]; ok {
			return &cell, i
		}
	}

	return nil, i
}
