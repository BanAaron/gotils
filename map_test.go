package gotils

import (
	"reflect"
	"testing"
)

func TestOrderedKeysEmptyMap(t *testing.T) {
	var emptyMap map[int]string
	keys := OrderedKeys(emptyMap, false)
	var target []int
	if !reflect.DeepEqual(keys, target) {
		t.Errorf("empty map failed: %v", keys)
	}
}

func TestOrderedKeysInt(t *testing.T) {
	mapInt := map[int]string{
		4: "",
		1: "",
		3: "",
		2: "",
	}
	target := []int{1, 2, 3, 4}
	keys := OrderedKeys(mapInt, false)
	if !reflect.DeepEqual(keys, target) {
		t.Errorf("int map failed: %v", keys)
	}
}

func TestOrderedKeysNegativeInt(t *testing.T) {
	mapInt := map[int]string{
		-1:  "",
		2:   "",
		5:   "",
		-10: "",
	}
	target := []int{-10, -1, 2, 5}
	keys := OrderedKeys(mapInt, false)
	if !reflect.DeepEqual(keys, target) {
		t.Errorf("int map failed: %v", keys)
	}
}

func TestOrderedKeysString(t *testing.T) {
	mapInt := map[string]int{
		"chris": 10,
		"aaron": 20,
		"jamie": 30,
		"drew":  40,
	}
	target := []string{"aaron", "chris", "drew", "jamie"}
	keys := OrderedKeys(mapInt, false)
	if !reflect.DeepEqual(keys, target) {
		t.Errorf("int map failed: %v", keys)
	}
}

func TestOrderedKeysFloat(t *testing.T) {
	mapInt := map[float64]string{
		3.14:   "pi",
		420.69: "lol",
		-55.55: "fives",
	}
	target := []float64{-55.55, 3.14, 420.69}
	keys := OrderedKeys(mapInt, false)
	if !reflect.DeepEqual(keys, target) {
		t.Errorf("int map failed: %v", keys)
	}
}

func TestOrderedKeysReverse(t *testing.T) {
	mapInt := map[int]string{
		5: "nice",
		4: "frog",
		3: "bro",
	}
	target := []int{5, 4, 3}
	keys := OrderedKeys(mapInt, true)
	if !reflect.DeepEqual(keys, target) {
		t.Errorf("reverse test failed: %v", keys)
	}
}
