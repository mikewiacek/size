package size

import "testing"
import "math"

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
