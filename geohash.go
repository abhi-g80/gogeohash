package gogeohash

const (
	MIN_LATITUDE = -90.0
	MAX_LATITUDE = 90.0

	MIN_LONGITUDE = -180.0
	MAX_LONGITUDE = 180.0

	STEP = 32

	TEXTREP = "0123456789bcdefghjkmnpqrstuvwxyz"
)

// Interleave 2 uint32's bit representation.
// Interleave lower bits of a and b, so the bits of a are in the even positions
// and bits from y in the odd. Checkout the following page many such bit tweaks
// https://graphics.stanford.edu/~seander/bithacks.html#InterleaveBMN
func interleave64(a, b uint32) uint64 {
	B := [5]uint64{
		0x5555555555555555,
		0x3333333333333333,
		0x0F0F0F0F0F0F0F0F,
		0x00FF00FF00FF00FF,
		0x0000FFFF0000FFFF,
	}
	S := [5]uint32{1, 2, 4, 8, 16}

	x := uint64(a)
	y := uint64(b)

	x = (x | (x << S[4])) & B[4]
	y = (y | (y << S[4])) & B[4]

	x = (x | (x << S[3])) & B[3]
	y = (y | (y << S[3])) & B[3]

	x = (x | (x << S[2])) & B[2]
	y = (y | (y << S[2])) & B[2]

	x = (x | (x << S[1])) & B[1]
	y = (y | (y << S[1])) & B[1]

	x = (x | (x << S[0])) & B[0]
	y = (y | (y << S[0])) & B[0]

	return x | (y << 1)
}

// De-Interleave 2 uint64's bit representation.
//
//lint:ignore U1000 will use it later
func deinterleave64(i uint64) uint64 {
	B := [6]uint64{
		0x5555555555555555,
		0x3333333333333333,
		0x0F0F0F0F0F0F0F0F,
		0x00FF00FF00FF00FF,
		0x0000FFFF0000FFFF,
		0x00000000FFFFFFFF,
	}
	S := [6]uint32{0, 1, 2, 4, 8, 16}

	x := i
	y := i >> 1

	x = (x | (x >> S[0])) & B[0]
	y = (y | (y >> S[0])) & B[0]

	x = (x | (x >> S[1])) & B[1]
	y = (y | (y >> S[1])) & B[1]

	x = (x | (x >> S[2])) & B[2]
	y = (y | (y >> S[2])) & B[2]

	x = (x | (x >> S[3])) & B[3]
	y = (y | (y >> S[3])) & B[3]

	x = (x | (x >> S[4])) & B[4]
	y = (y | (y >> S[4])) & B[4]

	x = (x | (x >> S[5])) & B[5]
	y = (y | (y >> S[5])) & B[5]

	return x | (y << 32)
}

// Map each 5 bit to TEXTREP and produce a string representation of the geohash.
func toChars(x uint64) string {
	var s string
	for i := 59; i >= 0; i -= 5 {
		s += string(TEXTREP[(x&(((1<<5)-1)<<i))>>i])
	}
	return s
}

// Map each 5 bit to TEXTREP and produce a string representation of the geohash.
func toCharsUnrolled(x uint64) string {
	var s string
	s += string(TEXTREP[(x&(((1<<5)-1)<<59))>>59])
	s += string(TEXTREP[(x&(((1<<5)-1)<<54))>>54])
	s += string(TEXTREP[(x&(((1<<5)-1)<<49))>>49])
	s += string(TEXTREP[(x&(((1<<5)-1)<<44))>>44])
	s += string(TEXTREP[(x&(((1<<5)-1)<<39))>>39])
	s += string(TEXTREP[(x&(((1<<5)-1)<<34))>>34])
	s += string(TEXTREP[(x&(((1<<5)-1)<<29))>>29])
	s += string(TEXTREP[(x&(((1<<5)-1)<<24))>>24])
	s += string(TEXTREP[(x&(((1<<5)-1)<<19))>>19])
	s += string(TEXTREP[(x&(((1<<5)-1)<<14))>>14])
	s += string(TEXTREP[(x&(((1<<5)-1)<<9))>>9])
	s += string(TEXTREP[(x&(((1<<5)-1)<<4))>>4])
	return s
}

func toCharsUnrolledBytes(x uint64) [12]byte {
	var s [12]byte
	s[0] = TEXTREP[(x&(((1<<5)-1)<<59))>>59]
	s[1] = TEXTREP[(x&(((1<<5)-1)<<54))>>54]
	s[2] = TEXTREP[(x&(((1<<5)-1)<<49))>>49]
	s[3] = TEXTREP[(x&(((1<<5)-1)<<44))>>44]
	s[4] = TEXTREP[(x&(((1<<5)-1)<<39))>>39]
	s[5] = TEXTREP[(x&(((1<<5)-1)<<34))>>34]
	s[6] = TEXTREP[(x&(((1<<5)-1)<<29))>>29]
	s[7] = TEXTREP[(x&(((1<<5)-1)<<24))>>24]
	s[8] = TEXTREP[(x&(((1<<5)-1)<<19))>>19]
	s[9] = TEXTREP[(x&(((1<<5)-1)<<14))>>14]
	s[10] = TEXTREP[(x&(((1<<5)-1)<<9))>>9]
	s[11] = TEXTREP[(x&(((1<<5)-1)<<4))>>4]
	return s
}

// For a latitude and longitude value, compute its GeoHash.
func GeoHashEncode(lat, lon float64) uint64 {
	var latOffset, lonOffset float64

	latOffset = (lat - MIN_LATITUDE) / (MAX_LATITUDE - MIN_LATITUDE)
	lonOffset = (lon - MIN_LONGITUDE) / (MAX_LONGITUDE - MIN_LONGITUDE)

	latOffset *= 1 << STEP
	lonOffset *= 1 << STEP

	return interleave64(uint32(latOffset), uint32(lonOffset))
}

func GeoHashDecode(s string) (float64, float64, error) {
	return 0.0, 0.0, nil
}
