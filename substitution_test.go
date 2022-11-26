package gocipher

import (
	"testing"
)

func TestSubstitutionNew(t *testing.T) {
	c, err := NewSubstitution(
		[]byte("CRYPTOGRAPHY"),
	)

	if c == nil || err != nil {
		t.Fatalf(`could not instantiate: "%s"`, err.Error())
	}
}

func TestSubstitutionBasicCrypt(t *testing.T) {
	c, err := NewSubstitution(
		[]byte("CRYPTOGRAPHY"),
	)
	if err != nil {
		t.Fatalf(`unexpected: could not instantiate: "%s"`, err.Error())
	}

	testCipherCrypt(c, t,
		[]byte("THEQUICKBROWNFOXJUMPSOVERTHELAZYDOG"),
		[]byte("QATLSHYDRMJVIOJWBSFKNJUTMQATECZXPJG"),
		[]byte("THEQUICKBROWNFOXJUMPSOVERTHELAZYDOG"),
	)
}

func TestSubstitutionNewWithPlainAlphabet(t *testing.T) {
	c, err := NewSubstitution(
		[]byte("CRYPTOGRAPHY"),
		SubstitutionWithPlainAlphabet([]byte("ABCD")),
	)

	if c == nil || err != nil {
		t.Fatalf(`could not instantiate: "%s"`, err.Error())
	}
}

func TestSubstitutionNewWithCipherAlphabet(t *testing.T) {
	c, err := NewSubstitution(
		[]byte("CRYPTOGRAPHY"),
		SubstitutionWithCipherAlphabet([]byte("ZYXWVUTSRQPONMLKJIHGFEDCBA")),
	)

	if c == nil || err != nil {
		t.Fatalf(`could not instantiate: "%s"`, err.Error())
	}
}

func TestSubstitutionNewWithAlphabets(t *testing.T) {
	c, err := NewSubstitution(
		[]byte("CRYPTOGRAPHY"),
		SubstitutionWithPlainAlphabet([]byte("ABCD")),
		SubstitutionWithCipherAlphabet([]byte("WXYZ")),
	)

	if c == nil || err != nil {
		t.Fatalf(`could not instantiate: "%s"`, err.Error())
	}
}

func TestSubstitutionNewErrorPlainBigger(t *testing.T) {
	c, err := NewSubstitution(
		[]byte("CRYPTOGRAPHY"),
		SubstitutionWithPlainAlphabet([]byte("ABCD")),
		SubstitutionWithCipherAlphabet([]byte("WXY")),
	)

	if c != nil || err == nil {
		t.Fatalf("did not fail")
	}
}

func TestSubstitutionNewErrorCipherBigger(t *testing.T) {
	c, err := NewSubstitution(
		[]byte("CRYPTOGRAPHY"),
		SubstitutionWithPlainAlphabet([]byte("ABC")),
		SubstitutionWithCipherAlphabet([]byte("WXYZ")),
	)

	if c != nil || err == nil {
		t.Fatalf("did not fail")
	}
}

func TestSubstitutionNewErrorPlainDuplicate(t *testing.T) {
	c, err := NewSubstitution(
		[]byte("CRYPTOGRAPHY"),
		SubstitutionWithPlainAlphabet([]byte("ABBD")),
		SubstitutionWithCipherAlphabet([]byte("WXYZ")),
	)

	if c != nil || err == nil {
		t.Fatalf("did not fail")
	}
}

func BenchmarkSubstitutionEncrypt(b *testing.B) {
	c, _ := NewSubstitution([]byte("CRYPTOGRAPHY"))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = c.Encrypt([]byte("INCRYPTOGRAPHYASUBSTITUTIONCIPHERISAMETHODOFENCRYPTINGINWHICHUNITSOFPLAINTEXTAREREPLACEDWITHTHECIPHERTEXTINADEFINEDMANNERWITHTHEHELPOFAKEY"))
	}
}

func BenchmarkSubstitutionDecrypt(b *testing.B) {
	c, _ := NewSubstitution([]byte("CRYPTOGRAPHY"))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = c.Decrypt([]byte("HIYMXKQJGMCKAXCNSRNQHQSQHJIYHKATMHNCFTQAJPJOTIYMXKQHIGHIVAHYASIHQNJOKECHIQTWQCMTMTKECYTPVHQAQATYHKATMQTWQHICPTOHITPFCIITMVHQAQATATEKJOCDTX"))
	}
}
