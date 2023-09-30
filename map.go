package gotils

import (
	"cmp"
	"sort"
)

// ordered is an interface that says "whatever datatype implements this interface"
// must be of the type cmp.Ordered.
type ordered interface {
	cmp.Ordered
}

// OrderedKeys takes an input map as an argument and returns a slice of keys in
// ascending order.
func OrderedKeys[Key ordered, Value any](input map[Key]Value) (keys []Key) {
	for k := range input {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	return
}
