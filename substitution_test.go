package gocipher

import (
	"bytes"
	"testing"
)

func TestSubstitutionNew(t *testing.T) {
	c, err := NewSubstitution([]byte("ABCD"), []byte("WXYZ"))

	if c == nil || err != nil {
		t.Fatalf(`could not instantiate`)
	}
}

func TestSubstitutionNewErrorPlainBigger(t *testing.T) {
	c, err := NewSubstitution([]byte("ABCD"), []byte("WXY"))

	if c != nil || err == nil {
		t.Fatalf("did not fail")
	}
}

func TestSubstitutionNewErrorCipherBigger(t *testing.T) {
	c, err := NewSubstitution([]byte("ABC"), []byte("WXYZ"))

	if c != nil || err == nil {
		t.Fatalf("did not fail")
	}
}

func TestSubstitutionNewErrorPlainDuplicate(t *testing.T) {
	c, err := NewSubstitution([]byte("ABBD"), []byte("WXYZ"))

	if c != nil || err == nil {
		t.Fatalf("did not fail")
	}
}

func TestSubstitutionNewErrorCipherDuplicate(t *testing.T) {
	c, err := NewSubstitution([]byte("ABCD"), []byte("WXYX"))

	if c != nil || err == nil {
		t.Fatalf("did not fail")
	}
}

func TestSubstitutionBasicEncrypt(t *testing.T) {
	c, _ := NewSubstitution([]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ"), []byte("CRYPTOGAHBDEFIJKLMNQSUVWXZ"))

	cipher := c.Encrypt([]byte("THEQUICKBROWNFOXJUMPSOVERTHELAZYDOG"))

	if !bytes.Equal(cipher, []byte("QATLSHYDRMJVIOJWBSFKNJUTMQATECZXPJG")) {
		t.Fatalf("invalid encryption: %s", cipher)
	}
}

func TestSubstitutionBasicDecrypt(t *testing.T) {
	c, _ := NewSubstitution([]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ"), []byte("CRYPTOGAHBDEFIJKLMNQSUVWXZ"))

	plain := c.Decrypt([]byte("QATLSHYDRMJVIOJWBSFKNJUTMQATECZXPJG"))

	if !bytes.Equal(plain, []byte("THEQUICKBROWNFOXJUMPSOVERTHELAZYDOG")) {
		t.Fatalf("invalid decryption: %s", string(plain))
	}
}

func BenchmarkSubstitutionEncrypt(b *testing.B) {
	c, _ := NewSubstitution([]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ"), []byte("CRYPTOGAHBDEFIJKLMNQSUVWXZ"))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = c.Encrypt([]byte("INCRYPTOGRAPHYASUBSTITUTIONCIPHERISAMETHODOFENCRYPTINGINWHICHUNITSOFPLAINTEXTAREREPLACEDWITHTHECIPHERTEXTINADEFINEDMANNERWITHTHEHELPOFAKEY"))
	}
}

func BenchmarkSubstitutionDecrypt(b *testing.B) {
	c, _ := NewSubstitution([]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ"), []byte("CRYPTOGAHBDEFIJKLMNQSUVWXZ"))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = c.Decrypt([]byte("HIYMXKQJGMCKAXCNSRNQHQSQHJIYHKATMHNCFTQAJPJOTIYMXKQHIGHIVAHYASIHQNJOKECHIQTWQCMTMTKECYTPVHQAQATYHKATMQTWQHICPTOHITPFCIITMVHQAQATATEKJOCDTX"))
	}
}
