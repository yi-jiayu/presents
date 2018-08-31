# presents
[![GoDoc](https://godoc.org/github.com/yi-jiayu/presents?status.svg)](https://godoc.org/github.com/yi-jiayu/presents)
[![Build Status](https://travis-ci.com/yi-jiayu/presents.svg?branch=master)](https://travis-ci.com/yi-jiayu/presents)
[![codecov](https://codecov.io/gh/yi-jiayu/presents/branch/master/graph/badge.svg)](https://codecov.io/gh/yi-jiayu/presents)
[![Go Report Card](https://goreportcard.com/badge/github.com/yi-jiayu/presents)](https://goreportcard.com/report/github.com/yi-jiayu/presents)

Like [hashids](https://hashids.org/), but based on block ciphers.

Inspired by [this StackOverflow answer](https://stackoverflow.com/a/8554984) suggesting using a block cipher to obfuscate IDs.

## How it works

The PRESENT block cipher [1] operates on 8-byte (64-bit) blocks and supports key lengths of 80-bits or 128-bits. It can be used to create a reversible mapping from 64-bit integers to 64-bit integers.

The resultant 64-bit integer is then converted to and from a string using an arbitrary change of base algorithm and a provided alphabet.

Triple DES, which also has a 64-bit block size, can be used as well.

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
	// with 3DES instead of PRESENT
	// key := make([]byte, 24)
	// p, err := presents.NewTripleDES(key, nil)
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

## Performance
Triple DES is about 20 times faster than PRESENT:

```console
$ go test -bench .
goos: windows
goarch: amd64
pkg: github.com/yi-jiayu/presents
BenchmarkPresents_Wrap-4                  100000             11688 ns/op
BenchmarkPresentsTripleDES_Wrap-4        1000000              1895 ns/op
PASS
ok      github.com/yi-jiayu/presents    3.566s
```

## References 
1. Bogdanov A. et al. (2007) PRESENT: An Ultra-Lightweight Block Cipher. In: Paillier P., Verbauwhede I. (eds) Cryptographic Hardware and Embedded Systems - CHES 2007. CHES 2007. Lecture Notes in Computer Science, vol 4727. Springer, Berlin, Heidelberg ([pdf](http://www.lightweightcrypto.org/present/present_ches2007.pdf))
