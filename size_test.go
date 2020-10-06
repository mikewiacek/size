package size

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestStrings(t *testing.T) {
	cases := []struct {
		numBytes  int64
		size      Size
		str       string
		binaryStr string
	}{
		{0, 0 * Byte, "0B", "0B"},
		{1, 1 * Byte, "1B", "1B"},
		{1, 1 * Byte, "1B", "1B"},
		{100, 100 * Byte, "100B", "100B"},
		{1000, 1 * Kilobyte, "1.00KB", "1000B"},
		{1024, 1 * Kibibyte, "1.02KB", "1.00KiB"},
		{1000000, 1 * Megabyte, "1.00MB", "976.56KiB"},
		{1048576, 1 * Mebibyte, "1.05MB", "1.00MiB"},
		{1000000000, 1 * Gigabyte, "1.00GB", "953.67MiB"},
		{1073741824, 1 * Gibibyte, "1.07GB", "1.00GiB"},
		{1000000000000, 1 * Terabyte, "1.00TB", "931.32GiB"},
		{1099511627776, 1 * Tebibyte, "1.10TB", "1.00TiB"},
		{1000000000000000, 1 * Petabyte, "1.00PB", "909.49TiB"},
		{1125899906842624, 1 * Pebibyte, "1.13PB", "1.00PiB"},
		{1000000000000000000, 1 * Exabyte, "1.00EB", "888.18PiB"},
		{1152921504606846976, 1 * Exbibyte, "1.15EB", "1.00EiB"},
		// This is the max value representable by Size.
		{math.MaxInt64, 7*Exbibyte + (1*Exbibyte - 1*Byte), "9.22EB", "8.00EiB"},
	}

	for _, c := range cases {
		s := Size(c.numBytes)
		if s != c.size {
			t.Errorf("Size from constant arithmetic got %d, want %d", s, c.size)
		}
		if s.String() != c.str {
			t.Errorf("%d bytes got a string of %q, want %q", c.numBytes, s.String(), c.str)
		}
		b := BinarySize(c.numBytes)
		if b.String() != c.binaryStr {
			t.Errorf("%d bytes got a binary size string of %q, want %q", c.numBytes, b.String(), c.binaryStr)
		}
	}
}

func TestConversions(t *testing.T) {
	// opt compares float64s such that extreme precision isn't required for equality.
	opt := cmp.Comparer(func(x, y float64) bool {
		delta := math.Abs(x - y)
		mean := math.Abs(x+y) / 2.0
		return delta/mean < 0.00001
	})

	cases := []struct {
		start        Size
		wantExbibyte float64
		wantExabyte  float64
		wantPebibyte float64
		wantPetabyte float64
		wantTebibyte float64
		wantTerabyte float64
		wantGibibyte float64
		wantGigabyte float64
		wantMebibyte float64
		wantMegabyte float64
		wantKibibyte float64
		wantKilobyte float64
		wantByte     int64
	}{
		{
			start:        1441151880758558720 * Byte, // 1.25 Exbibytes
			wantExbibyte: 1.25,
			wantExabyte:  1.441152,
			wantPebibyte: 1280.0,
			wantPetabyte: 1441.151881,
			wantTebibyte: 1310720.0,
			wantTerabyte: 1441151.880759,
			wantGibibyte: 1342177280.0,
			wantGigabyte: 1441151880.758559,
			wantMebibyte: 1374389534720.0,
			wantMegabyte: 1441151880758.558838,
			wantKibibyte: 1407374883553280.0,
			wantKilobyte: 1441151880758558.75,
			wantByte:     1441151880758558720, // Same as start
		},
	}

	for _, c := range cases {
		if got := c.start.Exbibytes(); !cmp.Equal(got, c.wantExbibyte, opt) {
			t.Errorf("For initial value (%d Byte).Exbibytes() got: %f, want: %f", c.start, got, c.wantExbibyte)
		}
		if got := c.start.Exabytes(); !cmp.Equal(got, c.wantExabyte, opt) {
			t.Errorf("For initial value (%d Byte).Exabytes() got: %f, want: %f", c.start, got, c.wantExabyte)
		}
		if got := c.start.Pebibytes(); !cmp.Equal(got, c.wantPebibyte, opt) {
			t.Errorf("For initial value (%d Byte).Pebibytes() got: %f, want: %f", c.start, got, c.wantPebibyte)
		}
		if got := c.start.Petabytes(); !cmp.Equal(got, c.wantPetabyte, opt) {
			t.Errorf("For initial value (%d Byte).Petabytes() got: %f, want: %f", c.start, got, c.wantPetabyte)
		}
		if got := c.start.Tebibytes(); !cmp.Equal(got, c.wantTebibyte, opt) {
			t.Errorf("For initial value (%d Byte).Tebibytes() got: %f, want: %f", c.start, got, c.wantTebibyte)
		}
		if got := c.start.Terabytes(); !cmp.Equal(got, c.wantTerabyte, opt) {
			t.Errorf("For initial value (%d Byte).Terabytes() got: %f, want: %f", c.start, got, c.wantTerabyte)
		}
		if got := c.start.Gibibytes(); !cmp.Equal(got, c.wantGibibyte, opt) {
			t.Errorf("For initial value (%d Byte).Gibibytes() got: %f, want: %f", c.start, got, c.wantGibibyte)
		}
		if got := c.start.Gigabytes(); !cmp.Equal(got, c.wantGigabyte, opt) {
			t.Errorf("For initial value (%d Byte).Gigabytes() got: %f, want: %f", c.start, got, c.wantGigabyte)
		}
		if got := c.start.Mebibytes(); !cmp.Equal(got, c.wantMebibyte, opt) {
			t.Errorf("For initial value (%d Byte).Mebibytes() got: %f, want: %f", c.start, got, c.wantMebibyte)
		}
		if got := c.start.Megabytes(); !cmp.Equal(got, c.wantMegabyte, opt) {
			t.Errorf("For initial value (%d Byte).Megabytes() got: %f, want: %f", c.start, got, c.wantMegabyte)
		}
		if got := c.start.Kibibytes(); !cmp.Equal(got, c.wantKibibyte, opt) {
			t.Errorf("For initial value (%d Byte).Kibibytes() got: %f, want: %f", c.start, got, c.wantKibibyte)
		}
		if got := c.start.Kilobytes(); !cmp.Equal(got, c.wantKilobyte, opt) {
			t.Errorf("For initial value (%d Byte).Kilobytes() got: %f, want: %f", c.start, got, c.wantKilobyte)
		}
		if got := c.start.Bytes(); got != c.wantByte {
			t.Errorf("For initial value (%d Byte).Bytes() got: %d, want: %d", c.start, got, c.wantByte)
		}
	}
}
