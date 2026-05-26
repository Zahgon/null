package null

import (
	"database/sql"
)

// Byte is an nullable byte.
// It does not consider zero values to be null.
// It will decode to null, not zero, if null.
type Byte struct {
	sql.NullByte
}

// NewByte creates a new Byte.
func NewByte(b byte, valid bool) Byte { _ = "STUB: not implemented"; return *new(Byte) }

// ByteFrom creates a new Byte that will always be valid.
func ByteFrom(b byte) Byte {
	_ = "STUB: not implemented"
	return *

	// ByteFromPtr creates a new Byte that be null if i is nil.
	new(Byte)
}

func ByteFromPtr(b *byte) Byte { _ = "STUB: not implemented"; return *new(Byte) }

// ValueOrZero returns the inner value if valid, otherwise zero.
func (b Byte) ValueOrZero() byte { _ = "STUB: not implemented"; return 0 }

// ValueOr returns the inner value if valid, otherwise v.
func (b Byte) ValueOr(v byte) byte { _ = "STUB: not implemented"; return 0 }

// UnmarshalJSON implements json.Unmarshaler.
// It supports number, string, and null input.
// 0 will not be considered a null Byte.
func (b *Byte) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalText implements encoding.TextUnmarshaler.
// It will unmarshal to a null Byte if the input is blank.
// It will return an error if the input is not an integer, blank, or "null".
func (b *Byte) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalJSON implements json.Marshaler.
// It will encode null if this Byte is null.
func (b Byte) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalText implements encoding.TextMarshaler.
// It will encode a blank string if this Byte is null.
func (b Byte) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// SetValid changes this Byte's value and also sets it to be non-null.
func (b *Byte) SetValid(n byte) { _ = "STUB: not implemented"; return }

// Ptr returns a pointer to this Byte's value, or a nil pointer if this Byte is null.
func (b Byte) Ptr() *byte { _ = "STUB: not implemented"; return nil }

// IsZero returns true for invalid Bytes, for future omitempty support (Go 1.4?)
// A non-null Byte with a 0 value will not be considered zero.
func (b Byte) IsZero() bool {
	_ = "STUB: not implemented"

	// Equal returns true if both ints have the same value or are both null.
	return false
}

func (b Byte) Equal(other Byte) bool { _ = "STUB: not implemented"; return false }

func (b Byte) value() (int64, bool) { _ = "STUB: not implemented"; return 0, false }
