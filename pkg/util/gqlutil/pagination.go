package gqlutil

import (
	"fmt"
)

// PageInfo represents the PageInfo type in GraphQL
type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor"`
	EndCursor       *string `json:"endCursor"`
}

// PaginateIDs returns a specific page of a slice of IDs with a corresponding
// new PageInfo.
// Deprecated (use PaginateItems instead).
func PaginateIDs(IDs []uint64, first int, cursor *uint64, seed *int64) ([]uint64, *PageInfo, error) {
	return PaginateIDEdges(IDs, nil, cursor, &first, nil, seed)
}

// PaginateIDEdges returns a specific page of a slice of IDs with a corresponding
// new PageInfo.
// Implements the relay cursor connection spec:
// https://relay.dev/graphql/connections.htm
func PaginateIDEdges(allEdges []uint64, before *uint64, after *uint64, first *int, last *int, seed *int64) ([]uint64, *PageInfo, error) {
	hasNextPage := false
	hasPreviousPage := false

	edges, hasPrevious, hasNext := applyCursorsToEdges(allEdges, before, after)

	if first != nil {
		hasNextPage = len(edges) > *first
	} else if before != nil {
		hasNextPage = hasNext
	}

	if last != nil {
		hasPreviousPage = len(edges) > *last
	} else if after != nil {
		hasPreviousPage = hasPrevious
	}

	if len(edges) == 0 {
		return edges, &PageInfo{
			StartCursor:     nil,
			EndCursor:       nil,
			HasPreviousPage: hasPreviousPage,
			HasNextPage:     hasNextPage,
		}, nil
	}

	if first != nil {
		if len(edges)-*first > 0 {
			edges = edges[:*first]
		}
	}

	if last != nil {
		min := len(edges) - *last
		if min > 0 {
			edges = edges[min:]
		}
	}

	var startCursor, endCursor string
	if seed != nil {
		startCursor = EncodeSeedAndIDCursor(*seed, edges[0])
		endCursor = EncodeSeedAndIDCursor(*seed, edges[len(edges)-1])
	} else {
		startCursor = EncodeIDCursor(edges[0])
		endCursor = EncodeIDCursor(edges[len(edges)-1])
	}

	return edges, &PageInfo{
		StartCursor:     &startCursor,
		EndCursor:       &endCursor,
		HasPreviousPage: hasPreviousPage,
		HasNextPage:     hasNextPage,
	}, nil
}

// ValidatePaginationArguments validates pagination arguments.
func ValidatePaginationArguments(before *uint64, after *uint64, first *int, last *int) error {
	if first != nil && last != nil {
		return fmt.Errorf("first and last arguments are mutually exclusive")
	}
	if after != nil && before != nil {
		return fmt.Errorf("after and before arguments are mutually exclusive")
	}
	if first != nil && *first <= 0 {
		return fmt.Errorf("first should be a positive integer")
	}
	if last != nil && *last <= 0 {
		return fmt.Errorf("last should be a positive integer")
	}
	return nil
}

func applyCursorsToEdges(allEdges []uint64, before *uint64, after *uint64) ([]uint64, bool, bool) {
	startIndex := 0
	hasPreviousEdges := false
	hasNextEdges := false
	if after != nil {
		for i, edge := range allEdges {
			if edge == *after {
				startIndex = i + 1
				break
			}
		}
		hasPreviousEdges = startIndex > 1
	}
	edges := allEdges[startIndex:]
	startIndex = 0
	endIndex := len(edges)
	if before != nil {
		for i, edge := range edges {
			if edge == *before {
				endIndex = i
				break
			}
		}
		hasNextEdges = len(edges) > endIndex+1
	}
	return edges[:endIndex], hasPreviousEdges, hasNextEdges
}

// PaginateUUIDs returns a specific page of a slice of UUIDs with a corresponding
// new PageInfo
func PaginateUUIDs(UUIDs []string, first int, cursor *string, seed *int64) ([]string, *PageInfo, error) {
	var paginatedUUIDs []string
	hasNextPage := false
	hasPreviousPage := false
	startIndex := 0

	if cursor != nil {
		for i, ID := range UUIDs {
			if ID == *cursor {
				startIndex = i + 1
				break
			}
		}
	}

	endIndex := startIndex + first
	if endIndex > len(UUIDs) {
		endIndex = len(UUIDs)
	}

	paginatedUUIDs = UUIDs[startIndex:endIndex]
	hasPreviousPage = startIndex > 0

	if len(paginatedUUIDs) == 0 {
		return paginatedUUIDs, &PageInfo{
			StartCursor:     nil,
			EndCursor:       nil,
			HasPreviousPage: hasPreviousPage,
			HasNextPage:     false,
		}, nil
	}

	hasNextPage = paginatedUUIDs[len(paginatedUUIDs)-1] != UUIDs[len(UUIDs)-1]

	var startCursor, endCursor string
	if seed != nil {
		startCursor = EncodeSeedAndUUIDCursor(*seed, paginatedUUIDs[0])
		endCursor = EncodeSeedAndUUIDCursor(*seed, paginatedUUIDs[len(paginatedUUIDs)-1])
	} else {
		startCursor = EncodeUUIDCursor(paginatedUUIDs[0])
		endCursor = EncodeUUIDCursor(paginatedUUIDs[len(paginatedUUIDs)-1])
	}

	return paginatedUUIDs, &PageInfo{
		StartCursor:     &startCursor,
		EndCursor:       &endCursor,
		HasPreviousPage: hasPreviousPage,
		HasNextPage:     hasNextPage,
	}, nil
}
