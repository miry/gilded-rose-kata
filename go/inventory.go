package main

// StockTaking or "inventory checking" is the physical verification of the quantities and condition of items
// held in an inventory or warehouse
func StockTaking(items []*Item) []*Item {
	for _, item := range items {
		NewInventoryItem(item).Process()
	}

	return items
}
