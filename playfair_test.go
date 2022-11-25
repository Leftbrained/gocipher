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

	cipher := c.Encrypt([]byte("THEQUICKBROWNFOXJUMPSOVERTHELAZYDOGS"))

	if !bytes.Equal(cipher, []byte("PBIMXDTDGTAUWNHUDXQRLBRMYCGINOWTLDBM")) {
		t.Fatalf("invalid encryption: %s", cipher)
	}
}

func TestPlayfairBasicDecrypt(t *testing.T) {
	c, err := NewPlayfair([]byte("CRYPTOGRAPHY"))
	if err != nil {
		t.Fatalf("unexpected: could not instantiate")
	}

	plain := c.Decrypt([]byte("PBIMXDTDGTAUWNHUDXQRLBRMYCGINOWTLDBM"))

	if !bytes.Equal(plain, []byte("THEQUICKBROWNFOXIUMPSOVERTHELAZYDOGS")) {
		t.Fatalf("invalid decryption: %s", string(plain))
	}
}

func TestPlayfairDoubleLetterEncrypt(t *testing.T) {
	c, err := NewPlayfair([]byte("CRYPTOGRAPHY"))
	if err != nil {
		t.Fatalf("unexpected: could not instantiate")
	}

	cipher := c.Encrypt([]byte("OOF"))

	if !bytes.Equal(cipher, []byte("HUAD")) {
		t.Fatalf("invalid encryption: %s", cipher)
	}
}

func TestPlayfairDoubleLetterDecrypt(t *testing.T) {
	c, err := NewPlayfair([]byte("CRYPTOGRAPHY"))
	if err != nil {
		t.Fatalf("unexpected: could not instantiate")
	}

	plain := c.Decrypt([]byte("HUAD"))

	if !bytes.Equal(plain, []byte("OXOF")) {
		t.Fatalf("invalid decryption: %s", string(plain))
	}
}

func TestPlayfairDoubleXEncrypt(t *testing.T) {
	c, err := NewPlayfair([]byte("CRYPTOGRAPHY"))
	if err != nil {
		t.Fatalf("unexpected: could not instantiate")
	}

	cipher := c.Encrypt([]byte("XXF"))

	if !bytes.Equal(cipher, []byte("PXWI")) {
		t.Fatalf("invalid encryption: %s", cipher)
	}
}

func TestPlayfairDoubleXDecrypt(t *testing.T) {
	c, err := NewPlayfair([]byte("CRYPTOGRAPHY"))
	if err != nil {
		t.Fatalf("unexpected: could not instantiate")
	}

	plain := c.Decrypt([]byte("PXWI"))

	if !bytes.Equal(plain, []byte("XQXF")) {
		t.Fatalf("invalid decryption: %s", string(plain))
	}
}

func TestPlayfairExtraLetterEncrypt(t *testing.T) {
	c, err := NewPlayfair([]byte("CRYPTOGRAPHY"))
	if err != nil {
		t.Fatalf("unexpected: could not instantiate")
	}

	cipher := c.Encrypt([]byte("D"))

	if !bytes.Equal(cipher, []byte("IU")) {
		t.Fatalf("invalid encryption: %s", cipher)
	}
}

func TestPlayfairExtraLetterDecrypt(t *testing.T) {
	c, err := NewPlayfair([]byte("CRYPTOGRAPHY"))
	if err != nil {
		t.Fatalf("unexpected: could not instantiate")
	}

	plain := c.Decrypt([]byte("IU"))

	if !bytes.Equal(plain, []byte("DX")) {
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
