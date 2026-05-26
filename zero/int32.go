package zero

import (
	"database/sql"
)

// Int32 is a nullable int32.
// JSON marshals to zero if null.
// Considered null to SQL if zero.
type Int32 struct {
	sql.NullInt32
}

// NewInt32 creates a new Int32
func NewInt32(i int32, valid bool) Int32 { _ = "STUB: not implemented"; return *new(Int32) }

// Int32From creates a new Int32 that will be null if zero.
func Int32From(i int32) Int32 { _ = "STUB: not implemented"; return *new(Int32) }

// Int32FromPtr creates a new Int32 that be null if i is nil.
func Int32FromPtr(i *int32) Int32 { _ = "STUB: not implemented"; return *new(Int32) }

// ValueOrZero returns the inner value if valid, otherwise zero.
func (i Int32) ValueOrZero() int32 { _ = "STUB: not implemented"; return 0 }

// ValueOr returns the inner value if valid, otherwise v.
func (i Int32) ValueOr(v int32) int32 { _ = "STUB: not implemented"; return 0 }

// UnmarshalJSON implements json.Unmarshaler.
// It supports number and null input.
// 0 will be considered a null Int32.
func (i *Int32) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalText implements encoding.TextUnmarshaler.
// It will unmarshal to a null Int32 if the input is a blank, or zero.
// It will return an error if the input is not an integer, blank, or "null".
func (i *Int32) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalJSON implements json.Marshaler.
// It will encode 0 if this Int32 is null.
func (i Int32) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalText implements encoding.TextMarshaler.
// It will encode a zero if this Int32 is null.
func (i Int32) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// SetValid changes this Int32's value and also sets it to be non-null.
func (i *Int32) SetValid(n int32) { _ = "STUB: not implemented"; return }

// Ptr returns a pointer to this Int32's value, or a nil pointer if this Int32 is null.
func (i Int32) Ptr() *int32 { _ = "STUB: not implemented"; return nil }

// IsZero returns true for null or zero Int32s, for future omitempty support (Go 1.4?)
func (i Int32) IsZero() bool { _ = "STUB: not implemented"; return false }

// Equal returns true if both ints have the same value or are both either null or zero.
func (i Int32) Equal(other Int32) bool { _ = "STUB: not implemented"; return false }

func (i Int32) value() (int64, bool) { _ = "STUB: not implemented"; return 0, false }
