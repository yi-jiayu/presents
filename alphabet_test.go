package presents

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlphabet_Encode(t *testing.T) {
	s := DefaultAlphabet.Encode(1213486160)
	expected := "QBf7K1"
	assert.Equal(t, expected, s)
}

func TestAlphabet_Decode(t *testing.T) {
	n, err := DefaultAlphabet.Decode("QBf7K1")
	if err != nil {
		t.Fatal(err)
	}
	var expected uint64 = 1213486160
	assert.Equal(t, expected, n)
}

func TestAlphabet_EncodedLen(t *testing.T) {
	actual := DefaultAlphabet.encodedLen(1213486160)
	expected := 6
	assert.Equal(t, expected, actual)
}

func TestAlphabet_Shuffle(t *testing.T) {
	shuffled := DefaultAlphabet.Shuffle(1)
	expected := "gd2J1bExUClwVnmNXoB6H0ifMqGKLkpz5cv8O9RQhAWrS7s4D3IujFePTatZyY"
	assert.Equal(t, expected, string(shuffled))
}
