package present

// key128 implements the PRESENT key schedule for 128-bit keys.
type key128 struct {
	A, B uint64
}

func (k *key128) copy() key {
	cpy := *k
	return &cpy
}

func (k *key128) rotate() {
	w := k.A & 0xfffffffffffffff8
	x := k.A & 0x7
	y := k.B & 0xfffffffffffffff8
	z := k.B & 0x7

	k.A = (x << 61) + (y >> 3)
	k.B = (z << 61) + (w >> 3)
}

func (k *key128) sBox() {
	x := (k.A >> 60) & 0xF
	y := (k.A >> 56) & 0xF

	p := sBox[x]
	q := sBox[y]

	a := uint64(p) << 60
	b := uint64(q) << 56
	c := k.A & 0x00ffffffffffffff

	k.A = a + b + c
}

func (k *key128) xor(ctr uint64) {
	w := (k.A & 0x7) << 2
	x := (k.B >> 62) & 0x3
	y := w + x
	z := y ^ ctr

	p := (z >> 2) & 0x7
	q := (z & 0x3) << 62
	r := k.A & 0xfffffffffffffff8
	s := k.B & 0x3fffffffffffffff

	k.A = p + r
	k.B = q + s
}

func (k *key128) roundKey() uint64 {
	return k.A
}

// newKey128 returns a new 128-bit PRESENT key register from the provided key bytes.
func newKey128(key []byte) *key128 {
	A, B := decompose(key)
	return &key128{A, B}
}
