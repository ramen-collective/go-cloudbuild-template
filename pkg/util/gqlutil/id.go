package gqlutil

import (
	"fmt"
	"io"
	"strconv"
)

// ID present a model ID, which is exposed as a string (ID type) in GraphQL
// but is a uint64 internally
type ID uint64

// NewIDFromString instanciate a new ID from a given string
func NewIDFromString(v string) ID {
	var ID ID
	ID.UnmarshalGQL(v)
	return ID
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (id *ID) UnmarshalGQL(v interface{}) error {
	stringID, ok := v.(string)
	if !ok {
		return fmt.Errorf("ID must be a string")
	}
	intID, err := strconv.ParseUint(stringID, 10, 64)
	if err != nil {
		return BadUserInputErrorf("ID must be a numeric string")
	}
	*id = ID(intID)
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (id ID) MarshalGQL(w io.Writer) {
	fmt.Fprintf(w, `"%s"`, id.String())
}

// Int returns actual uint64 value of the ID
func (id ID) Int() uint64 {
	return uint64(id)
}

// String returns string value of the ID
func (id ID) String() string {
	return strconv.FormatUint(uint64(id), 10)
}
