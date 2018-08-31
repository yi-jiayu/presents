package presents

import (
	"errors"
	"math"
	"math/rand"
	"strings"
)

type alphabet string

func newAlphabet(s string) (alphabet, error) {
	uniq := make(map[rune]struct{})
	for _, c := range s {
		if _, ok := uniq[c]; ok {
			return "", errors.New("presids: new alphabet: all characters in s must be unique")
		}
		uniq[c] = struct{}{}
	}
	return alphabet(s), nil
}

// Shuffle returns a new alphabet based on the shuffled characters of a
func (a alphabet) Shuffle(seed int64) alphabet {
	r := rand.New(rand.NewSource(seed))
	dst := make([]byte, len(a))
	perm := r.Perm(len(a))
	for i, j := range perm {
		dst[j] = byte(a[i])
	}
	return alphabet(dst)
}

func (a alphabet) Encode(n uint64) string {
	b := uint64(len(a))
	s := make([]byte, a.encodedLen(n))
	for i := 0; n > 0; i++ {
		s[i] = a[n%b]
		n /= b
	}
	return string(s)
}

func (a alphabet) Decode(s string) (uint64, error) {
	b := len(a)
	mag := 1
	var n uint64
	for _, c := range []byte(s) {
		x := strings.IndexByte(string(a), c)
		if x == -1 {
			return 0, errors.New("presids: decode alphabet: invalid input")
		}
		n += uint64(x * mag)
		mag *= b
	}
	return n, nil
}

func (a alphabet) encodedLen(n uint64) int {
	p := math.Log(float64(n))
	q := math.Log(float64(len(a)))
	r := p / q
	s := math.Ceil(r)
	return int(s)
}
