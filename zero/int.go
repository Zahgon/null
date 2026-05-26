package zero

import (
	"database/sql"
)

// Int is a nullable int64.
// JSON marshals to zero if null.
// Considered null to SQL if zero.
type Int struct {
	sql.NullInt64
}

// Int64 is an alias for Int.
type Int64 = Int

// NewInt creates a new Int
func NewInt(i int64, valid bool) Int { _ = "STUB: not implemented"; return *new(Int) }

// IntFrom creates a new Int that will be null if zero.
func IntFrom(i int64) Int {
	_ = "STUB: not implemented"
	return *

	// IntFromPtr creates a new Int that be null if i is nil.
	new(Int)
}

func IntFromPtr(i *int64) Int { _ = "STUB: not implemented"; return *new(Int) }

// ValueOrZero returns the inner value if valid, otherwise zero.
func (i Int) ValueOrZero() int64 { _ = "STUB: not implemented"; return 0 }

// ValueOr returns the inner value if valid, otherwise v.
func (i Int) ValueOr(v int64) int64 { _ = "STUB: not implemented"; return 0 }

// UnmarshalJSON implements json.Unmarshaler.
// It supports number and null input.
// 0 will be considered a null Int.
func (i *Int) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalText implements encoding.TextUnmarshaler.
// It will unmarshal to a null Int if the input is a blank, or zero.
// It will return an error if the input is not an integer, blank, or "null".
func (i *Int) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalJSON implements json.Marshaler.
// It will encode 0 if this Int is null.
func (i Int) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalText implements encoding.TextMarshaler.
// It will encode a zero if this Int is null.
func (i Int) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// SetValid changes this Int's value and also sets it to be non-null.
func (i *Int) SetValid(n int64) { _ = "STUB: not implemented"; return }

// Ptr returns a pointer to this Int's value, or a nil pointer if this Int is null.
func (i Int) Ptr() *int64 { _ = "STUB: not implemented"; return nil }

// IsZero returns true for null or zero Ints, for future omitempty support (Go 1.4?)
func (i Int) IsZero() bool { _ = "STUB: not implemented"; return false }

// Equal returns true if both ints have the same value or are both either null or zero.
func (i Int) Equal(other Int) bool { _ = "STUB: not implemented"; return false }

func (i Int) value() (int64, bool) { _ = "STUB: not implemented"; return 0, false }
