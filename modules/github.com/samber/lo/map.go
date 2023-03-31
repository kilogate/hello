package main

import (
	"fmt"

	"github.com/samber/lo"
)

func main() {
	m := map[string]string{
		"a": "A",
		"b": "B",
		"c": "C",
	}

	// Keys & Values & Entries & ToPairs & FromEntries & FromPairs
	keys := lo.Keys(m)
	values := lo.Values(m)
	entries := lo.Entries(m)
	pairs := lo.ToPairs(m)
	nm := lo.FromEntries(entries)
	nm = lo.FromPairs(pairs)
	fmt.Println(keys, values, entries, pairs, nm)

	// ValueOr
	valueOr := lo.ValueOr(m, "d", "D")
	fmt.Println(valueOr)

	// PickBy & PickByKeys & PickByValues
	pickBy := lo.PickBy(m, func(key string, value string) bool {
		return key == "a" && value == "A"
	})
	pickByKeys := lo.PickByKeys(m, []string{"a", "b"})
	pickByValues := lo.PickByValues(m, []string{"A", "B"})
	fmt.Println(pickBy, pickByKeys, pickByValues)

	// OmitBy & OmitByKeys & OmitByValues
	omitBy := lo.OmitBy(m, func(key string, value string) bool {
		return key == "a" && value == "A"
	})
	omitByKeys := lo.OmitByKeys(m, []string{"a", "b"})
	omitByValues := lo.OmitByValues(m, []string{"A", "B"})
	fmt.Println(omitBy, omitByKeys, omitByValues)

	// Invert
	invert := lo.Invert(m)
	fmt.Println(invert)

	// Assign
	otherMap := map[string]string{
		"a": "AA",
		"b": "BB",
		"c": "CC",
	}
	assign := lo.Assign(m, otherMap)
	fmt.Println(assign)

	// MapKeys & MapValues & MapEntries & MapToSlice
	mapKeys := lo.MapKeys(m, func(value string, key string) string {
		return key + value
	})
	mapValues := lo.MapValues(m, func(value string, key string) string {
		return key + value
	})
	mapEntries := lo.MapEntries(m, func(key string, value string) (string, string) {
		return "Key:" + key, "Value:" + value
	})
	mapToSlice := lo.MapToSlice(m, func(key string, value string) string {
		return key + value
	})
	fmt.Println(mapKeys, mapValues, mapEntries, mapToSlice)
}
