package gocipher

import (
	"bytes"
	"testing"
)

func testCipherCrypt(c Cipher, t *testing.T, inputPlain, expectedCipher, expectedPlain []byte) {

	cipher := c.Encrypt(inputPlain)

	if !bytes.Equal(cipher, expectedCipher) {
		t.Fatalf("invalid encryption: %s", cipher)
	}

	plain := c.Decrypt(cipher)

	if !bytes.Equal(plain, expectedPlain) {
		t.Fatalf("invalid decryption: %s", plain)
	}
}
