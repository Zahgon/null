package zero

import (
	"database/sql"
)

// Int16 is a nullable int16.
// JSON marshals to zero if null.
// Considered null to SQL if zero.
type Int16 struct {
	sql.NullInt16
}

// NewInt16 creates a new Int16
func NewInt16(i int16, valid bool) Int16 { _ = "STUB: not implemented"; return *new(Int16) }

// Int16From creates a new Int16 that will be null if zero.
func Int16From(i int16) Int16 { _ = "STUB: not implemented"; return *new(Int16) }

// Int16FromPtr creates a new Int16 that be null if i is nil.
func Int16FromPtr(i *int16) Int16 { _ = "STUB: not implemented"; return *new(Int16) }

// ValueOrZero returns the inner value if valid, otherwise zero.
func (i Int16) ValueOrZero() int16 { _ = "STUB: not implemented"; return 0 }

// ValueOr returns the inner value if valid, otherwise v.
func (i Int16) ValueOr(v int16) int16 { _ = "STUB: not implemented"; return 0 }

// UnmarshalJSON implements json.Unmarshaler.
// It supports number and null input.
// 0 will be considered a null Int16.
func (i *Int16) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalText implements encoding.TextUnmarshaler.
// It will unmarshal to a null Int16 if the input is a blank, or zero.
// It will return an error if the input is not an integer, blank, or "null".
func (i *Int16) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalJSON implements json.Marshaler.
// It will encode 0 if this Int16 is null.
func (i Int16) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalText implements encoding.TextMarshaler.
// It will encode a zero if this Int16 is null.
func (i Int16) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// SetValid changes this Int16's value and also sets it to be non-null.
func (i *Int16) SetValid(n int16) { _ = "STUB: not implemented"; return }

// Ptr returns a pointer to this Int16's value, or a nil pointer if this Int16 is null.
func (i Int16) Ptr() *int16 { _ = "STUB: not implemented"; return nil }

// IsZero returns true for null or zero Int16s, for future omitempty support (Go 1.4?)
func (i Int16) IsZero() bool { _ = "STUB: not implemented"; return false }

// Equal returns true if both ints have the same value or are both either null or zero.
func (i Int16) Equal(other Int16) bool { _ = "STUB: not implemented"; return false }

func (i Int16) value() (int64, bool) { _ = "STUB: not implemented"; return 0, false }
