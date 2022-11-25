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

func TestVigenereBasicEncrypt(t *testing.T) {
	c, err := NewVigenere([]byte("CRYPTOGRAPHY"))
	if err != nil {
		t.Fatalf("unexpected: could not instantiate")
	}

	cipher := c.Encrypt([]byte("THEQUICKBROWNFOXJUMPSOVERTHELAZYDOG"))

	if !bytes.Equal(cipher, []byte("VYCFNWIBBGVUPWMMCISGSDCCTKFTEOFPDDN")) {
		t.Fatalf("invalid encryption: %s", cipher)
	}
}

func TestVigenereBasicDecrypt(t *testing.T) {
	c, err := NewVigenere([]byte("CRYPTOGRAPHY"))
	if err != nil {
		t.Fatalf("unexpected: could not instantiate")
	}

	plain := c.Decrypt([]byte("VYCFNWIBBGVUPWMMCISGSDCCTKFTEOFPDDN"))

	if !bytes.Equal(plain, []byte("THEQUICKBROWNFOXJUMPSOVERTHELAZYDOG")) {
		t.Fatalf("invalid decryption: %s", string(plain))
	}
}

func BenchmarkVigenereEncrypt(b *testing.B) {
	c, _ := NewVigenere([]byte("CRYPTOGRAPHY"))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = c.Encrypt([]byte("INCRYPTOGRAPHYASUBSTITUTIONCIPHERISAMETHODOFENCRYPTINGINWHICHUNITSOFPLAINTEXTAREREPLACEDWITHTHECIPHERTEXTINADEFINEDMANNERWITHTHEHELPOFAKEY"))
	}
}

func BenchmarkVigenereDecrypt(b *testing.B) {
	c, _ := NewVigenere([]byte("CRYPTOGRAPHY"))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = c.Decrypt([]byte("KEAGRDZFGGHNJPYHNPYKIIBRKFLRBDNVRXZYOVRWHRUWECJPAGRXGUOEWWPAJLLXMGUWPAHGPKCMMOXVRTWJCTCSPWZYTWLAKGFTKHKOTXUYFVDXGSJDACUCTNGIAHNVHTSNQWYZXM"))
	}
}
