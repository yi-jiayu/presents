// Package present implements the ultra-lightweight block cipher PRESENT as defined by Bogdanov et al. [1].
//
// 1. Bogdanov A. et al. (2007) PRESENT: An Ultra-Lightweight Block Cipher.
// In: Paillier P., Verbauwhede I. (eds) Cryptographic Hardware and Embedded Systems - CHES 2007.
// CHES 2007. Lecture Notes in Computer Science, vol 4727. Springer, Berlin, Heidelberg
package present

import (
	"crypto/cipher"
	"strconv"
)

// KeySizeError represents an invalid PRESENT key length.
type KeySizeError int

func (k KeySizeError) Error() string {
	return "present: invalid key size " + strconv.Itoa(int(k))
}

// NewCipher creates a new cipher.Block.
// The argument should be the PRESENT key,
// which is either 10 or 16 bytes long
// for key lengths of 80 bits and 128 bits respectively.
func NewCipher(key []byte) (b cipher.Block, err error) {
	k := len(key)
	switch k {
	default:
		err = KeySizeError(k)
	case 10:
		b = &block{
			Key: newKey80(key),
		}
	case 16:
		b = &block{
			Key: newKey128(key),
		}
	}
	return
}
