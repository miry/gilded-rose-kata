package main

import "testing"

func TestItemValidate(t *testing.T) {
	t.Skip("not implemented yet")
	tests := []struct {
		name    string
		subject Item
		expect  bool
	}{
		{"empty object", Item{}, false},
		{"name value is nil", Item{}, false},
		{"name value is foo", Item{}, true},
		{"sell-in value is nil", Item{}, false},
		{"sell-in value is negative", Item{}, true},
		{"quality is nil", Item{}, false},
		{"quality is negative", Item{}, false},
		{"quality is more than 50", Item{}, false},
		{"sulfuras quality is 80", Item{}, true},
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
