package gqlutil

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testIDCursor = "ODc="
var testIDCursorValue = uint64(87)

var testSeed = int64(1641229116556671158)
var testSeedAndIDCursor = "MTY0MTIyOTExNjU1NjY3MTE1OF8xOTg="
var testSeedAndIDCursorValue = uint64(198)

func TestDecodeIDCursor(t *testing.T) {
	require := require.New(t)
	cursor, err := DecodeIDCursor(&testIDCursor)

	require.NoError(err)
	require.Exactly(testIDCursorValue, *cursor)
}

func TestDecodeIDCursorEmpty(t *testing.T) {
	require := require.New(t)
	cursor, err := DecodeIDCursor(nil)

	require.NoError(err)
	require.Nil(cursor)
}

func TestEncodeIDCursor(t *testing.T) {
	require := require.New(t)
	cursor := EncodeIDCursor(testIDCursorValue)
	require.Exactly(testIDCursor, cursor)
}

func TestDecodeSeedAndIDCursor(t *testing.T) {
	require := require.New(t)
	seed, cursor, err := DecodeSeedAndIDCursor(&testSeedAndIDCursor)
	require.NoError(err)
	require.Exactly(testSeed, *seed)
	require.Exactly(testSeedAndIDCursorValue, *cursor)
}

func TestDecodeSeedAndIDCursorEmpty(t *testing.T) {
	require := require.New(t)
	seed, cursor, err := DecodeSeedAndIDCursor(nil)
	require.NoError(err)
	require.Nil(seed)
	require.Nil(cursor)
}

func TestEncodeSeedAndIDCursor(t *testing.T) {
	require := require.New(t)
	cursor := EncodeSeedAndIDCursor(testSeed, testSeedAndIDCursorValue)
	require.Exactly(testSeedAndIDCursor, cursor)
}
