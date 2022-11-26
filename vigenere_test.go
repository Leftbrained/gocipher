package gocipher

import (
	"fmt"
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

func TestVigenereNewErrorInvalidSubstitutionCipher(t *testing.T) {
	c, err := NewVigenere([]byte("ABYZ"), VigenereWithNewSubstitution(func(plainAlphabet, cipherAlphabet []byte, opts ...SubstitutionOption) (*Substitution, error) {
		return nil, fmt.Errorf("random failure")
	}))

	if c != nil || err == nil {
		t.Fatalf("did not fail")
	}
}

func TestVigenereBasicCrypt(t *testing.T) {
	c, err := NewVigenere([]byte("CRYPTOGRAPHY"))
	if err != nil {
		t.Fatalf("unexpected: could not instantiate")
	}

	testCipherCrypt(c, t,
		[]byte("THEQUICKBROWNFOXJUMPSOVERTHELAZYDOG"),
		[]byte("VYCFNWIBBGVUPWMMCISGSDCCTKFTEOFPDDN"),
		[]byte("THEQUICKBROWNFOXJUMPSOVERTHELAZYDOG"),
	)
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
