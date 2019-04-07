package main

import "testing"

func TestLongestAgedItemProcess(t *testing.T) {
	subject := "Aged Brue"
	tests := []struct {
		name            string
		subject         *Item
		expectedSellIn  int
		expectedQuality int
	}{
		{"quality increase every day by 1", &Item{subject, 15, 20}, 14, 21},
		// NOTICE: Spec: Once the sell by date has passed, quality degrades twice as fast.
		//         It uses word `degrades`, there are no words for Cheese case to `change`
		// TODO: Clarify specs with product
		{"quality increase twice on passed", &Item{subject, -15, 20}, -16, 22},
		{"quality could be increased up to 50", &Item{subject, -15, 49}, -16, 50},
		{"quality max 50", &Item{subject, -15, 50}, -16, 50},
		// NOTICE: Spec: The quality of an item is never more than 50.
		//         There is mention only for Legendary items to not be altered.
		//         There are no cases, what should we do if there is quality set higher than 50
		// TODO: Clarify specs with product
		{"quality should not change on passed and highr max", &Item{subject, -15, 60}, -16, 60},
		{"quality should not change and highr max", &Item{subject, 15, 60}, 14, 60},
	}
	for _, tc := range tests {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := LongestAgedItem{tc.subject}
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
