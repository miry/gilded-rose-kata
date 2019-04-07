package main

// StockTaking or "inventory checking" is the physical verification of the quantities and condition of items
// held in an inventory or warehouse
func StockTaking(items []*Item) []*Item {
	for _, item := range items {
		processItem(NewInventoryItem(item))
	}

	return items
}

func processItem(item *InventoryItem) *InventoryItem {
	if item.name != "Aged Brie" && item.name != "Backstage passes to a TAFKAL80ETC concert" {
		if item.quality > 0 {
			if item.name != "Sulfuras, Hand of Ragnaros" {
				item.quality--
			}
		}
	} else {
		if item.quality < 50 {
			item.quality++
			if item.name == "Backstage passes to a TAFKAL80ETC concert" {
				if item.sellIn < 11 {
					if item.quality < 50 {
						item.quality++
					}
				}
				if item.sellIn < 6 {
					if item.quality < 50 {
						item.quality++
					}
				}
			}
		}
	}

	if item.name != "Sulfuras, Hand of Ragnaros" {
		item.sellIn--
	}

	if item.sellIn < 0 {
		if item.name != "Aged Brie" {
			if item.name != "Backstage passes to a TAFKAL80ETC concert" {
				if item.quality > 0 {
					if item.name != "Sulfuras, Hand of Ragnaros" {
						item.quality--
					}
				}
			} else {
				item.quality = 0
			}
		} else {
			if item.quality < 50 {
				item.quality++
			}
		}
	}
	return item
}
