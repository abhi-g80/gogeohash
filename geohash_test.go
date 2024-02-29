package gogeohash

import (
	"fmt"
	"testing"
)

func BenchmarkToChars(b *testing.B) {
	inputs := []uint64{
		14520001368503071193,
	}

	for _, v := range inputs {
		n := fmt.Sprintf("%s_%d", b.Name(), v)
		b.Run(n, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				toChars(v)
			}
		})
	}
}

func BenchmarkToCharsUnrolled(b *testing.B) {
	inputs := []uint64{
		14520001368503071193,
	}

	for _, v := range inputs {
		n := fmt.Sprintf("%s_%d", b.Name(), v)
		b.Run(n, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				toCharsUnrolled(v)
			}
		})
	}
}

func TestGeoHashEncode(t *testing.T) {
	testcases := []struct {
		lon float64
		lat float64
		exp uint64
	}{
		{
			lat: 12.34,
			lon: 56.78,
			exp: 14520001368503071193,
		},
	}

	for i, v := range testcases {
		n := fmt.Sprintf("%s_%d_%f_%f", t.Name(), i, v.lat, v.lon)
		t.Run(n, func(t *testing.T) {
			r := GeoHashEncode(v.lat, v.lon)
			if r != v.exp {
				t.Errorf("%s: expected: %d got: %d", t.Name(), v.exp, r)
			}
		})
	}
}

func TestToChars(t *testing.T) {
	testcases := []struct {
		input uint64
		chars string
	}{
		{
			input: 14520001368503071193,
			chars: "t60qhn3vtysx",
		},
	}

	for _, v := range testcases {
		n := fmt.Sprintf("%s_%d", t.Name(), v.input)
		t.Run(n, func(t *testing.T) {
			r := toChars(v.input)
			if r != v.chars {
				t.Errorf("%s: expected: %s got: %s", t.Name(), v.chars, r)
			}
		})
	}
}
