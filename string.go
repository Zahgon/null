// Package null contains SQL types that consider zero input and null input as separate values,
// with convenient support for JSON and text marshaling.
// Types in this package will always encode to their null value if null.
// Use the zero subpackage if you want zero values and null to be treated the same.
package null

import (
	"database/sql"
)

// String is a nullable string. It supports SQL and JSON serialization.
// It will marshal to null if null. Blank string input will be considered null.
type String struct {
	sql.NullString
}

// StringFrom creates a new String that will never be blank.
func StringFrom(s string) String {
	_ = "STUB: not implemented"
	return *

	// StringFromPtr creates a new String that be null if s is nil.
	new(String)
}

func StringFromPtr(s *string) String { _ = "STUB: not implemented"; return *new(String) }

// ValueOrZero returns the inner value if valid, otherwise zero.
func (s String) ValueOrZero() string { _ = "STUB: not implemented"; return "" }

// ValueOr returns the inner value if valid, otherwise v.
func (s String) ValueOr(v string) string { _ = "STUB: not implemented"; return "" }

// NewString creates a new String
func NewString(s string, valid bool) String { _ = "STUB: not implemented"; return *new(String) }

// UnmarshalJSON implements json.Unmarshaler.
// It supports string and null input. Blank string input does not produce a null String.
func (s *String) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalJSON implements json.Marshaler.
// It will encode null if this String is null.
func (s String) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalText implements encoding.TextMarshaler.
// It will encode a blank string when this String is null.
func (s String) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalText implements encoding.TextUnmarshaler.
// It will unmarshal to a null String if the input is a blank string.
func (s *String) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

// SetValid changes this String's value and also sets it to be non-null.
func (s *String) SetValid(v string) { _ = "STUB: not implemented"; return }

// Ptr returns a pointer to this String's value, or a nil pointer if this String is null.
func (s String) Ptr() *string { _ = "STUB: not implemented"; return nil }

// IsZero returns true for null strings, for potential future omitempty support.
func (s String) IsZero() bool {
	_ = "STUB: not implemented"

	// Equal returns true if both strings have the same value or are both null.
	return false
}

func (s String) Equal(other String) bool { _ = "STUB: not implemented"; return false }
