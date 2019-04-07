package main

// Item is represent inventory to be selling.
// NOTICE: do not alter the item function
// Validation Specs:
//   - All items have a sell-in value which denotes the number of days we have to sell the item
//   - All items have a quality value which denotes how valuable the item is
//   - The quality of an item is never negative
//   - The quality of an item is never more than 50
//   - "Sulfuras" is a legendary item and as such its quality is 80 and it never alters
type Item struct {
	name            string
	sellIn, quality int
}

var items = []Item{
	Item{"+5 Dexterity Vest", 10, 20},
	Item{"Aged Brie", 2, 0},
	Item{"Elixir of the Mongoose", 5, 7},
	Item{"Sulfuras, Hand of Ragnaros", 0, 80},
	Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
	Item{"Conjured Mana Cake", 3, 6},
}
