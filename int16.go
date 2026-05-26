package null

import (
	"database/sql"
)

// Int16 is an nullable int16.
// It does not consider zero values to be null.
// It will decode to null, not zero, if null.
type Int16 struct {
	sql.NullInt16
}

// NewInt16 creates a new Int16.
func NewInt16(i int16, valid bool) Int16 { _ = "STUB: not implemented"; return *new(Int16) }

// Int16From creates a new Int16 that will always be valid.
func Int16From(i int16) Int16 {
	_ = "STUB: not implemented"
	return *

	// Int16FromPtr creates a new Int16 that be null if i is nil.
	new(Int16)
}

func Int16FromPtr(i *int16) Int16 { _ = "STUB: not implemented"; return *new(Int16) }

// ValueOrZero returns the inner value if valid, otherwise zero.
func (i Int16) ValueOrZero() int16 { _ = "STUB: not implemented"; return 0 }

// ValueOr returns the inner value if valid, otherwise v.
func (i Int16) ValueOr(v int16) int16 { _ = "STUB: not implemented"; return 0 }

// UnmarshalJSON implements json.Unmarshaler.
// It supports number, string, and null input.
// 0 will not be considered a null Int16.
func (i *Int16) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalText implements encoding.TextUnmarshaler.
// It will unmarshal to a null Int16 if the input is blank.
// It will return an error if the input is not an integer, blank, or "null".
func (i *Int16) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalJSON implements json.Marshaler.
// It will encode null if this Int16 is null.
func (i Int16) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalText implements encoding.TextMarshaler.
// It will encode a blank string if this Int16 is null.
func (i Int16) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// SetValid changes this Int16's value and also sets it to be non-null.
func (i *Int16) SetValid(n int16) { _ = "STUB: not implemented"; return }

// Ptr returns a pointer to this Int16's value, or a nil pointer if this Int16 is null.
func (i Int16) Ptr() *int16 { _ = "STUB: not implemented"; return nil }

// IsZero returns true for invalid Int16s, for future omitempty support (Go 1.4?)
// A non-null Int16 with a 0 value will not be considered zero.
func (i Int16) IsZero() bool {
	_ = "STUB: not implemented"

	// Equal returns true if both ints have the same value or are both null.
	return false
}

func (i Int16) Equal(other Int16) bool { _ = "STUB: not implemented"; return false }

func (i Int16) value() (int64, bool) { _ = "STUB: not implemented"; return 0, false }
