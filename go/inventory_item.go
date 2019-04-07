package main

type empty struct{}
type set map[string]empty

const (
	AGED_BRIE        = "Aged Brie"
	BACKSTAGE_PASSES = "Backstage passes to a TAFKAL80ETC concert"
	SULFURAS         = "Sulfuras, Hand of Ragnaros"
)

var (
	LONGEST_AGED = set{AGED_BRIE: empty{}}
	PROMOTED     = set{BACKSTAGE_PASSES: empty{}}
	LEGENDARY    = set{SULFURAS: empty{}}
)

// InventoryItem extends Item with extra functionality
type InventoryItem struct {
	*Item
}

func NewInventoryItem(item *Item) *InventoryItem {
	return &InventoryItem{item}
}

func (i *InventoryItem) IsLongestAged() bool {
	_, ok := LONGEST_AGED[i.name]
	return ok
}

func (i *InventoryItem) IsPromoted() bool {
	_, ok := PROMOTED[i.name]
	return ok
}

func (i *InventoryItem) IsLegendary() bool {
	_, ok := LEGENDARY[i.name]
	return ok
}
