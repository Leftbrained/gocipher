package cipher

import (
	"fmt"
)

type Substitution struct {
	encrypt map[byte]byte
	decrypt map[byte]byte
}

type SubstitutionConfig struct {
	plainAlphabet  []byte
	cipherAlphabet []byte
}

type SubstitutionOption func(*SubstitutionConfig)

func NewSubstitution(key []byte, opts ...SubstitutionOption) (*Substitution, error) {
	cfg := &SubstitutionConfig{
		plainAlphabet: []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ"),
	}

	for _, opt := range opts {
		opt(cfg)
	}

	if len(cfg.cipherAlphabet) == 0 {
		cfg.cipherAlphabet = cfg.plainAlphabet
	}
	cfg.cipherAlphabet = GetKeyedAlphabet(key, cfg.cipherAlphabet)

	size := len(cfg.plainAlphabet)

	if len(cfg.cipherAlphabet) != size {
		return nil, fmt.Errorf("size mismatch between plain and cipher alphabets, found: plain=%d, cipher=%d", size, len(cfg.cipherAlphabet))
	}

	c := Substitution{
		encrypt: make(map[byte]byte, size),
		decrypt: make(map[byte]byte, size),
	}

	for i, plain := range cfg.plainAlphabet {
		cipher := cfg.cipherAlphabet[i]

		if _, ok := c.encrypt[plain]; ok {
			return nil, fmt.Errorf("plain alphabet has duplicates")
		}

		c.encrypt[plain] = cipher
		c.decrypt[cipher] = plain
	}

	return &c, nil
}

func SubstitutionWithPlainAlphabet(plainAlphabet []byte) SubstitutionOption {
	return func(cfg *SubstitutionConfig) {
		cfg.plainAlphabet = plainAlphabet
	}
}

func SubstitutionWithCipherAlphabet(cipherAlphabet []byte) SubstitutionOption {
	return func(cfg *SubstitutionConfig) {
		cfg.cipherAlphabet = cipherAlphabet
	}
}

func (c *Substitution) Encrypt(text []byte) []byte {
	for i, from := range text {
		if to, ok := c.encrypt[from]; ok {
			text[i] = to
		}
	}

	return text
}

func (c *Substitution) Decrypt(text []byte) []byte {
	for i, from := range text {
		if to, ok := c.decrypt[from]; ok {
			text[i] = to
		}
	}

	return text
}
