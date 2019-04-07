package main

import "fmt"

func main() {
	items := []*Item{
		&Item{"+5 Dexterity Vest", 10, 20},
		&Item{"Aged Brie", 2, 0},
		&Item{"Elixir of the Mongoose", 5, 7},
		&Item{"Sulfuras, Hand of Ragnaros", 0, 80},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		&Item{"Conjured Mana Cake", 3, 6},
	}

	StockTaking(items)

	for _, item := range items {
		fmt.Printf("- %s { sell in %d , quality: %d }\n", item.name, item.sellIn, item.quality)
	}
}
