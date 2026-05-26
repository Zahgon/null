package zero

import (
	"database/sql"
)

// Bool is a nullable bool. False input is considered null.
// JSON marshals to false if null.
// Considered null to SQL unmarshaled from a false value.
type Bool struct {
	sql.NullBool
}

// NewBool creates a new Bool
func NewBool(b bool, valid bool) Bool { _ = "STUB: not implemented"; return *new(Bool) }

// BoolFrom creates a new Bool that will be null if false.
func BoolFrom(b bool) Bool {
	_ = "STUB: not implemented"
	return *

	// BoolFromPtr creates a new Bool that be null if b is nil.
	new(Bool)
}

func BoolFromPtr(b *bool) Bool { _ = "STUB: not implemented"; return *new(Bool) }

// ValueOrZero returns the inner value if valid, otherwise false.
func (b Bool) ValueOrZero() bool { _ = "STUB: not implemented"; return false }

// ValueOr returns the inner value if valid, otherwise v.
func (b Bool) ValueOr(v bool) bool { _ = "STUB: not implemented"; return false }

// UnmarshalJSON implements json.Unmarshaler.
// "false" will be considered a null Bool.
func (b *Bool) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalText implements encoding.TextUnmarshaler.
// It will unmarshal to a null Bool if the input is false or blank.
// It will return an error if the input is not a float, blank, or "null".
func (b *Bool) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalJSON implements json.Marshaler.
// It will encode null if this Bool is null.
func (b Bool) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalText implements encoding.TextMarshaler.
// It will encode a zero if this Bool is null.
func (b Bool) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// SetValid changes this Bool's value and also sets it to be non-null.
func (b *Bool) SetValid(v bool) { _ = "STUB: not implemented"; return }

// Ptr returns a poBooler to this Bool's value, or a nil poBooler if this Bool is null.
func (b Bool) Ptr() *bool { _ = "STUB: not implemented"; return nil }

// IsZero returns true for null or zero Bools, for future omitempty support (Go 1.4?)
func (b Bool) IsZero() bool { _ = "STUB: not implemented"; return false }

// Equal returns true if both booleans are true and valid, or if both booleans are either false or invalid.
func (b Bool) Equal(other Bool) bool { _ = "STUB: not implemented"; return false }
