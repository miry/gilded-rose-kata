package main

const (
	AGED_BRIE = "Aged Brie"
)

// InventoryItem extends Item with extra functionality
type InventoryItem struct {
	*Item
}

func NewInventoryItem(item *Item) *InventoryItem {
	return &InventoryItem{item}
}

func (i *InventoryItem) IsLongestAged() bool {
	return i.name == AGED_BRIE
}
