package main

// InventoryItem extends Item with extra functionality
type InventoryItem struct {
	*Item
}

func NewInventoryItem(item *Item) *InventoryItem {
	return &InventoryItem{item}
}
