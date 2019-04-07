package main

import "testing"

func TestPromotedItemProcess(t *testing.T) {
	tests := []struct {
		name            string
		subject         *PromotedItem
		expectedSellIn  int
		expectedQuality int
	}{
		{"quality is over maximum", &PromotedItem{&Item{name: "Backstage passes", sellIn: 10, quality: 80}}, 9, 80},
		{"quality reset on sell pass", &PromotedItem{&Item{name: "Backstage passes", sellIn: 0, quality: 80}}, -1, 0},
		{"quality reset on sell pass", &PromotedItem{&Item{name: "Backstage passes", sellIn: -5, quality: 30}}, -6, 0},
		{"quality close to maximum", &PromotedItem{&Item{name: "Backstage passes", sellIn: 5, quality: 49}}, 4, 50},
		{"quality close to maximum close to sell", &PromotedItem{&Item{name: "Backstage passes", sellIn: 5, quality: 48}}, 4, 50},
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
