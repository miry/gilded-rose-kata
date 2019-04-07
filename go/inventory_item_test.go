package main

import "testing"

func TestInventoryItemValidate(t *testing.T) {
	t.Skip("not implemented yet")
	tests := []struct {
		name    string
		subject InventoryItem
		expect  bool
	}{
		{"empty object", InventoryItem{}, false},
		{"name value is nil", InventoryItem{}, false},
		{"name value is foo", InventoryItem{}, true},
		{"sell-in value is nil", InventoryItem{}, false},
		{"sell-in value is negative", InventoryItem{}, true},
		{"quality is nil", InventoryItem{}, false},
		{"quality is negative", InventoryItem{}, false},
		{"quality is more than 50", InventoryItem{}, false},
		{"sulfuras quality is 80", InventoryItem{}, true},
	}
	for _, tc := range tests {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := false
			if tc.expect != actual {
				t.Errorf("Validate() == %v; want %v", actual, tc.expect)
			}
		})
	}
}

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
