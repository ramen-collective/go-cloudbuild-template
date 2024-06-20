package gqlutil

import (
	"fmt"
	"io"
)

// UUID present a model UUID, which is exposed as a string (UUID type) in GraphQL
// but is a uint64 internally
type UUID string

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (uuid *UUID) UnmarshalGQL(v interface{}) error {
	stringUUID, ok := v.(string)
	if !ok {
		return fmt.Errorf("ID must be a string")
	}
	*uuid = UUID(stringUUID)
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (uuid UUID) MarshalGQL(w io.Writer) {
	fmt.Fprintf(w, `"%s"`, uuid.String())
}

// String returns string value of the UUID
func (uuid UUID) String() string {
	return string(uuid)
}
