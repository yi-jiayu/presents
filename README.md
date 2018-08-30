# presents
Like [hashids](https://hashids.org/), but based on the PRESENT block cipher as defined by Bogdanov et al. [1].

Inspired by [this StackOverflow answer](https://stackoverflow.com/a/8554984) suggesting using a block cipher to obfuscate IDs.

## How it works

The PRESENT block cipher operates on 8-byte (64-bit) blocks and supports key lengths of 80-bits or 128-bits. It can be used to create a reversible mapping from 64-bit integers to 64-bit integers.

The resultant 64-bit integer is then converted to and from a string using an arbitrary change of base algorithm and a provided alphabet.

## Usage
```go
package main

import (
	"fmt"
	"log"

	"github.com/yi-jiayu/presents"
)

func main() {
	// 80-bit PRESENT block cipher key
	key := make([]byte, 10)
	p, err := presents.New(key, nil)
	if err != nil {
		log.Fatal(err)
	}

	s := p.Wrap(1213486160)
	fmt.Println(s) // 90NyXHLckhA

	n, err := p.Unwrap(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n) // 1213486160
}
```


## References 
1. Bogdanov A. et al. (2007) PRESENT: An Ultra-Lightweight Block Cipher. In: Paillier P., Verbauwhede I. (eds) Cryptographic Hardware and Embedded Systems - CHES 2007. CHES 2007. Lecture Notes in Computer Science, vol 4727. Springer, Berlin, Heidelberg ([pdf](http://www.lightweightcrypto.org/present/present_ches2007.pdf))
