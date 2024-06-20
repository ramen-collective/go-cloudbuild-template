package gqlutil

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// DecodeIDCursor decodes a cursor only
// supposed to contain a numeric ID
func DecodeIDCursor(cursor *string) (*uint64, error) {
	if cursor == nil || *cursor == "" {
		return nil, nil
	}
	decodedString, err := base64.StdEncoding.DecodeString(*cursor)
	if err != nil {
		return nil, err
	}
	intCursor, err := strconv.ParseUint(string(decodedString), 10, 64)
	return &intCursor, err
}

// DecodeUUIDCursor decodes a cursor only
// supposed to contain a numeric UUID
func DecodeUUIDCursor(cursor *string) (*string, error) {
	if cursor == nil {
		return nil, nil
	}
	decodedString, err := base64.StdEncoding.DecodeString(*cursor)
	if err != nil {
		return nil, err
	}
	strCursor := string(decodedString)
	return &strCursor, err
}

// EncodeIDCursor encodes a cursor from a given ID
func EncodeIDCursor(ID uint64) string {
	return base64.StdEncoding.EncodeToString([]byte(strconv.FormatUint(ID, 10)))
}

// EncodeUUIDCursor encodes a cursor from a given UUID
func EncodeUUIDCursor(UUID string) string {
	return base64.StdEncoding.EncodeToString([]byte(UUID))
}

// DecodeSeedAndIDCursor decodes a cursor
// supposed to contain a seed and a numeric ID
func DecodeSeedAndIDCursor(cursor *string) (*int64, *uint64, error) {
	if cursor == nil || *cursor == "" {
		return nil, nil, nil
	}
	decodedString, err := base64.StdEncoding.DecodeString(*cursor)
	if err != nil {
		return nil, nil, err
	}
	splitedString := strings.Split(string(decodedString), "_")
	if len(splitedString) != 2 {
		return nil, nil, errors.New("Error parsing the cursor")
	}
	seed, err := strconv.ParseInt(splitedString[0], 10, 64)
	ID, err := strconv.ParseUint(splitedString[1], 10, 64)
	return &seed, &ID, err
}

// EncodeSeedAndIDCursor encodes a cursor from a given
// seed and ID
func EncodeSeedAndIDCursor(seed int64, ID uint64) string {
	stringCursor := fmt.Sprintf("%s_%s", strconv.FormatInt(seed, 10), strconv.FormatUint(ID, 10))
	return base64.StdEncoding.EncodeToString([]byte(stringCursor))
}

// EncodeSeedAndUUIDCursor encodes a cursor from a given
// seed and UUID
func EncodeSeedAndUUIDCursor(seed int64, UUID string) string {
	stringCursor := fmt.Sprintf("%s_%s", strconv.FormatInt(seed, 10), UUID)
	return base64.StdEncoding.EncodeToString([]byte(stringCursor))
}
