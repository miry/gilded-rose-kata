package main

import (
	"testing"
)

func TestStockTakingNoPanic(t *testing.T) {
	items := []Item{
		Item{"+5 Dexterity Vest", 10, 20},
		Item{"Aged Brie", 2, 0},
		Item{"Elixir of the Mongoose", 5, 7},
		Item{"Sulfuras, Hand of Ragnaros", 0, 80},
		Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		Item{"Conjured Mana Cake", 3, 6},
	}
	StockTaking(items)
}

func TestStockTakingEmptyItems(t *testing.T) {
	subject := []Item{}
	actual := StockTaking(subject)

	if actual == nil {
		t.Errorf("actual should not be nil")
	}

	if len(actual) != 0 {
		t.Errorf("actual should be empty, instead of %v", len(actual))
	}
}

func TestStockTakingLowersSellAndQuality(t *testing.T) {
	subject := []Item{
		Item{"+5 Dexterity Vest", 10, 20},
	}

	actual := StockTaking(subject)

	if len(actual) != 1 {
		t.Errorf("actual should have 1 ite, instead of %v", len(actual))
	}

	if actual[0].sellIn != 9 {
		t.Errorf("item's sell in should be decreased by 1")
	}

	if actual[0].quality != 19 {
		t.Errorf("item's quality should be decreased by 1")
	}
}

func TestStockTakingItemQualityNeverNegative(t *testing.T) {
	subject := []Item{
		Item{"+5 Dexterity Vest", 10, 0},
	}

	actual := StockTaking(subject)[0].quality
	expected := 0

	if actual != expected {
		t.Errorf("The quality of an item is never negative. Got %v", actual)
	}
}

func TestStockTakingSellPassedQualityTwiceDecrease(t *testing.T) {
	subject := []Item{
		Item{"+5 Dexterity Vest", -1, 50},
	}

	actual := StockTaking(subject)[0].quality
	expected := 48

	if actual != expected {
		t.Errorf("Expect to decreased by 2 and be %v. Got %v", expected, actual)
	}
}

func TestStockTakingCaseAgedBrieIncrQuality(t *testing.T) {
	subject := []Item{
		Item{"Aged Brie", 10, 30},
	}

	actual := StockTaking(subject)[0].quality
	expected := 31

	if actual != expected {
		t.Errorf("Expect to increase by 1 for Aged Brie and be %v. Got %v", expected, actual)
	}
}

func TestStockTakingCaseAgedBrieDecrSellIn(t *testing.T) {
	subject := []Item{
		Item{"Aged Brie", 10, 30},
	}

	actual := StockTaking(subject)[0].sellIn
	expected := 9

	if actual != expected {
		t.Errorf("Expect to decrease by 1 sell in for Aged Brie and be %v. Got %v", expected, actual)
	}
}

func TestStockTakingCaseQualityMax50(t *testing.T) {
	subject := []Item{
		Item{"Aged Brie", 10, 50},
	}

	actual := StockTaking(subject)[0].quality
	expected := 50

	if actual != expected {
		t.Errorf("Expect to have max %v quality. Got %v", expected, actual)
	}
}

func TestStockTakingLegendaryItemQualitySame(t *testing.T) {
	subject := []Item{
		Item{"Sulfuras, Hand of Ragnaros", 10, 80},
	}

	actual := StockTaking(subject)[0].quality
	expected := 80

	if actual != expected {
		t.Errorf("Expect to not change quality and be %v. Got %v", expected, actual)
	}
}

func TestStockTakingLegendaryItemSellInSame(t *testing.T) {
	subject := []Item{
		Item{"Sulfuras, Hand of Ragnaros", 10, 80},
	}

	actual := StockTaking(subject)[0].sellIn
	expected := 10

	if actual != expected {
		t.Errorf("Expect to not change sell day and be %v. Got %v", expected, actual)
	}
}

func TestStockTakinBackstagePassesDecrSellIn(t *testing.T) {
	subject := []Item{
		Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
	}

	actual := StockTaking(subject)[0].sellIn
	expected := 14

	if actual != expected {
		t.Errorf("Expect to not change sell day and be %v. Got %v", expected, actual)
	}
}

func TestStockTakinBackstagePassesQuality(t *testing.T) {
	backstage := "Backstage passes to a TAFKAL80ETC concert"
	tests := []struct {
		name     string
		item     Item
		expected int
	}{
		{"have a lot of time increase by 1", Item{backstage, 15, 20}, 21},
		{"concert in 11 days increse by 1", Item{backstage, 11, 20}, 21},
		{"concert in 10 days increse by 2", Item{backstage, 10, 20}, 22},
		{"concert in 6 days increse by 2", Item{backstage, 6, 20}, 22},
		{"concert in 5 days increse by 3", Item{backstage, 5, 20}, 23},
		{"concert pass drop to 0", Item{backstage, -1, 20}, 0},
	}

	for _, tc := range tests {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := StockTaking([]Item{tc.item})[0].quality
			if tc.expected != actual {
				t.Errorf("Expect quality to be %v for sell in %v. Got %v", tc.expected, tc.item.sellIn, actual)
			}
		})
	}
}
