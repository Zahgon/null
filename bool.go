package null

import (
	"database/sql"
)

// Bool is a nullable bool.
// It does not consider false values to be null.
// It will decode to null, not false, if null.
type Bool struct {
	sql.NullBool
}

// NewBool creates a new Bool
func NewBool(b bool, valid bool) Bool { _ = "STUB: not implemented"; return *new(Bool) }

// BoolFrom creates a new Bool that will always be valid.
func BoolFrom(b bool) Bool {
	_ = "STUB: not implemented"
	return *

	// BoolFromPtr creates a new Bool that will be null if f is nil.
	new(Bool)
}

func BoolFromPtr(b *bool) Bool { _ = "STUB: not implemented"; return *new(Bool) }

// ValueOrZero returns the inner value if valid, otherwise false.
func (b Bool) ValueOrZero() bool { _ = "STUB: not implemented"; return false }

// ValueOr returns the inner value if valid, otherwise v.
func (b Bool) ValueOr(v bool) bool { _ = "STUB: not implemented"; return false }

// UnmarshalJSON implements json.Unmarshaler.
// It supports number and null input.
// 0 will not be considered a null Bool.
func (b *Bool) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalText implements encoding.TextUnmarshaler.
// It will unmarshal to a null Bool if the input is blank.
// It will return an error if the input is not an integer, blank, or "null".
func (b *Bool) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalJSON implements json.Marshaler.
// It will encode null if this Bool is null.
func (b Bool) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalText implements encoding.TextMarshaler.
// It will encode a blank string if this Bool is null.
func (b Bool) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// SetValid changes this Bool's value and also sets it to be non-null.
func (b *Bool) SetValid(v bool) { _ = "STUB: not implemented"; return }

// Ptr returns a pointer to this Bool's value, or a nil pointer if this Bool is null.
func (b Bool) Ptr() *bool { _ = "STUB: not implemented"; return nil }

// IsZero returns true for invalid Bools, for future omitempty support (Go 1.4?)
// A non-null Bool with a 0 value will not be considered zero.
func (b Bool) IsZero() bool {
	_ = "STUB: not implemented"

	// Equal returns true if both booleans have the same value or are both null.
	return false
}

func (b Bool) Equal(other Bool) bool { _ = "STUB: not implemented"; return false }
