package main

import "testing"

func TestLegendaryItemProcess(t *testing.T) {
	tests := []struct {
		name            string
		subject         *LegendaryItem
		expectedSellIn  int
		expectedQuality int
	}{
		{"quality is over maximum", &LegendaryItem{&Item{name: "Sulfrus", sellIn: 10, quality: 80}}, 10, 80},
		{"sell in soon", &LegendaryItem{&Item{name: "Sulfrus", sellIn: 1, quality: 50}}, 1, 50},
		{"sell in passed", &LegendaryItem{&Item{name: "Sulfrus", sellIn: -2, quality: 20}}, -2, 20},
	}
	for _, tc := range tests {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc.subject.Process()
			actual := tc.subject
			if tc.expectedSellIn != actual.sellIn {
				t.Errorf("sellIn %v; want %v", actual.sellIn, tc.expectedSellIn)
			}
			if tc.expectedQuality != actual.quality {
				t.Errorf("quality %v; want %v", actual.quality, tc.expectedQuality)
			}
		})
	}
}
