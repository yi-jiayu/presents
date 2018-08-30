package present

// key represents an implementation of the PRESENT key schedule for deriving the round keys.
type key interface {
	copy() key
	rotate()
	sBox()
	xor(ctr uint64)
	roundKey() uint64
}

// decompose converts an 80-bit or 128-bit key into a pair of 64-bit integers.
func decompose(key []byte) (A, B uint64) {
	for i, x := range key {
		if i < 8 {
			shift := 56 - i*8
			A |= uint64(x) << uint64(shift)
		} else {
			shift := 120 - i*8
			B |= uint64(x) << uint64(shift)
		}
	}
	return
}

// updateKey updates key based on the key schedule and current round counter.
func updateKey(k key, ctr int) {
	k.rotate()
	k.sBox()
	k.xor(uint64(ctr))
}

// expandKey copies and expands k, sending the derived round keys to roundKeys.
func expandKey(k key, roundKeys chan<- uint64) {
	k = k.copy()
	for ctr := 0; ctr < numRounds; ctr++ {
		roundKeys <- k.roundKey()
		updateKey(k, ctr+1)
	}
	roundKeys <- k.roundKey()
	close(roundKeys)
}
