package gnull

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Null[T any] struct {
	Value T
	Valid bool // Valid is true if Value is not NULL
}

func New[T any](value T, valid bool) Null[T] {
	return Null[T]{
		Value: value,
		Valid: true,
	}
}

func NewFromPtr[T any](value *T) Null[T] {
	if value == nil {
		return Null[T]{
			Valid: false,
		}
	}

	return New(*value, true)
}

// nullBytes is a JSON null literal
var nullBytes = []byte("null")

// UnmarshalJSON implements json.Unmarshaler.
func (n *Null[T]) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, nullBytes) {
		n.Valid = false
		return nil
	}

	if err := json.Unmarshal(data, &n.Value); err != nil {
		return fmt.Errorf("gnull: couldn't unmarshal JSON: %w", err)
	}

	n.Valid = true
	return nil
}

// MarshalJSON implements json.Marshaler.
func (n *Null[T]) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(n.Value)
}
