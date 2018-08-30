package present

// key80 implements the PRESENT key schedule for 80-bit keys.
type key80 struct {
	A, B uint64
}

func (k *key80) copy() key {
	cpy := *k
	return &cpy
}

func (k *key80) rotate() {
	w := k.A & 0xfffffffffff80000
	x := k.A & 0x7fff8
	y := k.A & 0x7
	z := k.B & 0xffff000000000000

	k.A = (y << 61) + (z >> 3) + (w >> 19)
	k.B = x << 45
}

func (k *key80) sBox() {
	w := (k.A >> 60) & 0xF
	x := sBox[w]
	y := uint64(x) << 60
	z := k.A & 0x0fffffffffffffff

	k.A = y + z
}

func (k *key80) xor(ctr uint64) {
	w := (k.A & 0xF) << 1
	x := (k.B >> 63) & 1
	y := w + x
	z := y ^ ctr

	p := (z & 0x1e) >> 1
	q := (z & 1) << 63
	r := k.A & 0xfffffffffffffff0
	s := k.B & 0x7fffffffffffffff

	k.A = p + r
	k.B = q + s
}

func (k *key80) roundKey() uint64 {
	return k.A
}

// newKey80 returns a new 80-bit PRESENT key register from the provided key bytes.
func newKey80(key []byte) *key80 {
	A, B := decompose(key)
	return &key80{A, B}
}
