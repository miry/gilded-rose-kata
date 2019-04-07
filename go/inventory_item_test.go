package main

import (
	"fmt"
	"testing"
)

func TestNewInventory(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{"Aged Brie", "*main.LongestAgedItem"},
		{"+5 Dexterity Vest", "*main.GeneralItem"},
		{"Backstage passes to a TAFKAL80ETC concert", "*main.PromotedItem"},
		{"Sulfuras, Hand of Ragnaros", "*main.LegendaryItem"},
	}
	for _, tc := range tests {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := fmt.Sprintf("%T", NewInventoryItem(&Item{name: tc.name}))
			if tc.expected != actual {
				t.Errorf("Expected type %s, got %s", tc.expected, actual)
			}
		})
	}
}
