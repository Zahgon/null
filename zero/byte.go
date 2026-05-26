package zero

import (
	"database/sql"
)

// Byte is a nullable byte.
// JSON marshals to zero if null.
// Considered null to SQL if zero.
type Byte struct {
	sql.NullByte
}

// NewByte creates a new Byte
func NewByte(i byte, valid bool) Byte { _ = "STUB: not implemented"; return *new(Byte) }

// ByteFrom creates a new Byte that will be null if zero.
func ByteFrom(i byte) Byte {
	_ = "STUB: not implemented"
	return *

	// ByteFromPtr creates a new Byte that be null if i is nil.
	new(Byte)
}

func ByteFromPtr(i *byte) Byte { _ = "STUB: not implemented"; return *new(Byte) }

// ValueOrZero returns the inner value if valid, otherwise zero.
func (b Byte) ValueOrZero() byte { _ = "STUB: not implemented"; return 0 }

// ValueOr returns the inner value if valid, otherwise v.
func (b Byte) ValueOr(v byte) byte { _ = "STUB: not implemented"; return 0 }

// UnmarshalJSON implements json.Unmarshaler.
// It supports number and null input.
// 0 will be considered a null Byte.
func (b *Byte) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalText implements encoding.TextUnmarshaler.
// It will unmarshal to a null Byte if the input is a blank, or zero.
// It will return an error if the input is not an integer, blank, or "null".
func (b *Byte) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalJSON implements json.Marshaler.
// It will encode 0 if this Byte is null.
func (b Byte) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalText implements encoding.TextMarshaler.
// It will encode a zero if this Byte is null.
func (b Byte) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// SetValid changes this Byte's value and also sets it to be non-null.
func (b *Byte) SetValid(n byte) { _ = "STUB: not implemented"; return }

// Ptr returns a pointer to this Byte's value, or a nil pointer if this Byte is null.
func (b Byte) Ptr() *byte { _ = "STUB: not implemented"; return nil }

// IsZero returns true for null or zero Bytes, for future omitempty support (Go 1.4?)
func (b Byte) IsZero() bool { _ = "STUB: not implemented"; return false }

// Equal returns true if both ints have the same value or are both either null or zero.
func (b Byte) Equal(other Byte) bool { _ = "STUB: not implemented"; return false }

func (b Byte) value() (int64, bool) { _ = "STUB: not implemented"; return 0, false }
