package gocipher

import (
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

func TestSubstitutionBasicCrypt(t *testing.T) {
	c, err := NewSubstitution([]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ"), []byte("CRYPTOGAHBDEFIJKLMNQSUVWXZ"))
	if err != nil {
		t.Fatalf("unexpected: could not instantiate")
	}

	testCipherCrypt(c, t,
		[]byte("THEQUICKBROWNFOXJUMPSOVERTHELAZYDOG"),
		[]byte("QATLSHYDRMJVIOJWBSFKNJUTMQATECZXPJG"),
		[]byte("THEQUICKBROWNFOXJUMPSOVERTHELAZYDOG"),
	)
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
