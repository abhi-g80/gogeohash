package gogeohash

import (
	"fmt"
	"testing"
)

func BenchmarkToChars(b *testing.B) {
	inputs := []uint64{
		14520001368503071193,
	}

	for i, v := range inputs {
		n := fmt.Sprintf("%s_%d", b.Name(), i)
		b.Run(n, func(b *testing.B) {
			for b.Loop() {
				toChars(v)
			}
		})
	}
}

func BenchmarkToCharsUnrolled(b *testing.B) {
	inputs := []uint64{
		14520001368503071193,
	}

	for i, v := range inputs {
		n := fmt.Sprintf("%s_%d", b.Name(), i)
		b.Run(n, func(b *testing.B) {
			for b.Loop() {
				toCharsUnrolled(v)
			}
		})
	}
}

func BenchmarkToCharsUnrolledBytes(b *testing.B) {
	inputs := []uint64{
		14520001368503071193,
	}

	for i, v := range inputs {
		n := fmt.Sprintf("%s_%d", b.Name(), i)
		b.Run(n, func(b *testing.B) {
			for b.Loop() {
				toCharsUnrolledBytes(v)
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

func TestToCharsUnrolledBytes(t *testing.T) {
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
			r := toCharsUnrolledBytes(v.input)
			if string(r[:]) != v.chars {
				t.Errorf("%s: expected: %s got: %+v", t.Name(), v.chars, r)
			}
		})
	}
}

func TestToCharsUnrolled(t *testing.T) {
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
			r := toCharsUnrolled(v.input)
			if r != v.chars {
				t.Errorf("%s: expected: %s got: %+v", t.Name(), v.chars, r)
			}
		})
	}
}
