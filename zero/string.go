// Package zero contains SQL types that consider zero input and null input to be equivalent
// with convenient support for JSON and text marshaling.
// Types in this package will JSON marshal to their zero value, even if null.
// Use the null parent package if you don't want this.
package zero

import (
	"database/sql"
)

// String is a nullable string.
// JSON marshals to a blank string if null.
// Considered null to SQL if zero.
type String struct {
	sql.NullString
}

// NewString creates a new String
func NewString(s string, valid bool) String { _ = "STUB: not implemented"; return *new(String) }

// StringFrom creates a new String that will be null if s is blank.
func StringFrom(s string) String { _ = "STUB: not implemented"; return *new(String) }

// StringFromPtr creates a new String that be null if s is nil or blank.
// It will make s point to the String's value.
func StringFromPtr(s *string) String { _ = "STUB: not implemented"; return *new(String) }

// ValueOrZero returns the inner value if valid, otherwise zero.
func (s String) ValueOrZero() string { _ = "STUB: not implemented"; return "" }

// ValueOr returns the inner value if valid, otherwise v.
func (s String) ValueOr(v string) string { _ = "STUB: not implemented"; return "" }

// UnmarshalJSON implements json.Unmarshaler.
// It supports string and null input. Blank string input produces a null String.
func (s *String) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

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

// IsZero returns true for null or empty strings, for potential future omitempty support.
func (s String) IsZero() bool { _ = "STUB: not implemented"; return false }

// Equal returns true if both strings have the same value or are both either null or empty.
func (s String) Equal(other String) bool { _ = "STUB: not implemented"; return false }
