package gotils

import (
	"testing"
)

func TestLeftPad(t *testing.T) {
	result := LeftPad("aaron", 4)
	if result != "    aaron" {
		t.Errorf("Standard test failed: %s", result)
	}
}

func TestLeftPadTwoSpaces(t *testing.T) {
	result := LeftPad("aaron", 2)
	if result != "  aaron" {
		t.Errorf("Two spaces test failed: %s", result)
	}
}

func TestLeftPadZeroSpaces(t *testing.T) {
	result := LeftPad("aaron", 0)
	if result != "aaron" {
		t.Errorf("Zero spaces test failed: %s", result)
	}
}

func TestLeftPadOnlySpaces(t *testing.T) {
	result := LeftPad(" ", 4)
	if result != "     " {
		t.Errorf("Only spaces test failed: %s", result)
	}
}

func TestLeftPadLeadingSpaces(t *testing.T) {
	result := LeftPad(" aaron", 4)
	if result != "     aaron" {
		t.Errorf("Leading spaces test failed: %s", result)
	}
}

func TestLeftPadTrailingSpaces(t *testing.T) {
	result := LeftPad("aaron ", 4)
	if result != "    aaron " {
		t.Errorf("Trailing spaces test failed: %s", result)
	}
}

func TestLeftPadEmptyString(t *testing.T) {
	var foo string
	result := LeftPad(foo, 4)
	if result != "    " {
		t.Errorf("Empty string test failed: %s", result)
	}
}
