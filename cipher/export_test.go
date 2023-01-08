package cipher

import (
	"bytes"
	"testing"

	"github.com/leftbrained/gocipher"
)

func testCipherCrypt(c gocipher.Cipher, t *testing.T, input, expectedCipher, expectedPlain []byte) {

	cipher := c.Encrypt(input)

	if !bytes.Equal(cipher, expectedCipher) {
		t.Fatalf("invalid encryption: expected=%q actual=%q", expectedCipher, cipher)
	}

	plain := c.Decrypt(cipher)

	if !bytes.Equal(plain, expectedPlain) {
		t.Fatalf("invalid decryption: expected=%q actual=%q", expectedPlain, plain)
	}
}

func testCipherCryptByte(c gocipher.CharacterCipher, t *testing.T, input, expectedCipher, expectedPlain byte) {

	cipher := c.EncryptByte(input)

	if cipher != expectedCipher {
		t.Fatalf("invalid encryption: expected=%q actual=%q", expectedCipher, cipher)
	}

	plain := c.DecryptByte(cipher)

	if plain != expectedPlain {
		t.Fatalf("invalid decryption: expected=%q actual=%q", expectedPlain, plain)
	}
}
