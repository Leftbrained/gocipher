package gocipher

import "fmt"

type Adfgvx struct {
	polybius      Cipher
	transposition Cipher
}

type AdfgvxConfig struct {
	newPolybius      func(alphabet []byte, opts ...PolybiusOption) (*Polybius, error)
	newTransposition func(key []byte, opts ...TranspositionOption) (*Transposition, error)
}

type AdfgvxOption func(*AdfgvxConfig)

func NewAdfgvx(alphabet, key []byte, opts ...AdfgvxOption) (*Adfgvx, error) {
	cfg := &AdfgvxConfig{
		newPolybius:      NewPolybius,
		newTransposition: NewTransposition,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	alphabetSize := len(alphabet)
	if alphabetSize != 36 {
		return nil, fmt.Errorf("expecting alphabet size to be 36, found: %d", alphabetSize)
	}

	polybius, err := cfg.newPolybius(alphabet, PolybiusWithCoords([]byte("ADFGVX")))
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

func AdfgvxWithNewPolybius(newPolybius func(alphabet []byte, opts ...PolybiusOption) (*Polybius, error)) AdfgvxOption {
	return func(cfg *AdfgvxConfig) {
		cfg.newPolybius = newPolybius
	}
}

func AdfgvxWithNewTransposition(newTransposition func(key []byte, opts ...TranspositionOption) (*Transposition, error)) AdfgvxOption {
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
