package cipher

import (
	"bytes"
	"fmt"
	"math"
)

type Polybius struct {
	encrypt map[byte][2]byte
	decrypt map[[2]byte]byte
}

type PolybiusConfig struct {
	alphabet []byte
	coords   []byte
}

type PolybiusOption func(*PolybiusConfig)

func NewPolybius(key []byte, opts ...PolybiusOption) (*Polybius, error) {
	cfg := &PolybiusConfig{
		alphabet: []byte("ABCDEFGHIKLMNOPQRSTUVWXYZ"),
		coords:   []byte("123456789"),
	}

	for _, opt := range opts {
		opt(cfg)
	}

	cfg.alphabet = GetKeyedAlphabet(key, cfg.alphabet)

	alphabetSize := len(cfg.alphabet)
	size := int(math.Sqrt(float64(alphabetSize)))

	if size*size != alphabetSize {
		return nil, fmt.Errorf("expecting alphabet size to be a perfect square, found: %d", alphabetSize)
	}

	c := Polybius{
		encrypt: make(map[byte][2]byte, alphabetSize),
		decrypt: make(map[[2]byte]byte, alphabetSize),
	}

	if size > len(cfg.coords) {
		return nil, fmt.Errorf("expecting coords size to be at least %d, found: %d", size, len(cfg.coords))
	}

	for i, plain := range cfg.alphabet {
		x, y := i/size, i%size

		cipher := [2]byte{cfg.coords[x], cfg.coords[y]}

		c.encrypt[plain] = cipher
		c.decrypt[cipher] = plain
	}

	return &c, nil
}

func PolybiusWithAlphabet(alphabet []byte) PolybiusOption {
	return func(cfg *PolybiusConfig) {
		cfg.alphabet = alphabet
	}
}

func PolybiusWithCoords(coords []byte) PolybiusOption {
	return func(cfg *PolybiusConfig) {
		cfg.coords = coords
	}
}

func (c *Polybius) Encrypt(text []byte) []byte {
	buffer := bytes.NewBuffer(make([]byte, 0, len(text)*2))

	for _, from := range text {
		if to, ok := c.encrypt[from]; ok {
			buffer.Write(to[:])
		}
	}

	return buffer.Bytes()
}

func (c *Polybius) Decrypt(text []byte) []byte {
	length := len(text)
	if length%2 > 0 {
		length--
	}

	buffer := bytes.NewBuffer(make([]byte, 0, length/2))

	for i := 0; i < length; i += 2 {
		to := [2]byte{text[i], text[i+1]}
		if from, ok := c.decrypt[to]; ok {
			buffer.WriteByte(from)
		}
	}

	return buffer.Bytes()
}
