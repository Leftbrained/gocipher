package gocipher

import (
	"fmt"
)

type Substitution struct {
	encrypt map[byte]byte
	decrypt map[byte]byte
}

func NewSubstitution(plainAlphabet, cipherAlphabet []byte) (*Substitution, error) {

	size := len(plainAlphabet)

	if len(cipherAlphabet) != size {
		return nil, fmt.Errorf("size mismatch between plain and cipher alphabets")
	}

	if hasDuplicates(plainAlphabet) {
		return nil, fmt.Errorf("plain alphabet has duplicates")
	}

	if hasDuplicates(cipherAlphabet) {
		return nil, fmt.Errorf("cipher alphabet has duplicates")
	}

	c := Substitution{
		encrypt: make(map[byte]byte, size),
		decrypt: make(map[byte]byte, size),
	}

	for i, plain := range plainAlphabet {
		cipher := cipherAlphabet[i]
		c.encrypt[plain] = cipher
		c.decrypt[cipher] = plain
	}

	return &c, nil
}

func (c *Substitution) crypt(mapping map[byte]byte, text []byte) []byte {
	for i, from := range text {
		if to, ok := mapping[from]; ok {
			text[i] = to
		}
	}

	return text
}

func (c *Substitution) Encrypt(plaintext []byte) []byte {
	return c.crypt(c.encrypt, plaintext)
}

func (c *Substitution) Decrypt(ciphertext []byte) []byte {
	return c.crypt(c.decrypt, ciphertext)
}

func hasDuplicates(s []byte) bool {
	found := make(map[byte]bool, len(s))

	for _, b := range s {
		if _, ok := found[b]; ok {
			return true
		}
		found[b] = true
	}

	return false
}
