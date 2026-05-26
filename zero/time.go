package zero

import (
	"database/sql"
	"database/sql/driver"
	"time"
)

// Time is a nullable time.Time.
// JSON marshals to the zero value for time.Time if null.
// Considered to be null to SQL if zero.
type Time struct {
	sql.NullTime
}

// Value implements the driver Valuer interface.
func (t Time) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

// NewTime creates a new Time.
func NewTime(t time.Time, valid bool) Time { _ = "STUB: not implemented"; return *new(Time) }

// TimeFrom creates a new Time that will
// be null if t is the zero value.
func TimeFrom(t time.Time) Time { _ = "STUB: not implemented"; return *new(Time) }

// TimeFromPtr creates a new Time that will
// be null if t is nil or *t is the zero value.
func TimeFromPtr(t *time.Time) Time { _ = "STUB: not implemented"; return *new(Time) }

// ValueOrZero returns the inner value if valid, otherwise zero.
func (t Time) ValueOrZero() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// ValueOr returns the inner value if valid, otherwise v.
func (t Time) ValueOr(v time.Time) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// MarshalJSON implements json.Marshaler.
// It will encode the zero value of time.Time
// if this time is invalid.
func (t Time) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON implements json.Unmarshaler.
// It supports string and null input.
func (t *Time) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalText implements encoding.TextMarshaler.
// It will encode to an empty time.Time if invalid.
func (t Time) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalText implements encoding.TextUnmarshaler.
// It has compatibility with the null package in that it will accept empty strings as invalid values,
// which will be unmarshaled to an invalid zero value.
func (t *Time) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"

	// allowing "null" is for backwards compatibility with v3
	return nil
}

// SetValid changes this Time's value and
// sets it to be non-null.
func (t *Time) SetValid(v time.Time) { _ = "STUB: not implemented"; return }

// Ptr returns a pointer to this Time's value,
// or a nil pointer if this Time is zero.
func (t Time) Ptr() *time.Time { _ = "STUB: not implemented"; return nil }

// IsZero returns true for null or zero Times, for potential future omitempty support.
func (t Time) IsZero() bool { _ = "STUB: not implemented"; return false }

// Equal returns true if both Time objects encode the same time or are both are either null or zero.
// Two times can be equal even if they are in different locations.
// For example, 6:00 +0200 CEST and 4:00 UTC are Equal.
func (t Time) Equal(other Time) bool { _ = "STUB: not implemented"; return false }

// ExactEqual returns true if both Time objects are equal or both are either null or zero.
// ExactEqual returns false for times that are in different locations or
// have a different monotonic clock reading.
func (t Time) ExactEqual(other Time) bool { _ = "STUB: not implemented"; return false }
