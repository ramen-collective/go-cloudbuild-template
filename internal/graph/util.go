package graph

import "github.com/ramen-collective/go-cloudbuild-template/pkg/util/gqlutil"

func ConvertArrayUUIDToString(uuids []gqlutil.UUID) []string {
	var arr []string
	for _, val := range uuids {
		arr = append(arr, val.String())
	}
	return arr
}
