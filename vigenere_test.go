package gocipher

import (
	"bytes"
	"testing"
)

func TestVigenereNew(t *testing.T) {

	c, err := NewVigenere([]byte("ABYZ"))

	if c == nil || err != nil {
		t.Fatalf(`could not instantiate`)
	}
}

func TestVigenereNewErrorLowercase(t *testing.T) {

	c, err := NewVigenere([]byte("AbYZ"))

	if c != nil || err == nil {
		t.Fatalf("did not fail")
	}
}

func TestVigenereNewErrorLowerBound(t *testing.T) {

	c, err := NewVigenere([]byte("A@YZ"))

	if c != nil || err == nil {
		t.Fatalf("did not fail")
	}
}

func TestVigenereNewErrorUpperBound(t *testing.T) {

	c, err := NewVigenere([]byte("A[YZ"))

	if c != nil || err == nil {
		t.Fatalf("did not fail")
	}
}

func TestVigenereEncrypt(t *testing.T) {
	c, err := NewVigenere([]byte("VIG"))
	if err != nil {
		t.Fatalf("unexpected: could not instantiate")
	}

	cipher := c.Encrypt([]byte("CRYPTOGRAPHY"))

	if !bytes.Equal(cipher, []byte("XZEKBUBZGKPE")) {
		t.Fatalf("invalid encryption")
	}
}
