package gotils

import (
	"testing"
)

var input = "aaron barratt"
var brownFox = "The quick brown fox jumps over the lazy dog"
var allCaps = "MF DOOM"

func TestToTitleCase(t *testing.T) {
	res := ToTitleCase(input)
	if res != "Aaron Barratt" {
		t.Errorf("to title case failed: %s", res)
	}
}

func TestToTitleCaseBrowFox(t *testing.T) {
	res := ToTitleCase(brownFox)
	if res != "The Quick Brown Fox Jumps Over The Lazy Dog" {
		t.Errorf("to title case failed: %s", res)
	}
}

func TestToTitleCaseAllCaps(t *testing.T) {
	res := ToTitleCase(allCaps)
	if res != "Mf Doom" {
		t.Errorf("to title case failed: %s", res)
	}
}
