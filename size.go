// Package size provides functionality for handling bytes and the conversion
// between various representations such as kilobytes and mebibytes and all the
// other alternatives.
package size

import "fmt"

// Size represents a number of bytes as an int64. The largest size that can be
// representated is math.Int64Max bytes.
type Size int64

// BinarySize simply provides an alternative fmt.Stringer implementation for Size
// that prints values using binary names, i.e. Mebibyte instead of Size's default
// Megabyte.
type BinarySize int64

const (
	Byte Size = 1 << (10 * iota)
	Kibibyte
	Mebibyte
	Gibibyte
	Tebibyte
	Pebibyte
	Exbibyte
)

const (
	Kilobyte Size = 1000 * Byte
	Megabyte Size = 1000 * Kilobyte
	Gigabyte Size = 1000 * Megabyte
	Terabyte Size = 1000 * Gigabyte
	Petabyte Size = 1000 * Terabyte
	Exabyte  Size = 1000 * Petabyte
)

// String implements fmt.Stringer for BinarySize. It prints out
// the resulting data in terms of Bytes, Kibibytes, Mebibytes, and
// so forth.
func (b BinarySize) String() string {
	switch {
	case Size(b) >= Exbibyte:
		return fmt.Sprintf("%.2fEiB", float64(b)/float64(Exbibyte))
	case Size(b) >= Pebibyte:
		return fmt.Sprintf("%.2fPiB", float64(b)/float64(Pebibyte))
	case Size(b) >= Tebibyte:
		return fmt.Sprintf("%.2fTiB", float64(b)/float64(Tebibyte))
	case Size(b) >= Gibibyte:
		return fmt.Sprintf("%.2fGiB", float64(b)/float64(Gibibyte))
	case Size(b) >= Mebibyte:
		return fmt.Sprintf("%.2fMiB", float64(b)/float64(Mebibyte))
	case Size(b) >= Kibibyte:
		return fmt.Sprintf("%.2fKiB", float64(b)/float64(Kibibyte))
	}
	return fmt.Sprintf("%.0fB", float64(b))
}

// String implements fmt.Stringer for Size. It prints out the resulting
// data in terms of Bytes, Kilobytes, Megabytes, and so forth.
func (s Size) String() string {
	switch {
	case s >= Exabyte:
		return fmt.Sprintf("%.2fEB", float64(s)/float64(Exabyte))
	case s >= Petabyte:
		return fmt.Sprintf("%.2fPB", float64(s)/float64(Petabyte))
	case s >= Terabyte:
		return fmt.Sprintf("%.2fTB", float64(s)/float64(Terabyte))
	case s >= Gigabyte:
		return fmt.Sprintf("%.2fGB", float64(s)/float64(Gigabyte))
	case s >= Megabyte:
		return fmt.Sprintf("%.2fMB", float64(s)/float64(Megabyte))
	case s >= Kilobyte:
		return fmt.Sprintf("%.2fKB", float64(s)/float64(Kilobyte))
	}
	return fmt.Sprintf("%.0fB", float64(s))
}

// Bytes returns the size as an integer byte count.
func (s Size) Bytes() int64 {
	return int64(s)
}

// Kibibytes returns the size as a floating point kibibyte count.
func (s Size) Kibibytes() float64 {
	whole := s / Kibibyte
	part := s % Kibibyte
	return float64(whole) + float64(part)/(1<<(10))
}

// Kilobytes returns the size as a floating point kilobyte count.
func (s Size) Kilobytes() float64 {
	whole := s / Kilobyte
	part := s % Kilobyte
	return float64(whole) + float64(part)/1e3
}

// Mebibytes returns the size as a floating point mebibyte count.
func (s Size) Mebibytes() float64 {
	whole := s / Mebibyte
	part := s % Mebibyte
	return float64(whole) + float64(part)/(1<<(10*2))
}

// Megabytes returns the size as a floating point megabyte count.
func (s Size) Megabytes() float64 {
	whole := s / Megabyte
	part := s % Megabyte
	return float64(whole) + float64(part)/1e6
}

// Gibibytes returns the size as a floating point gibibytes count.
func (s Size) Gibibytes() float64 {
	whole := s / Gibibyte
	part := s % Gibibyte
	return float64(whole) + float64(part)/(1<<(10*3))
}

// Gigabytes returns the size as a floating point gigabyte count.
func (s Size) Gigabytes() float64 {
	whole := s / Gigabyte
	part := s % Gigabyte
	return float64(whole) + float64(part)/1e9
}

// Tebibytes returns the size as a floating point tebibytes count.
func (s Size) Tebibytes() float64 {
	whole := s / Tebibyte
	part := s % Tebibyte
	return float64(whole) + float64(part)/(1<<(10*4))
}

// Terabytes returns the size as a floating point terabyte count.
func (s Size) Terabytes() float64 {
	whole := s / Terabyte
	part := s % Terabyte
	return float64(whole) + float64(part)/1e12
}

// Pebibytes returns the size as a floating point pebibyte count.
func (s Size) Pebibytes() float64 {
	whole := s / Pebibyte
	part := s % Pebibyte
	return float64(whole) + float64(part)/(1<<(10*5))
}

// Petabytes returns the size as a floating point petabyte count.
func (s Size) Petabytes() float64 {
	whole := s / Petabyte
	part := s % Petabyte
	return float64(whole) + float64(part)/1e15
}

// Exbibytes returns the size as a floating point petabyte count.
func (s Size) Exbibytes() float64 {
	whole := s / Exbibyte
	part := s % Exbibyte
	return float64(whole) + float64(part)/(1<<(10*6))
}

// Exabytes returns the size as a floating point petabyte count.
func (s Size) Exabytes() float64 {
	whole := s / Exabyte
	part := s % Exabyte
	return float64(whole) + float64(part)/1e18
}
