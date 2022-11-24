package gocipher

import (
	"fmt"
)

type MonoalphabeticSubstitution struct {
	encryptMapping map[byte]byte
	decryptMapping map[byte]byte
}

func NewMonoalphabeticSubstitution(plainAlphabet, cipherAlphabet []byte) (*MonoalphabeticSubstitution, error) {

	size := len(plainAlphabet)

	if len(cipherAlphabet) != size {
		return nil, fmt.Errorf("size mismatch between plain and cipher alphabets")
	}

	c := MonoalphabeticSubstitution{
		encryptMapping: make(map[byte]byte, size),
		decryptMapping: make(map[byte]byte, size),
	}

	for i, plain := range plainAlphabet {
		cipher := cipherAlphabet[i]
		c.encryptMapping[plain] = cipher
		c.decryptMapping[cipher] = plain
	}

	return &c, nil
}

func (c *MonoalphabeticSubstitution) crypt(mapping map[byte]byte, text []byte) []byte {
	for i, from := range text {
		if to, ok := mapping[from]; ok {
			text[i] = to
		}
	}

	return text
}

func (c *MonoalphabeticSubstitution) Encrypt(plaintext []byte) []byte {
	return c.crypt(c.encryptMapping, plaintext)
}

func (c *MonoalphabeticSubstitution) Decrypt(ciphertext []byte) []byte {
	return c.crypt(c.decryptMapping, ciphertext)
}
