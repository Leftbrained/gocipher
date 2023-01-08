package cipher

import (
	"fmt"

	"github.com/leftbrained/gocipher"
)

type Vigenere struct {
	keyLen  int
	ciphers []gocipher.CharacterCipher
}

type VigenereConfig struct {
	newSubstitution func(key []byte, opts ...SubstitutionOption) (gocipher.CharacterCipher, error)
}

type VigenereOption func(*VigenereConfig)

var vignereAlphabet = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZ")

func NewVigenere(key []byte, opts ...VigenereOption) (*Vigenere, error) {
	size := len(key)
	c := Vigenere{
		keyLen:  size,
		ciphers: make([]gocipher.CharacterCipher, size),
	}

	cfg := &VigenereConfig{
		newSubstitution: func(key []byte, opts ...SubstitutionOption) (gocipher.CharacterCipher, error) {
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

func VigenereWithNewSubstitution(newSubstitution func(key []byte, opts ...SubstitutionOption) (gocipher.CharacterCipher, error)) VigenereOption {
	return func(cfg *VigenereConfig) {
		cfg.newSubstitution = newSubstitution
	}
}

func (c *Vigenere) Encrypt(text []byte) []byte {
	for i := 0; i < len(text); i++ {
		text[i] = c.ciphers[i%c.keyLen].EncryptByte(text[i])
	}
	return text
}

func (c *Vigenere) Decrypt(text []byte) []byte {
	for i := 0; i < len(text); i++ {
		text[i] = c.ciphers[i%c.keyLen].DecryptByte(text[i])
	}
	return text
}

func (c *Vigenere) EncryptByte(from byte) byte {
	return c.ciphers[0].EncryptByte(from)
}

func (c *Vigenere) DecryptByte(from byte) byte {
	return c.ciphers[0].DecryptByte(from)
}
