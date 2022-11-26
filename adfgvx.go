package gocipher

import "fmt"

type Adfgvx struct {
	polybius      Cipher
	transposition Cipher
}

func NewAdfgvx(alphabet, key []byte) (*Adfgvx, error) {
	alphabetSize := len(alphabet)
	if alphabetSize != 36 {
		return nil, fmt.Errorf("expecting alphabet size to be 25, found: %d", alphabetSize)
	}

	polybius, err := NewPolybius(alphabet, PolybiusWithCoords([]byte("ADFGVX")))
	if err != nil {
		return nil, fmt.Errorf("polybius: %s", err.Error())
	}

	transposition, err := NewTransposition(key)
	if err != nil {
		return nil, fmt.Errorf("transposition: %s", err.Error())
	}

	c := Adfgvx{
		polybius:      polybius,
		transposition: transposition,
	}

	return &c, nil
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
