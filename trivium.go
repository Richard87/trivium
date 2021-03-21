package trivium

import (
	"fmt"
	"github.com/dropbox/godropbox/container/bitvector"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

type Trivium struct {
	iv    *bitvector.BitVector
	key   *bitvector.BitVector
	state *bitvector.BitVector
}

func NewTrivium(iv []byte, key []byte) *Trivium {
	state := make([]byte, 36)

	bv := bitvector.NewBitVector(state, 288)
	kBv := bitvector.NewBitVector(key, 80)
	ivBv := bitvector.NewBitVector(iv, 80)

	for i := 0; i < kBv.Length(); i++ {
		bv.Set(kBv.Element(i), i)
	}

	for i := 0; i < ivBv.Length(); i++ {
		bv.Set(kBv.Element(i), i+93)
	}
	bv.Set(1, 285)
	bv.Set(1, 286)
	bv.Set(1, 287)

	for i := 0; i < 288; i++ {

		t1 := bv.Element(65) ^ (bv.Element(90) ^ bv.Element(91) ^ bv.Element(92)) ^ bv.Element(170)
		t2 := bv.Element(161) ^ (bv.Element(174) ^ bv.Element(175) ^ bv.Element(177)) ^ bv.Element(263)
		t3 := bv.Element(242) ^ (bv.Element(85) ^ bv.Element(286) ^ bv.Element(287)) ^ bv.Element(65)

		bv.Insert(bv.Element(bv.Length()-1), 0)

		bv.Set(t3, 0)
		bv.Set(t1, 93)
		bv.Set(t2, 177)
	}

	t := &Trivium{
		iv:    kBv,
		key:   ivBv,
		state: bv,
	}
	return t
}

func (t *Trivium) Encrypt(content []byte) []byte {
	bv := bitvector.NewBitVector(content, 8*len(content))
	ciphertext := bitvector.NewBitVector(make([]byte, len(content)), 8*len(content))

	for i := 0; i < bv.Length(); i++ {

		t1 := bv.Element(65) ^ (bv.Element(90) ^ bv.Element(91) ^ bv.Element(92)) ^ bv.Element(170)
		t2 := bv.Element(161) ^ (bv.Element(174) ^ bv.Element(175) ^ bv.Element(177)) ^ bv.Element(263)
		t3 := bv.Element(242) ^ (bv.Element(85) ^ bv.Element(286) ^ bv.Element(287)) ^ bv.Element(65)

		bv.Insert(bv.Element(bv.Length()-1), 0)

		bv.Set(t3, 0)
		bv.Set(t1, 93)
		bv.Set(t2, 177)

		keybit := bv.Element(65) ^ bv.Element(92) ^ bv.Element(161) ^ bv.Element(177) ^ bv.Element(242) ^ bv.Element(287)
		result := keybit ^ bv.Element(i)

		ciphertext.Set(result, i)
	}

	return ciphertext.Bytes()
}
