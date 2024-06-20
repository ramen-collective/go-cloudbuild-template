package gqlutil

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var paginationTestIDs = []uint64{1, 33, 5, 62, 855, 21, 89, 72}

func TestPaginateIDsWithoutCursor(t *testing.T) {
	require := require.New(t)
	first := 5
	expectedIDs := paginationTestIDs[:first]
	startCursor := EncodeIDCursor(expectedIDs[0])
	endCursor := EncodeIDCursor(expectedIDs[len(expectedIDs)-1])
	paginatedIDs, pageInfo, err := PaginateIDs(paginationTestIDs, first, nil, nil)
	require.Nil(err)
	require.Exactly(expectedIDs, paginatedIDs)
	require.Exactly(&PageInfo{
		StartCursor:     &startCursor,
		EndCursor:       &endCursor,
		HasPreviousPage: false,
		HasNextPage:     true,
	}, pageInfo)
}

func TestPaginateIDsWithCursorAndSeed(t *testing.T) {
	require := require.New(t)
	first := 5
	cursor := uint64(5)
	seed := int64(123456789)
	expectedIDs := paginationTestIDs[3 : 3+first]
	startCursor := EncodeSeedAndIDCursor(seed, expectedIDs[0])
	endCursor := EncodeSeedAndIDCursor(seed, expectedIDs[len(expectedIDs)-1])
	paginatedIDs, pageInfo, err := PaginateIDs(paginationTestIDs, first, &cursor, &seed)
	require.Nil(err)
	require.Exactly(expectedIDs, paginatedIDs)
	require.Exactly(&PageInfo{
		StartCursor:     &startCursor,
		EndCursor:       &endCursor,
		HasPreviousPage: true,
		HasNextPage:     false,
	}, pageInfo)
}

func TestPaginateIDsFullResults(t *testing.T) {
	require := require.New(t)
	first := 8
	expectedIDs := paginationTestIDs[:first]
	startCursor := EncodeIDCursor(expectedIDs[0])
	endCursor := EncodeIDCursor(expectedIDs[len(expectedIDs)-1])
	paginatedIDs, pageInfo, err := PaginateIDs(paginationTestIDs, first, nil, nil)
	require.Nil(err)
	require.Exactly(expectedIDs, paginatedIDs)
	require.Exactly(&PageInfo{
		StartCursor:     &startCursor,
		EndCursor:       &endCursor,
		HasPreviousPage: false,
		HasNextPage:     false,
	}, pageInfo)
}

func TestPaginateEdgesWithLastUnderLimitOnly(t *testing.T) {
	require := require.New(t)
	last := 7
	expectedIDs := paginationTestIDs[len(paginationTestIDs)-last:]
	startCursor := EncodeIDCursor(expectedIDs[0])
	endCursor := EncodeIDCursor(expectedIDs[len(expectedIDs)-1])
	paginatedIDs, pageInfo, err := PaginateIDEdges(paginationTestIDs, nil, nil, nil, &last, nil)
	require.Nil(err)
	require.Exactly(expectedIDs, paginatedIDs)
	require.Exactly(&PageInfo{
		StartCursor:     &startCursor,
		EndCursor:       &endCursor,
		HasPreviousPage: true,
		HasNextPage:     false,
	}, pageInfo)
}

func TestPaginateEdgesWithLastAboveLimitOnly(t *testing.T) {
	require := require.New(t)
	last := 12
	expectedIDs := paginationTestIDs
	startCursor := EncodeIDCursor(expectedIDs[0])
	endCursor := EncodeIDCursor(expectedIDs[len(expectedIDs)-1])
	paginatedIDs, pageInfo, err := PaginateIDEdges(paginationTestIDs, nil, nil, nil, &last, nil)
	require.Nil(err)
	require.Exactly(expectedIDs, paginatedIDs)
	require.Exactly(&PageInfo{
		StartCursor:     &startCursor,
		EndCursor:       &endCursor,
		HasPreviousPage: false,
		HasNextPage:     false,
	}, pageInfo)
}

func TestPaginateEdgesWithFirstUnderLimitOnly(t *testing.T) {
	require := require.New(t)
	first := 7
	expectedIDs := paginationTestIDs[:first]
	startCursor := EncodeIDCursor(expectedIDs[0])
	endCursor := EncodeIDCursor(expectedIDs[len(expectedIDs)-1])
	paginatedIDs, pageInfo, err := PaginateIDEdges(paginationTestIDs, nil, nil, &first, nil, nil)
	require.Nil(err)
	require.Exactly(expectedIDs, paginatedIDs)
	require.Exactly(&PageInfo{
		StartCursor:     &startCursor,
		EndCursor:       &endCursor,
		HasPreviousPage: false,
		HasNextPage:     true,
	}, pageInfo)
}

