package present

import (
	"encoding/binary"
)

// BlockSize is the PRESENT block size in bytes.
const BlockSize = 8

const numRounds = 31

// Substitution and permutation tables for PRESENT.
var (
	sBox    = []byte{0xC, 5, 6, 0xB, 9, 0, 0xA, 0xD, 3, 0xE, 0xF, 8, 4, 7, 1, 2}
	sBoxInv = []byte{5, 0xE, 0xF, 8, 0xC, 1, 2, 0xD, 0xB, 4, 6, 3, 0, 7, 9, 0xA}
	p       = []byte{
		0, 16, 32, 48, 1, 17, 33, 49, 2, 18, 34, 50, 3, 19, 35, 51, 4, 20, 36, 52, 5, 21, 37, 53, 6,
		22, 38, 54, 7, 23, 39, 55, 8, 24, 40, 56, 9, 25, 41, 57, 10, 26, 42, 58, 11, 27, 43, 59, 12,
		28, 44, 60, 13, 29, 45, 61, 14, 30, 46, 62, 15, 31, 47, 63,
	}
	pInv = []byte{
		0, 4, 8, 12, 16, 20, 24, 28, 32, 36, 40, 44, 48, 52, 56, 60, 1, 5, 9, 13, 17, 21, 25, 29, 33,
		37, 41, 45, 49, 53, 57, 61, 2, 6, 10, 14, 18, 22, 26, 30, 34, 38, 42, 46, 50, 54, 58, 62, 3, 7,
		11, 15, 19, 23, 27, 31, 35, 39, 43, 47, 51, 55, 59, 63,
	}
)

type block struct {
	Key key
}

func (b *block) BlockSize() int {
	return BlockSize
}

func (b *block) Encrypt(dst, src []byte) {
	if len(src) < BlockSize {
		panic("present: input not full block")
	}
	if len(dst) < BlockSize {
		panic("present: output not full block")
	}
	state := binary.BigEndian.Uint64(src)
	roundKeys := make(chan uint64, numRounds+1)
	go expandKey(b.Key, roundKeys)
	for i := 0; i < numRounds; i++ {
		roundKey := <-roundKeys
		state ^= roundKey
		state = sBoxLayer(state, sBox)
		state = pLayer(state, p)
	}
	roundKey := <-roundKeys
	state ^= roundKey
	binary.BigEndian.PutUint64(dst, state)
}

func (b *block) Decrypt(dst, src []byte) {
	if len(src) < BlockSize {
		panic("present: input not full block")
	}
	if len(dst) < BlockSize {
		panic("present: output not full block")
	}
	c := make(chan uint64)
	go expandKey(b.Key, c)
	var roundKeys []uint64
	for roundKey := range c {
		roundKeys = append(roundKeys, roundKey)
	}
	state := binary.BigEndian.Uint64(src)
	state ^= roundKeys[numRounds]
	for i := numRounds - 1; i >= 0; i-- {
		state = pLayer(state, pInv)
		state = sBoxLayer(state, sBoxInv)
		state ^= roundKeys[i]
	}
	binary.BigEndian.PutUint64(dst, state)
}

func sBoxLayer(state uint64, s []byte) (result uint64) {
	for i := 0; i < 16; i++ {
		shift := 4 * uint(i)
		var mask uint64 = 0xF << shift
		x := (state & mask) >> shift
		y := uint64(s[x])
		z := y << shift
		result |= z
	}
	return
}

func pLayer(state uint64, p []byte) (result uint64) {
	for i, pi := range p {
		var mask uint64 = 1 << uint(i)
		x := (state & mask) >> uint(i)
		y := x << pi
		result |= y
	}
	return
}
