package substitution

import (
	"bytes"
	"fmt"
	"testing"
)

func TestSimpleNew(t *testing.T) {
	plain, cipher := []byte("cBda"), []byte("aDcB")

	c, err := NewSimple(plain, cipher)

	if c == nil || err != nil {
		t.Fatalf(`could not instantiate`)
	}
}

func TestSimpleNewErrorPlainBigger(t *testing.T) {
	plain, cipher := []byte("cBda"), []byte("aDc")

	c, err := NewSimple(plain, cipher)

	if c != nil || err == nil {
		t.Fatalf("did not fail")
	}
}

func TestSimpleNewErrorCipherBigger(t *testing.T) {
	plain, cipher := []byte("cBda"), []byte("aDc")

	c, err := NewSimple(plain, cipher)

	if c != nil || err == nil {
		t.Fatalf("did not fail")
	}
}

func TestSimpleEncrypt(t *testing.T) {
	c, _ := NewSimple([]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"), []byte("BCDEFGHIJKLMNOPQRSTUVWXYZAbcdefghijklmnopqrstuvwxyza"))

	cipher := c.Encrypt([]byte("A Man, A Plan, A Canal - Panama!"))

	if !bytes.Equal(cipher, []byte("B Nbo, B Qmbo, B Dbobm - Qbobnb!")) {
		t.Fatalf("invalid encryption")
	}
}

func TestSimpleDecrypt(t *testing.T) {
	c, _ := NewSimple([]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"), []byte("BCDEFGHIJKLMNOPQRSTUVWXYZAbcdefghijklmnopqrstuvwxyza"))

	cipher := c.Decrypt([]byte("B Nbo, B Qmbo, B Dbobm - Qbobnb!"))

	if !bytes.Equal(cipher, []byte("A Man, A Plan, A Canal - Panama!")) {
		t.Fatalf("invalid encryption")
	}
}

var benchmarkArgs = []struct {
	plain  []byte
	cipher []byte
}{
	{
		plain:  []byte(""),
		cipher: []byte(""),
	},
	{
		plain:  []byte("AB"),
		cipher: []byte("BA"),
	},
	{
		plain:  []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"),
		cipher: []byte("9876543210ZYXWVUTSRQPONMLKJIHGFEDCBA"),
	},
	{
		plain:  []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"),
		cipher: []byte("IHGFEDCBA9876543210ZYXWVUTSRQPONMLKJIHGFEDCBA9876543210ZYXWVUTSRQPONMLKJIHGFEDCBA9876543210ZYXWVUTSRQPONMLKJIHGFEDCBA9876543210ZYXWVUTSRQPONMLKJIHGFEDCBA9876543210ZYXWVUTSRQPONMLKJ"),
	},
}

func BenchmarkSimpleNew(b *testing.B) {
	for _, a := range benchmarkArgs {
		b.Run(fmt.Sprintf("size=%d", len(a.plain)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = NewSimple(a.plain, a.cipher)
			}
		})
	}
}

var benchmarkPlaintext = []byte(`In cryptography, a substitution cipher is a method of encrypting in which units of plaintext are replaced with the ciphertext, in a defined manner, with the help of a key; the "units" may be single letters (the most common), pairs of letters, triplets of letters, mixtures of the above, and so forth. The receiver deciphers the text by performing the inverse substitution process to extract the original message.

Substitution ciphers can be compared with transposition ciphers. In a transposition cipher, the units of the plaintext are rearranged in a different and usually quite complex order, but the units themselves are left unchanged. By contrast, in a substitution cipher, the units of the plaintext are retained in the same sequence in the ciphertext, but the units themselves are altered.

There are a number of different types of substitution cipher. If the cipher operates on single letters, it is termed a simple substitution cipher; a cipher that operates on larger groups of letters is termed polygraphic. A monoalphabetic cipher uses fixed substitution over the entire message, whereas a polyalphabetic cipher uses a number of substitutions at different positions in the message, where a unit from the plaintext is mapped to one of several possibilities in the ciphertext and vice versa.
`)

func BenchmarkSimpleCrypt(b *testing.B) {
	c, _ := NewSimple([]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"), []byte("ZYXWVUTSRQPONMLKJIHGFEDCBAzyxwvutsrqponmlkjihgfedcba"))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = c.crypt(c.encryptMapping, benchmarkPlaintext)
	}
}