func TestPaginateEdgesWithFirstAboveLimitOnly(t *testing.T) {
	require := require.New(t)
	first := 15
	expectedIDs := paginationTestIDs
	startCursor := EncodeIDCursor(expectedIDs[0])
	endCursor := EncodeIDCursor(expectedIDs[len(expectedIDs)-1])
	paginatedIDs, pageInfo, err := PaginateIDEdges(paginationTestIDs, nil, nil, &first, nil, nil)
	require.Nil(err)
	require.Exactly(expectedIDs, paginatedIDs)
	require.Exactly(&PageInfo{
		StartCursor:     &startCursor,
		EndCursor:       &endCursor,
		HasPreviousPage: false,
		HasNextPage:     false,
	}, pageInfo)
}

func TestPaginateEdgesWithAfterAndFirstOnly(t *testing.T) {
	require := require.New(t)
	first := 3
	expectedIDs := paginationTestIDs[4:7]
	startCursor := EncodeIDCursor(expectedIDs[0])
	endCursor := EncodeIDCursor(expectedIDs[len(expectedIDs)-1])
	paginatedIDs, pageInfo, err := PaginateIDEdges(paginationTestIDs, nil, &paginationTestIDs[3], &first, nil, nil)
	require.Nil(err)
	require.Exactly(expectedIDs, paginatedIDs)
	require.Exactly(&PageInfo{
		StartCursor:     &startCursor,
		EndCursor:       &endCursor,
		HasPreviousPage: true,
		HasNextPage:     true,
	}, pageInfo)
}

func TestPaginateEdgesWithBeforeAndLastOnly(t *testing.T) {
	require := require.New(t)
	last := 2
	expectedIDs := paginationTestIDs[1:3]
	startCursor := EncodeIDCursor(expectedIDs[0])
	endCursor := EncodeIDCursor(expectedIDs[len(expectedIDs)-1])
	paginatedIDs, pageInfo, err := PaginateIDEdges(paginationTestIDs, &paginationTestIDs[3], nil, nil, &last, nil)
	require.Nil(err)
	require.Exactly(expectedIDs, paginatedIDs)
	require.Exactly(&PageInfo{
		StartCursor:     &startCursor,
		EndCursor:       &endCursor,
		HasPreviousPage: true,
		HasNextPage:     true,
	}, pageInfo)
}

func TestPaginateEdgesWithAllParameters(t *testing.T) {
	require := require.New(t)
	last := 2
	expectedIDs := paginationTestIDs[4:6]
	startCursor := EncodeIDCursor(expectedIDs[0])
	endCursor := EncodeIDCursor(expectedIDs[len(expectedIDs)-1])
	paginatedIDs, pageInfo, err := PaginateIDEdges(paginationTestIDs, &paginationTestIDs[6], &paginationTestIDs[0], nil, &last, nil)
	require.Nil(err)
	require.Exactly(expectedIDs, paginatedIDs)
	require.Exactly(&PageInfo{
		StartCursor:     &startCursor,
		EndCursor:       &endCursor,
		HasPreviousPage: true,
		HasNextPage:     true,
	}, pageInfo)
}

func TestValidatePaginationArguments(t *testing.T) {
	require := require.New(t)
	before := uint64(2)
	after := uint64(3)
	first := 10
	last := 10

	err := ValidatePaginationArguments(&before, &after, nil, nil)
	require.Error(err)

	err = ValidatePaginationArguments(&before, nil, nil, nil)
	require.NoError(err)

	err = ValidatePaginationArguments(nil, &after, nil, nil)
	require.NoError(err)

	err = ValidatePaginationArguments(nil, nil, &first, &last)
	require.Error(err)

	err = ValidatePaginationArguments(nil, nil, &first, nil)
	require.NoError(err)

	err = ValidatePaginationArguments(nil, nil, nil, &last)
	require.NoError(err)

	first = 0
	last = 0
	err = ValidatePaginationArguments(nil, nil, &first, nil)
	require.Error(err)

	err = ValidatePaginationArguments(nil, nil, nil, &last)
	require.Error(err)
}
