package main

import "testing"

func TestConjuredItemProcess(t *testing.T) {
	tests := []struct {
		name            string
		subject         *Item
		expectedSellIn  int
		expectedQuality int
	}{
		{"quality is over maximum", &Item{"Foo", 10, 80}, 9, 78},
		{"quality decrease every day by 2", &Item{"Quux", 15, 20}, 14, 18},
		{"quality decrease every day by 2 even quality is higher maximum", &Item{"Quux", 15, 60}, 14, 58},
		{"quality decrease faster on sell passed", &Item{"Quux", -15, 20}, -16, 16},
		{"quality could be zero", &Item{"Quux", 15, 2}, 14, 0},
		{"quality could be zero after passed", &Item{"Quux", -15, 4}, -16, 0},
		{"quality is not negative", &Item{"Quux", 15, 0}, 14, 0},
		{"quality is not negative even for passed", &Item{"Quux", -15, 0}, -16, 0},
	}
	for _, tc := range tests {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := ConjuredItem{tc.subject}
			actual.Process()

			if tc.expectedSellIn != actual.sellIn {
				t.Errorf("sellIn %v; want %v", actual.sellIn, tc.expectedSellIn)
			}

			if tc.expectedQuality != actual.quality {
				t.Errorf("quality %v; want %v", actual.quality, tc.expectedQuality)
			}
		})
	}
}
