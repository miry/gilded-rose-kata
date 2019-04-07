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
	if !item.IsLongestAged() && !item.IsPromoted() {
		if item.quality > 0 {
			if !item.IsLegendary() {
				item.quality--
			}
		}
	} else {
		if item.quality < 50 {
			item.quality++
			if item.IsPromoted() {
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

	if !item.IsLegendary() {
		item.sellIn--
	}

	if item.sellIn >= 0 {
		return item
	}

	if !item.IsLongestAged() {
		if !item.IsPromoted() {
			if item.quality > 0 {
				if !item.IsLegendary() {
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

	return item
}
