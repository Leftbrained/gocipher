package cipher

import (
	"fmt"

	"github.com/leftbrained/gocipher"
)

type Adfgx struct {
	polybius      gocipher.Cipher
	transposition gocipher.Cipher
}

type AdfgxConfig struct {
	newPolybius      func(key []byte, opts ...PolybiusOption) (gocipher.Cipher, error)
	newTransposition func(key []byte, opts ...TranspositionOption) (gocipher.Cipher, error)
}

type AdfgxOption func(*AdfgxConfig)

func NewAdfgx(key []byte, opts ...AdfgxOption) (*Adfgx, error) {
	cfg := &AdfgxConfig{
		newPolybius: func(key []byte, opts ...PolybiusOption) (gocipher.Cipher, error) {
			return NewPolybius(key, opts...)
		},
		newTransposition: func(key []byte, opts ...TranspositionOption) (gocipher.Cipher, error) {
			return NewTransposition(key, opts...)
		},
	}

	for _, opt := range opts {
		opt(cfg)
	}

	polybius, err := cfg.newPolybius(
		[]byte{},
		PolybiusWithAlphabet([]byte("ABCDEFGHIKLMNOPQRSTUVWXYZ"), map[byte]byte{
			'J': 'I',
		}),
		PolybiusWithCoords([]byte("ADFGX")),
	)
	if err != nil {
		return nil, fmt.Errorf("polybius: %s", err.Error())
	}

	transposition, err := cfg.newTransposition(key)
	if err != nil {
		return nil, fmt.Errorf("transposition: %s", err.Error())
	}

	c := Adfgx{
		polybius:      polybius,
		transposition: transposition,
	}

	return &c, nil
}

func AdfgxWithNewPolybius(newPolybius func(key []byte, opts ...PolybiusOption) (gocipher.Cipher, error)) AdfgxOption {
	return func(cfg *AdfgxConfig) {
		cfg.newPolybius = newPolybius
	}
}

func AdfgxWithNewTransposition(newTransposition func(key []byte, opts ...TranspositionOption) (gocipher.Cipher, error)) AdfgxOption {
	return func(cfg *AdfgxConfig) {
		cfg.newTransposition = newTransposition
	}
}

func (c *Adfgx) Encrypt(text []byte) []byte {
	return c.transposition.Encrypt(
		c.polybius.Encrypt(text),
	)
}

func (c *Adfgx) Decrypt(text []byte) []byte {
	return c.polybius.Decrypt(
		c.transposition.Decrypt(text),
	)
}
