package gocipher

import (
	"bytes"
	"fmt"
)

type Playfair struct {
	grid    [][]PlayfairCell
	letters map[byte]PlayfairCell
}

type PlayfairCell struct {
	letter byte
	x      int
	y      int
}

func NewPlayfair(key []byte) (*Playfair, error) {
	c := Playfair{
		grid:    make([][]PlayfairCell, 5),
		letters: make(map[byte]PlayfairCell, 25),
	}

	for y := 0; y < 5; y++ {
		c.grid[y] = make([]PlayfairCell, 5)
	}

	var i int

	for _, k := range key {
		if k < 65 || k > 90 {
			return nil, fmt.Errorf("invalid character in key: %s", string(k))
		}
		if k == 74 {
			k = 73
		}

		if _, ok := c.letters[k]; ok {
			continue
		}

		cell := PlayfairCell{
			letter: k,
			x:      i % 5,
			y:      i / 5,
		}

		c.grid[cell.y][cell.x] = cell
		c.letters[cell.letter] = cell
		i++
	}

	for _, k := range []byte("ABCDEFGHIKLMNOPQRSTUVWXYZ") {
		if _, ok := c.letters[k]; ok {
			continue
		}

		cell := PlayfairCell{
			letter: k,
			x:      i % 5,
			y:      i / 5,
		}

		c.grid[cell.y][cell.x] = cell
		c.letters[cell.letter] = cell
		i++
	}

	return &c, nil
}

func (c *Playfair) crypt(text []byte, swap func(x1, y1, x2, y2 int) (int, int, int, int)) []byte {
	buffer := bytes.NewBuffer([]byte{})

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

		x1, y1, x2, y2 := swap(cell1.x, cell1.y, cell2.x, cell2.y)

		buffer.WriteByte(c.grid[y1][x1].letter)
		buffer.WriteByte(c.grid[y2][x2].letter)
	}

	return buffer.Bytes()
}

func (c *Playfair) Encrypt(text []byte) []byte {
	return c.crypt(text, func(x1, y1, x2, y2 int) (int, int, int, int) {
		switch {
		case x1 == x2:
			y1, y2 = (y1+1)%5, (y2+1)%5
		case y1 == y2:
			x1, x2 = (x1+1)%5, (x2+1)%5
		default:
			x2, x1 = x1, x2
		}
		return x1, y1, x2, y2
	})
}

func (c *Playfair) Decrypt(text []byte) []byte {
	return c.crypt(text, func(x1, y1, x2, y2 int) (int, int, int, int) {
		switch {
		case x1 == x2:
			y1, y2 = (y1+4)%5, (y2+4)%5
		case y1 == y2:
			x1, x2 = (x1+4)%5, (x2+4)%5
		default:
			x2, x1 = x1, x2
		}
		return x1, y1, x2, y2
	})
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
