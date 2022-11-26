package gocipher

import "fmt"

type Adfgx struct {
	polybius      Cipher
	transposition Cipher
}

type AdfgxConfig struct {
	newPolybius      func(alphabet []byte, opts ...PolybiusOption) (*Polybius, error)
	newTransposition func(key []byte) (*Transposition, error)
}

type AdfgxOption func(*AdfgxConfig)

func NewAdfgx(alphabet, key []byte, opts ...AdfgxOption) (*Adfgx, error) {
	cfg := &AdfgxConfig{
		newPolybius:      NewPolybius,
		newTransposition: NewTransposition,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	alphabetSize := len(alphabet)
	if alphabetSize != 25 {
		return nil, fmt.Errorf("expecting alphabet size to be 25, found: %d", alphabetSize)
	}

	polybius, err := cfg.newPolybius(alphabet, PolybiusWithCoords([]byte("ADFGX")))
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

func AdfgxWithNewPolybius(newPolybius func(alphabet []byte, opts ...PolybiusOption) (*Polybius, error)) AdfgxOption {
	return func(cfg *AdfgxConfig) {
		cfg.newPolybius = newPolybius
	}
}

func AdfgxWithNewTransposition(newTransposition func(key []byte) (*Transposition, error)) AdfgxOption {
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
