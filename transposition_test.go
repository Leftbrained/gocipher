package gocipher

import (
	"testing"
)

func TestTranspositionNew(t *testing.T) {
	c, err := NewTransposition([]byte("ABCD"))

	if c == nil || err != nil {
		t.Fatalf(`could not instantiate`)
	}
}

func TestTranspositionBasicCrypt(t *testing.T) {
	c, err := NewTransposition([]byte("CRYPTOGRAPHY"))
	if err != nil {
		t.Fatalf("unexpected: could not instantiate")
	}

	testCipherCrypt(c, t,
		[]byte("THEQUICKBROWNFOXJUMPSOVERTHELAZYDOG"),
		[]byte("BSDTNRCMZOVGIUAQXEROOHFTKPYUJLEOHWE"),
		[]byte("THEQUICKBROWNFOXJUMPSOVERTHELAZYDOG"),
	)
}

func BenchmarkTranspositionEncrypt(b *testing.B) {
	c, _ := NewTransposition([]byte("CRYPTOGRAPHY"))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = c.Encrypt([]byte("INCRYPTOGRAPHYASUBSTITUTIONCIPHERISAMETHODOFENCRYPTINGINWHICHUNITSOFPLAINTEXTAREREPLACEDWITHTHECIPHERTEXTINADEFINEDMANNERWITHTHEHELPOFAKEY"))
	}
}

func BenchmarkTranspositionDecrypt(b *testing.B) {
	c, _ := NewTransposition([]byte("CRYPTOGRAPHY"))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = c.Decrypt([]byte("PRMHXNTTILSEHEAITCIHNAAHYPNCIROECIMEALNHNYNCRNELSAEUAPCIYTTPUCRNDTIPPEHOBEWIEIPHTXOFSDITFNHEOTDATWTSIGERGAOKIIHONIRTRRFETTTFENIEAEEYUHHPDW"))
	}
}
