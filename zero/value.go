//go:build go1.22

package zero

import (
	"database/sql"
)

type Value[T comparable] struct {
	sql.Null[T]
}

// NewValue creates a new Value.
func NewValue[T comparable](t T, valid bool) Value[T] { _ = "STUB: not implemented"; return nil }

// ValueFrom creates a new Value that will always be valid.
func ValueFrom[T comparable](t T) Value[T] { _ = "STUB: not implemented"; return nil }

// ValueFromPtr creates a new Value that will be null if t is nil.
func ValueFromPtr[T comparable](t *T) Value[T] { _ = "STUB: not implemented"; return nil }

// ValueOrZero returns the inner value if valid, otherwise zero.
func (t Value[T]) ValueOrZero() T { _ = "STUB: not implemented"; return *new(T) }

// ValueOr returns the inner value if valid, otherwise v.
func (t Value[T]) ValueOr(v T) T { _ = "STUB: not implemented"; return *new(T) }

// MarshalJSON implements json.Marshaler.
// It will encode null if this value is null or zero.
func (t Value[T]) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON implements json.Unmarshaler.
// It supports string and null input.
func (t *Value[T]) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// SetValid changes this Value's value and sets it to be non-null.
func (t *Value[T]) SetValid(v T) { _ = "STUB: not implemented"; return }

// Ptr returns a pointer to this Value's value, or a nil pointer if this Value is null.
func (t Value[T]) Ptr() *T { _ = "STUB: not implemented"; return nil }

// IsZero returns true for invalid or zero Values, hopefully for future omitempty support.
func (t Value[T]) IsZero() bool { _ = "STUB: not implemented"; return false }

// Equal returns true if both Value objects encode the same value or are both null.
func (t Value[T]) Equal(other Value[T]) bool { _ = "STUB: not implemented"; return false }
