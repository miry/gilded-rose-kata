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
