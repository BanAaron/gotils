package gotils

import (
	"cmp"
	"sort"
)

type ordered interface {
	cmp.Ordered
}

func OrderedKeys[Key ordered, Value any](input map[Key]Value) (keys []Key) {
	for k := range input {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	return
}
