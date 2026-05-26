package null

import (
	"database/sql"
)

// Float is a nullable float64.
// It does not consider zero values to be null.
// It will decode to null, not zero, if null.
type Float struct {
	sql.NullFloat64
}

// NewFloat creates a new Float
func NewFloat(f float64, valid bool) Float { _ = "STUB: not implemented"; return *new(Float) }

// FloatFrom creates a new Float that will always be valid.
func FloatFrom(f float64) Float {
	_ = "STUB: not implemented"
	return *

	// FloatFromPtr creates a new Float that be null if f is nil.
	new(Float)
}

func FloatFromPtr(f *float64) Float { _ = "STUB: not implemented"; return *new(Float) }

// ValueOrZero returns the inner value if valid, otherwise zero.
func (f Float) ValueOrZero() float64 { _ = "STUB: not implemented"; return 0 }

// ValueOr returns the inner value if valid, otherwise v.
func (f Float) ValueOr(v float64) float64 { _ = "STUB: not implemented"; return 0 }

// UnmarshalJSON implements json.Unmarshaler.
// It supports number and null input.
// 0 will not be considered a null Float.
func (f *Float) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalText implements encoding.TextUnmarshaler.
// It will unmarshal to a null Float if the input is blank.
// It will return an error if the input is not an integer, blank, or "null".
func (f *Float) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalJSON implements json.Marshaler.
// It will encode null if this Float is null.
func (f Float) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalText implements encoding.TextMarshaler.
// It will encode a blank string if this Float is null.
func (f Float) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// SetValid changes this Float's value and also sets it to be non-null.
func (f *Float) SetValid(n float64) { _ = "STUB: not implemented"; return }

// Ptr returns a pointer to this Float's value, or a nil pointer if this Float is null.
func (f Float) Ptr() *float64 { _ = "STUB: not implemented"; return nil }

// IsZero returns true for invalid Floats, for future omitempty support (Go 1.4?)
// A non-null Float with a 0 value will not be considered zero.
func (f Float) IsZero() bool {
	_ = "STUB: not implemented"

	// Equal returns true if both floats have the same value or are both null.
	// Warning: calculations using floating point numbers can result in different ways
	// the numbers are stored in memory. Therefore, this function is not suitable to
	// compare the result of a calculation. Use this method only to check if the value
	// has changed in comparison to some previous value.
	return false
}

func (f Float) Equal(other Float) bool { _ = "STUB: not implemented"; return false }
