package null

import (
	"database/sql"
)

// Int32 is an nullable int32.
// It does not consider zero values to be null.
// It will decode to null, not zero, if null.
type Int32 struct {
	sql.NullInt32
}

// NewInt32 creates a new Int32.
func NewInt32(i int32, valid bool) Int32 { _ = "STUB: not implemented"; return *new(Int32) }

// Int32From creates a new Int32 that will always be valid.
func Int32From(i int32) Int32 {
	_ = "STUB: not implemented"
	return *

	// Int32FromPtr creates a new Int32 that be null if i is nil.
	new(Int32)
}

func Int32FromPtr(i *int32) Int32 { _ = "STUB: not implemented"; return *new(Int32) }

// ValueOrZero returns the inner value if valid, otherwise zero.
func (i Int32) ValueOrZero() int32 { _ = "STUB: not implemented"; return 0 }

// ValueOr returns the inner value if valid, otherwise v.
func (i Int32) ValueOr(v int32) int32 { _ = "STUB: not implemented"; return 0 }

// UnmarshalJSON implements json.Unmarshaler.
// It supports number, string, and null input.
// 0 will not be considered a null Int32.
func (i *Int32) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalText implements encoding.TextUnmarshaler.
// It will unmarshal to a null Int32 if the input is blank.
// It will return an error if the input is not an integer, blank, or "null".
func (i *Int32) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalJSON implements json.Marshaler.
// It will encode null if this Int32 is null.
func (i Int32) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalText implements encoding.TextMarshaler.
// It will encode a blank string if this Int32 is null.
func (i Int32) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// SetValid changes this Int32's value and also sets it to be non-null.
func (i *Int32) SetValid(n int32) { _ = "STUB: not implemented"; return }

// Ptr returns a pointer to this Int32's value, or a nil pointer if this Int32 is null.
func (i Int32) Ptr() *int32 { _ = "STUB: not implemented"; return nil }

// IsZero returns true for invalid Int32s, for future omitempty support (Go 1.4?)
// A non-null Int32 with a 0 value will not be considered zero.
func (i Int32) IsZero() bool {
	_ = "STUB: not implemented"

	// Equal returns true if both ints have the same value or are both null.
	return false
}

func (i Int32) Equal(other Int32) bool { _ = "STUB: not implemented"; return false }

func (i Int32) value() (int64, bool) { _ = "STUB: not implemented"; return 0, false }
