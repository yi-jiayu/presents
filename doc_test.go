package presents_test

import (
	"fmt"
	"log"

	"github.com/yi-jiayu/presents"
	"golang.org/x/crypto/blowfish"
)

// This example show how to encode and decode IDs.
func Example() {
	// 80-bit PRESENT block cipher key
	key := make([]byte, 10)
	p, err := presents.New(key, nil)
	if err != nil {
		log.Fatal(err)
	}

	s := p.Wrap(1213486160)
	fmt.Println(s)

	n, err := p.Unwrap(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
	// Output:
	// 90NyXHLckhA
	// 1213486160
}

// This example shows how to use the triple DES cipher instead of PRESENT.
func ExampleNewTripleDES() {
	// 24-byte triple DES key
	key := make([]byte, 24)
	p, err := presents.NewTripleDES(key, nil)
	if err != nil {
		log.Fatal(err)
	}

	s := p.Wrap(1213486160)
	fmt.Println(s)

	n, err := p.Unwrap(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
	// Output:
	// Yq4drxXQtcJ
	// 1213486160
}

// This example shows how to use a custom alphabet as well as shuffling it.
func Example_customAlphabet() {
	// 80-bit PRESENT block cipher key
	key := make([]byte, 10)
	options := &presents.Options{
		Alphabet: "_0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
		Shuffle:  true,
		Seed:     12,
	}
	p, err := presents.New(key, options)
	if err != nil {
		log.Fatal(err)
	}

	s := p.Wrap(1213486160)
	fmt.Println(s)

	n, err := p.Unwrap(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
	// Output:
	// tiKYxtU_2ZB
	// 1213486160
}

// You can also provide your own cipher.Block implementation as long as it has a 64-bit block size.
// This example shows how you can use the Blowfish cipher provided by the golang.org/x/crypto/blowfish package.
func ExampleNewWithCipher() {
	// 56-byte Blowfish key
	key := make([]byte, 56)
	c, err := blowfish.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}
	p, err := presents.NewWithCipher(c, nil)
	if err != nil {
		log.Fatal(err)
	}

	s := p.Wrap(1213486160)
	fmt.Println(s)

	n, err := p.Unwrap(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
	// Output:
	// woSdQdAYuiK
	// 1213486160
}
