package stax

import (
	"testing"
)

func TestStax(t *testing.T) {
	items := []Item{
		Item{},
		Item{},
		Item{},
	}
	_ = NewStack(items)
}
