package gocipher

import (
	"bytes"
	"testing"
)

func TestPlayfairNew(t *testing.T) {
	c, err := NewPlayfair([]byte("ABYZ"))

	if c == nil || err != nil {
		t.Fatalf(`could not instantiate`)
	}
}

func TestPlayfairNewErrorLowercase(t *testing.T) {
	c, err := NewPlayfair([]byte("AbYZ"))

	if c != nil || err == nil {
		t.Fatalf("did not fail")
	}
}

func TestPlayfairNewErrorLowerBound(t *testing.T) {
	c, err := NewPlayfair([]byte("A@YZ"))

	if c != nil || err == nil {
		t.Fatalf("did not fail")
	}
}

func TestPlayfairNewErrorUpperBound(t *testing.T) {
	c, err := NewPlayfair([]byte("A[YZ"))

	if c != nil || err == nil {
		t.Fatalf("did not fail")
	}
}

func TestPlayfairBasicEncrypt(t *testing.T) {
	c, err := NewPlayfair([]byte("CRYPTOGRAPHY"))
	if err != nil {
		t.Fatalf("unexpected: could not instantiate")
	}

	cipher := c.Encrypt([]byte("THEQUICKBROWNFOXJUMPSOVERTHELAZYDOG"))

	if !bytes.Equal(cipher, []byte("PBIMXDTDGTAUWNHUDXQRLBRMYCGINOWTLDHV")) {
		t.Fatalf("invalid encryption: %s", cipher)
	}
}

func TestPlayfairBasicDecrypt(t *testing.T) {
	c, err := NewPlayfair([]byte("CRYPTOGRAPHY"))
	if err != nil {
		t.Fatalf("unexpected: could not instantiate")
	}

	plain := c.Decrypt([]byte("PBIMXDTDGTAUWNHUDXQRLBRMYCGINOWTLDHV"))

	if !bytes.Equal(plain, []byte("THEQUICKBROWNFOXIUMPSOVERTHELAZYDOGX")) {
		t.Fatalf("invalid decryption: %s", string(plain))
	}
}

func BenchmarkPlayfairEncrypt(b *testing.B) {
	c, _ := NewPlayfair([]byte("CRYPTOGRAPHY"))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = c.Encrypt([]byte("INCRYPTOGRAPHYASUBSTITUTIONCIPHERISAMETHODOFENCRYPTINGINWHICHUNITSOFPLAINTEXTAREREPLACEDWITHTHECIPHERTEXTINADEFINEDMANNERWITHTHEHELPOFAKEY"))
	}
}

func BenchmarkPlayfairDecrypt(b *testing.B) {
	c, _ := NewPlayfair([]byte("CRYPTOGRAPHY"))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = c.Decrypt([]byte("FQRYPTCBEGHYAPBNZOZBKPZCDHLYQHGIPENBVMPBDLADFMRYPTPKMAFQXADPOXQFBZADCQHFSYIVYBGMGMCQOYFEXFPBPBDRQHGIYCIVPKWFEFIKMFELFWMFYVKPBPGIGIQCADBFFR"))
	}
}
