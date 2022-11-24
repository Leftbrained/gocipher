package substitution

import (
	"fmt"
)

type SimpleCipher struct {
	encryptMapping map[byte]byte
	decryptMapping map[byte]byte
}

func NewSimple(plainAlphabet, cipherAlphabet []byte) (*SimpleCipher, error) {

	size := len(plainAlphabet)

	if len(cipherAlphabet) != size {
		return nil, fmt.Errorf("size mismatch between plain and cipher alphabets")
	}

	c := SimpleCipher{
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

func (c *SimpleCipher) crypt(mapping map[byte]byte, text []byte) []byte {
	for i, from := range text {
		if to, ok := mapping[from]; ok {
			text[i] = to
		}
	}

	return text
}

func (c *SimpleCipher) Encrypt(plaintext []byte) []byte {
	return c.crypt(c.encryptMapping, plaintext)
}

func (c *SimpleCipher) Decrypt(ciphertext []byte) []byte {
	return c.crypt(c.decryptMapping, ciphertext)
}
