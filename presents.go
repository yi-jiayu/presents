// Package presents implements a block cipher-based method of converting 64-bit unsigned integers to and from strings.
//
// The intended application is towards the obfuscation of sequential database IDs.
package presents

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/yi-jiayu/PRESENT.go"
)

// DefaultAlphabet contains printable characters from 0-9, A-Z and a-z, similar to a base62 encoding.
const DefaultAlphabet = alphabet("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

// Presents contains a cipher.Block implementing PRESENT
// and an alphabet for converting between 64-bit integers and strings.
type Presents struct {
	cipher   cipher.Block
	alphabet alphabet
}

// Options can be passed to New to customise the alphabet to be used.
type Options struct {
	Alphabet string
	Shuffle  bool
	Seed     int64
}

// New creates a new Presents struct using the PRESENT block cipher.
// If options.Alphabet is not the empty string, it will be used as the alphabet.
// If options.Shuffle is true, the alphabet will be shuffled based on options.Seed.
func New(key []byte, options *Options) (*Presents, error) {
	c, err := present.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("presents: New: %v", err)
	}
	return NewWithCipher(c, options)
}

// NewWithCipher returns a new Presents instance from the provided cipher.Block and options.
// The provided cipher.Block should have a 64-bit block size.
func NewWithCipher(c cipher.Block, options *Options) (*Presents, error) {
	if c.BlockSize() != 8 {
		return nil, errors.New("presents: NewWithCipher: cipher should have a 64-bit block size")
	}

	a := DefaultAlphabet
	if options != nil {
		if options.Alphabet != "" {
			var err error
			a, err = newAlphabet(options.Alphabet)
			if err != nil {
				return nil, err
			}
		}
		if options.Shuffle {
			a = a.Shuffle(options.Seed)
		}
	}
	return &Presents{
		cipher:   c,
		alphabet: a,
	}, nil
}

// NewTripleDES creates a new Presents struct using Triple DES instead of PRESENT.
// If options.Alphabet is not the empty string, it will be used as the alphabet.
// If options.Shuffle is true, the alphabet will be shuffled based on options.Seed.
func NewTripleDES(key []byte, options *Options) (*Presents, error) {
	c, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, fmt.Errorf("presents: NewTripleDES: %v", err)
	}
	return NewWithCipher(c, options)
}

// Wrap converts an unsigned 64-bit integer to a string.
func (p *Presents) Wrap(n uint64) string {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, n)
	dst := make([]byte, 8)
	p.cipher.Encrypt(dst, b)
	n = binary.BigEndian.Uint64(dst)
	return p.alphabet.Encode(n)
}

// Unwrap converts a string back to an unsigned 64-bit integer.
// It returns an error if the string cannot be converted using the given alphabet.
func (p *Presents) Unwrap(s string) (uint64, error) {
	n, err := p.alphabet.Decode(s)
	if err != nil {
		return 0, err
	}
	src := make([]byte, 8)
	binary.BigEndian.PutUint64(src, n)
	dst := make([]byte, 8)
	p.cipher.Decrypt(dst, src)
	n = binary.BigEndian.Uint64(dst)
	return n, nil
}
