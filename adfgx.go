package gocipher

import "fmt"

type Adfgx struct {
	polybius      Cipher
	transposition Cipher
}

type AdfgxConfig struct {
	newPolybius      func(key []byte, opts ...PolybiusOption) (Cipher, error)
	newTransposition func(key []byte, opts ...TranspositionOption) (Cipher, error)
}

type AdfgxOption func(*AdfgxConfig)

func NewAdfgx(key []byte, opts ...AdfgxOption) (*Adfgx, error) {
	cfg := &AdfgxConfig{
		newPolybius: func(key []byte, opts ...PolybiusOption) (Cipher, error) {
			return NewPolybius(key, opts...)
		},
		newTransposition: func(key []byte, opts ...TranspositionOption) (Cipher, error) {
			return NewTransposition(key, opts...)
		},
	}

	for _, opt := range opts {
		opt(cfg)
	}

	polybius, err := cfg.newPolybius(
		[]byte{},
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

func AdfgxWithNewPolybius(newPolybius func(key []byte, opts ...PolybiusOption) (Cipher, error)) AdfgxOption {
	return func(cfg *AdfgxConfig) {
		cfg.newPolybius = newPolybius
	}
}

func AdfgxWithNewTransposition(newTransposition func(key []byte, opts ...TranspositionOption) (Cipher, error)) AdfgxOption {
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
