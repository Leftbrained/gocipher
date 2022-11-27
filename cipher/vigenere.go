package cipher

import (
	"fmt"

	"github.com/leftbrained/gocipher"
)

type Vigenere struct {
	keyLen  int
	ciphers []gocipher.Cipher
}

type VigenereConfig struct {
	newSubstitution func(key []byte, opts ...SubstitutionOption) (gocipher.Cipher, error)
}

type VigenereOption func(*VigenereConfig)

var vignereAlphabet = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZ")

func NewVigenere(key []byte, opts ...VigenereOption) (*Vigenere, error) {
	size := len(key)
	c := Vigenere{
		keyLen:  size,
		ciphers: make([]gocipher.Cipher, size),
	}

	cfg := &VigenereConfig{
		newSubstitution: func(key []byte, opts ...SubstitutionOption) (gocipher.Cipher, error) {
			return NewSubstitution(key, opts...)
		},
	}

	for _, opt := range opts {
		opt(cfg)
	}

	for i, k := range key {
		if k < 65 || k > 90 {
			return nil, fmt.Errorf("invalid character in key: %s", string(k))
		}

		k -= 65

		cipher, err := cfg.newSubstitution(
			[]byte{},
			SubstitutionWithPlainAlphabet(vignereAlphabet[:26]),
			SubstitutionWithCipherAlphabet(vignereAlphabet[k:k+26]),
		)
		if err != nil {
			return nil, err
		}
		c.ciphers[i] = cipher
	}

	return &c, nil
}

func VigenereWithNewSubstitution(newSubstitution func(key []byte, opts ...SubstitutionOption) (gocipher.Cipher, error)) VigenereOption {
	return func(cfg *VigenereConfig) {
		cfg.newSubstitution = newSubstitution
	}
}

func (c *Vigenere) Encrypt(text []byte) []byte {
	for i := 0; i < len(text); i++ {
		text[i] = c.ciphers[i%c.keyLen].Encrypt(text[i : i+1])[0]
	}
	return text
}

func (c *Vigenere) Decrypt(text []byte) []byte {
	for i := 0; i < len(text); i++ {
		text[i] = c.ciphers[i%c.keyLen].Decrypt(text[i : i+1])[0]
	}
	return text
}