package null

import (
	"database/sql"
	"database/sql/driver"
	"time"
)

// Time is a nullable time.Time. It supports SQL and JSON serialization.
// It will marshal to null if null.
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

// TimeFrom creates a new Time that will always be valid.
func TimeFrom(t time.Time) Time {
	_ = "STUB: not implemented"
	return *

	// TimeFromPtr creates a new Time that will be null if t is nil.
	new(Time)
}

func TimeFromPtr(t *time.Time) Time { _ = "STUB: not implemented"; return *new(Time) }

// ValueOrZero returns the inner value if valid, otherwise zero.
func (t Time) ValueOrZero() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// ValueOr returns the inner value if valid, otherwise v.
func (t Time) ValueOr(v time.Time) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// MarshalJSON implements json.Marshaler.
// It will encode null if this time is null.
func (t Time) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON implements json.Unmarshaler.
// It supports string and null input.
func (t *Time) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalText implements encoding.TextMarshaler.
// It returns an empty string if invalid, otherwise time.Time's MarshalText.
func (t Time) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalText implements encoding.TextUnmarshaler.
// It has backwards compatibility with v3 in that the string "null" is considered equivalent to an empty string
// and unmarshaling will succeed. This may be removed in a future version.
func (t *Time) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"

	// allowing "null" is for backwards compatibility with v3
	return nil
}

// SetValid changes this Time's value and sets it to be non-null.
func (t *Time) SetValid(v time.Time) { _ = "STUB: not implemented"; return }

// Ptr returns a pointer to this Time's value, or a nil pointer if this Time is null.
func (t Time) Ptr() *time.Time { _ = "STUB: not implemented"; return nil }

// IsZero returns true for invalid Times, hopefully for future omitempty support.
// A non-null Time with a zero value will not be considered zero.
func (t Time) IsZero() bool {
	_ = "STUB: not implemented"

	// Equal returns true if both Time objects encode the same time or are both null.
	// Two times can be equal even if they are in different locations.
	// For example, 6:00 +0200 CEST and 4:00 UTC are Equal.
	return false
}

func (t Time) Equal(other Time) bool { _ = "STUB: not implemented"; return false }

// ExactEqual returns true if both Time objects are equal or both null.
// ExactEqual returns false for times that are in different locations or
// have a different monotonic clock reading.
func (t Time) ExactEqual(other Time) bool { _ = "STUB: not implemented"; return false }
