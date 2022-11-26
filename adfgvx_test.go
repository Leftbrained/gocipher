package gocipher

import (
	"testing"
)

func TestAdfgvxNew(t *testing.T) {
	c, err := NewAdfgvx([]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"), []byte("ABYZ"))

	if c == nil || err != nil {
		t.Fatalf(`could not instantiate`)
	}
}

func TestAdfgvxNewErrorAlphabetSize(t *testing.T) {
	c, err := NewAdfgvx([]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ012345678"), []byte("ABYZ"))

	if c != nil || err == nil {
		t.Fatalf("did not fail")
	}
}

func TestAdfgvxBasicCrypt(t *testing.T) {
	c, err := NewAdfgvx([]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"), []byte("CRYPTOGRAPHY"))
	if err != nil {
		t.Fatalf("unexpected: could not instantiate")
	}

	testCipherCrypt(c, t,
		[]byte("THEQUICKBROWNFOXJUMPSOVERTHELAZYDOG"),
		[]byte("GFDGDDGAFFFVFFGFAFDGGAAVDFADGDVXGDAFFGGXADFDAXDVXXFVFAAFGDADDAFGVFVFVA"),
		[]byte("THEQUICKBROWNFOXJUMPSOVERTHELAZYDOG"),
	)
}

func BenchmarkAdfgvxEncrypt(b *testing.B) {
	c, _ := NewAdfgvx([]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"), []byte("CRYPTOGRAPHY"))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = c.Encrypt([]byte("INCRYPTOGRAPHYASUBSTITUTIONCIPHERISAMETHODOFENCRYPTINGINWHICHUNITSOFPLAINTEXTAREREPLACEDWITHTHECIPHERTEXTINADEFINEDMANNERWITHTHEHELPOFAKEY"))
	}
}

func BenchmarkAdfgvxDecrypt(b *testing.B) {
	c, _ := NewAdfgvx([]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"), []byte("CRYPTOGRAPHY"))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = c.Decrypt([]byte("VAGGDGFAFDGAGFGAFFFFDDADGDGDDFFVDDFFFAGDAAAFDFFFGGADDFDDDDGAADADDFGADFFAGFAAFDAGDADDAGAAAGFVFAAFDXDVDVDGVXVDDDXAFDADFADFVVXGDFXDVFDGXVAVVXAAFFFAFFDFDADGVVXDDDDXVFDDAFDAFAFDFDXADFVGGXDFXXADFFDDFDFXXVGDVFFDDVVADADFFGAGGFFAFAGDGAADDAFFVGFAAAFFGAGAADFGAFGAAGGDDGAGXAFAFAXFFDAVVDGA"))
	}
}
