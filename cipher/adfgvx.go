package cipher

import (
	"fmt"

	"github.com/leftbrained/gocipher"
)

type Adfgvx struct {
	polybius      gocipher.Cipher
	transposition gocipher.Cipher
}

type AdfgvxConfig struct {
	newPolybius      func(key []byte, opts ...PolybiusOption) (gocipher.Cipher, error)
	newTransposition func(key []byte, opts ...TranspositionOption) (gocipher.Cipher, error)
}

type AdfgvxOption func(*AdfgvxConfig)

func NewAdfgvx(key []byte, opts ...AdfgvxOption) (*Adfgvx, error) {
	cfg := &AdfgvxConfig{
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
		PolybiusWithAlphabet([]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")),
		PolybiusWithCoords([]byte("ADFGVX")),
	)
	if err != nil {
		return nil, fmt.Errorf("polybius: %s", err.Error())
	}

	transposition, err := cfg.newTransposition(key)
	if err != nil {
		return nil, fmt.Errorf("transposition: %s", err.Error())
	}

	c := Adfgvx{
		polybius:      polybius,
		transposition: transposition,
	}

	return &c, nil
}

func AdfgvxWithNewPolybius(newPolybius func(key []byte, opts ...PolybiusOption) (gocipher.Cipher, error)) AdfgvxOption {
	return func(cfg *AdfgvxConfig) {
		cfg.newPolybius = newPolybius
	}
}

func AdfgvxWithNewTransposition(newTransposition func(key []byte, opts ...TranspositionOption) (gocipher.Cipher, error)) AdfgvxOption {
	return func(cfg *AdfgvxConfig) {
		cfg.newTransposition = newTransposition
	}
}

func (c *Adfgvx) Encrypt(text []byte) []byte {
	return c.transposition.Encrypt(
		c.polybius.Encrypt(text),
	)
}

func (c *Adfgvx) Decrypt(text []byte) []byte {
	return c.polybius.Decrypt(
		c.transposition.Decrypt(text),
	)
}
