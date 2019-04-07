package main

type empty struct{}
type set map[string]empty

const (
	MAX_QUALITY = 50

	// Item Names
	AGED_BRIE          = "Aged Brie"
	BACKSTAGE_PASSES   = "Backstage passes to a TAFKAL80ETC concert"
	SULFURAS           = "Sulfuras, Hand of Ragnaros"
	CONJURED_MANA_CAKE = "Conjured Mana Cake"
)

var (
	LONGEST_AGED = set{AGED_BRIE: empty{}}
	PROMOTED     = set{BACKSTAGE_PASSES: empty{}}
	LEGENDARY    = set{SULFURAS: empty{}}
	CONJURED     = set{CONJURED_MANA_CAKE: empty{}}
)

// InventoryItem extends Item with extra functionality
type InventoryItem interface {
	Process()
}

func NewInventoryItem(item *Item) InventoryItem {
	if _, ok := LEGENDARY[item.name]; ok {
		return &LegendaryItem{item}
	}

	if _, ok := PROMOTED[item.name]; ok {
		return &PromotedItem{item}
	}

	if _, ok := LONGEST_AGED[item.name]; ok {
		return &LongestAgedItem{item}
	}

	if _, ok := CONJURED[item.name]; ok {
		return &ConjuredItem{item}
	}

	return &GeneralItem{item}
}
