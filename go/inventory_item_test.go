package main

import "testing"

func TestInventoryItemLongAged(t *testing.T) {
	tests := []struct {
		name     string
		subject  *InventoryItem
		expected bool
	}{
		{"Aged Brie", &InventoryItem{&Item{name: "Aged Brie"}}, true},
		{"General", &InventoryItem{&Item{name: "+5 Dexterity Vest"}}, false},
	}
	for _, tc := range tests {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := tc.subject.IsLongestAged()
			if tc.expected != actual {
				t.Errorf("IsLongestAged() == %v; want %v", actual, tc.expected)
			}
		})
	}
}

func TestInventoryItemIsPromoted(t *testing.T) {
	tests := []struct {
		name     string
		subject  *InventoryItem
		expected bool
	}{
		{"Cheese", &InventoryItem{&Item{name: "Aged Brie"}}, false},
		{"Promoted", &InventoryItem{&Item{name: "Backstage passes to a TAFKAL80ETC concert"}}, true},
		{"General", &InventoryItem{&Item{name: "+5 Dexterity Vest"}}, false},
	}
	for _, tc := range tests {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := tc.subject.IsPromoted()
			if tc.expected != actual {
				t.Errorf("IsPromoted() == %v; want %v", actual, tc.expected)
			}
		})
	}
}

func TestInventoryItemIsLegendary(t *testing.T) {
	tests := []struct {
		name     string
		subject  *InventoryItem
		expected bool
	}{
		{"Cheese", &InventoryItem{&Item{name: "Aged Brie"}}, false},
		{"Promoted", &InventoryItem{&Item{name: "Backstage passes to a TAFKAL80ETC concert"}}, false},
		{"Legendary", &InventoryItem{&Item{name: "Sulfuras, Hand of Ragnaros"}}, true},
		{"General", &InventoryItem{&Item{name: "+5 Dexterity Vest"}}, false},
	}
	for _, tc := range tests {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := tc.subject.IsLegendary()
			if tc.expected != actual {
				t.Errorf("IsLegendary() == %v; want %v", actual, tc.expected)
			}
		})
	}
}
