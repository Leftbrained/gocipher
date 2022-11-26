package gocipher

import "fmt"

type Adfgx struct {
	polybius      Cipher
	transposition Cipher
}

func NewAdfgx(alphabet, key []byte) (*Adfgx, error) {
	alphabetSize := len(alphabet)
	if alphabetSize != 25 {
		return nil, fmt.Errorf("expecting alphabet size to be 25, found: %d", alphabetSize)
	}

	polybius, err := NewPolybius(alphabet, PolybiusWithCoords([]byte("ADFGX")))
	if err != nil {
		return nil, fmt.Errorf("polybius: %s", err.Error())
	}

	transposition, err := NewTransposition(key)
	if err != nil {
		return nil, fmt.Errorf("transposition: %s", err.Error())
	}

	c := Adfgx{
		polybius:      polybius,
		transposition: transposition,
	}

	return &c, nil
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
