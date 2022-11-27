package cipher

import (
	"fmt"
	"testing"

	"github.com/leftbrained/gocipher"
)

func TestAdfgvxNew(t *testing.T) {
	c, err := NewAdfgvx(
		[]byte("ABYZ"),
	)

	if c == nil || err != nil {
		t.Fatalf(`could not instantiate`)
	}
}

func TestAdfgvxNewErrorAlphabetSize(t *testing.T) {
	c, err := NewAdfgvx(
		[]byte("ABYZ"),
		AdfgvxWithNewPolybius(func(key []byte, opts ...PolybiusOption) (gocipher.Cipher, error) {
			opts = append(opts, PolybiusWithAlphabet([]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ012345678"), map[byte]byte{
				'J': 'I',
			}))
			return NewPolybius(key, opts...)
		}),
	)

	if c != nil || err == nil {
		t.Fatalf("did not fail")
	}
}

func TestAdfgvxNewErrorInvalidPolybiusCipher(t *testing.T) {
	c, err := NewAdfgvx(
		[]byte("ABYZ"),
		AdfgvxWithNewPolybius(func(alphabet []byte, opts ...PolybiusOption) (gocipher.Cipher, error) {
			return nil, fmt.Errorf("random failure")
		}),
	)

	if c != nil || err == nil {
		t.Fatalf("did not fail")
	}
}

func TestAdfgvxNewErrorInvalidTranspositionCipher(t *testing.T) {
	c, err := NewAdfgvx(
		[]byte("ABYZ"),
		AdfgvxWithNewTransposition(func(key []byte, opts ...TranspositionOption) (gocipher.Cipher, error) {
			return nil, fmt.Errorf("random failure")
		}),
	)

	if c != nil || err == nil {
		t.Fatalf("did not fail")
	}
}

func TestAdfgvxBasicCrypt(t *testing.T) {
	c, err := NewAdfgvx(
		[]byte("CRYPTOGRAPHY"),
	)
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
	c, _ := NewAdfgvx(
		[]byte("CRYPTOGRAPHY"),
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = c.Encrypt([]byte("INCRYPTOGRAPHYASUBSTITUTIONCIPHERISAMETHODOFENCRYPTINGINWHICHUNITSOFPLAINTEXTAREREPLACEDWITHTHECIPHERTEXTINADEFINEDMANNERWITHTHEHELPOFAKEY"))
	}
}

func BenchmarkAdfgvxDecrypt(b *testing.B) {
	c, _ := NewAdfgvx(
		[]byte("CRYPTOGRAPHY"),
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = c.Decrypt([]byte("VAGGDGFAFDGAGFGAFFFFDDADGDGDDFFVDDFFFAGDAAAFDFFFGGADDFDDDDGAADADDFGADFFAGFAAFDAGDADDAGAAAGFVFAAFDXDVDVDGVXVDDDXAFDADFADFVVXGDFXDVFDGXVAVVXAAFFFAFFDFDADGVVXDDDDXVFDDAFDAFAFDFDXADFVGGXDFXXADFFDDFDFXXVGDVFFDDVVADADFFGAGGFFAFAGDGAADDAFFVGFAAAFFGAGAADFGAFGAAGGDDGAGXAFAFAXFFDAVVDGA"))
	}
}
