package gocipher

import (
	"fmt"
)

type Substitution struct {
	encrypt map[byte]byte
	decrypt map[byte]byte
}

type SubstitutionConfig struct{}

type SubstitutionOption func(*SubstitutionConfig)

func NewSubstitution(plainAlphabet, cipherAlphabet []byte, opts ...SubstitutionOption) (*Substitution, error) {
	cfg := &SubstitutionConfig{}

	for _, opt := range opts {
		opt(cfg)
	}

	size := len(plainAlphabet)

	if len(cipherAlphabet) != size {
		return nil, fmt.Errorf("size mismatch between plain and cipher alphabets")
	}

	c := Substitution{
		encrypt: make(map[byte]byte, size),
		decrypt: make(map[byte]byte, size),
	}

	for i, plain := range plainAlphabet {
		cipher := cipherAlphabet[i]

		if _, ok := c.encrypt[plain]; ok {
			return nil, fmt.Errorf("plain alphabet has duplicates")
		}

		if _, ok := c.decrypt[cipher]; ok {
			return nil, fmt.Errorf("cipher alphabet has duplicates")
		}

		c.encrypt[plain] = cipher
		c.decrypt[cipher] = plain
	}

	return &c, nil
}

func (c *Substitution) Encrypt(text []byte) []byte {
	for i, from := range text {
		if to, ok := c.encrypt[from]; ok {
			text[i] = to
		}
	}

	return text
}

func (c *Substitution) Decrypt(text []byte) []byte {
	for i, from := range text {
		if to, ok := c.decrypt[from]; ok {
			text[i] = to
		}
	}

	return text
}
