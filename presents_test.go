package presents_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yi-jiayu/presents"
)

func TestPresents_Wrap(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		key := make([]byte, 10)
		p, err := presents.New(key, nil)
		if err != nil {
			t.Fatal(err)
		}
		s := p.Wrap(1213486160)
		expected := "90NyXHLckhA"
		assert.Equal(t, expected, s)
	})
	t.Run("shuffled", func(t *testing.T) {
		key := make([]byte, 10)
		p, err := presents.New(key, &presents.Options{
			Shuffle: true,
		})
		if err != nil {
			t.Fatal(err)
		}
		s := p.Wrap(1213486160)
		expected := "w3CBcIAvNMd"
		assert.Equal(t, expected, s)
	})
	t.Run("custom alphabet", func(t *testing.T) {
		key := make([]byte, 10)
		p, err := presents.New(key, &presents.Options{
			Alphabet: "zyxwvutsrqponmlkjihgfedcbaZYXWVUTSRQPONMLKJIHGFEDCBA9876543210",
		})
		if err != nil {
			t.Fatal(err)
		}
		s := p.Wrap(1213486160)
		expected := "qzc1SieNFIp"
		assert.Equal(t, expected, s)
	})
}

func TestPresents_Unwrap(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		key := make([]byte, 10)
		p, err := presents.New(key, nil)
		if err != nil {
			t.Fatal(err)
		}
		n, err := p.Unwrap("90NyXHLckhA")
		if err != nil {
			t.Fatal(err)
		}
		var expected uint64 = 1213486160
		assert.Equal(t, expected, n)
	})
	t.Run("shuffled", func(t *testing.T) {
		key := make([]byte, 10)
		p, err := presents.New(key, &presents.Options{
			Shuffle: true,
		})
		if err != nil {
			t.Fatal(err)
		}
		s, err := p.Unwrap("w3CBcIAvNMd")
		if err != nil {
			t.Fatal(err)
		}
		var expected uint64 = 1213486160
		assert.Equal(t, expected, s)
	})
	t.Run("custom alphabet", func(t *testing.T) {
		key := make([]byte, 10)
		p, err := presents.New(key, &presents.Options{
			Alphabet: "zyxwvutsrqponmlkjihgfedcbaZYXWVUTSRQPONMLKJIHGFEDCBA9876543210",
		})
		if err != nil {
			t.Fatal(err)
		}
		s, err := p.Unwrap("qzc1SieNFIp")
		if err != nil {
			t.Fatal(err)
		}
		var expected uint64 = 1213486160
		assert.Equal(t, expected, s)
	})
}

func TestPresentsTripleDES_Wrap(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		key := make([]byte, 24)
		p, err := presents.NewTripleDES(key, nil)
		if err != nil {
			t.Fatal(err)
		}
		s := p.Wrap(1213486160)
		expected := "Yq4drxXQtcJ"
		assert.Equal(t, expected, s)
	})
	t.Run("shuffled", func(t *testing.T) {
		key := make([]byte, 24)
		p, err := presents.NewTripleDES(key, &presents.Options{
			Shuffle: true,
		})
		if err != nil {
			t.Fatal(err)
		}
		s := p.Wrap(1213486160)
		expected := "D7sYkicVqvE"
		assert.Equal(t, expected, s)
	})
	t.Run("custom alphabet", func(t *testing.T) {
		key := make([]byte, 24)
		p, err := presents.NewTripleDES(key, &presents.Options{
			Alphabet: "zyxwvutsrqponmlkjihgfedcbaZYXWVUTSRQPONMLKJIHGFEDCBA9876543210",
		})
		if err != nil {
			t.Fatal(err)
		}
		s := p.Wrap(1213486160)
		expected := "R9vM82SZ6Ng"
		assert.Equal(t, expected, s)
	})
}

func TestPresentsTripleDES_Unwrap(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		key := make([]byte, 24)
		p, err := presents.NewTripleDES(key, nil)
		if err != nil {
			t.Fatal(err)
		}
		n, err := p.Unwrap("Yq4drxXQtcJ")
		if err != nil {
			t.Fatal(err)
		}
		var expected uint64 = 1213486160
		assert.Equal(t, expected, n)
	})
	t.Run("shuffled", func(t *testing.T) {
		key := make([]byte, 24)
		p, err := presents.NewTripleDES(key, &presents.Options{
			Shuffle: true,
		})
		if err != nil {
			t.Fatal(err)
		}
		s, err := p.Unwrap("D7sYkicVqvE")
		if err != nil {
			t.Fatal(err)
		}
		var expected uint64 = 1213486160
		assert.Equal(t, expected, s)
	})
	t.Run("custom alphabet", func(t *testing.T) {
		key := make([]byte, 24)
		p, err := presents.NewTripleDES(key, &presents.Options{
			Alphabet: "zyxwvutsrqponmlkjihgfedcbaZYXWVUTSRQPONMLKJIHGFEDCBA9876543210",
		})
		if err != nil {
			t.Fatal(err)
		}
		s, err := p.Unwrap("R9vM82SZ6Ng")
		if err != nil {
			t.Fatal(err)
		}
		var expected uint64 = 1213486160
		assert.Equal(t, expected, s)
	})
}
