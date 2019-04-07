package main

type empty struct{}
type set map[string]empty

const (
	MAX_QUALITY = 50

	// Item Names
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

func (item *InventoryItem) Process() {
	if item.IsLegendary() {
		return
	}

	// Modify quality
	if !item.IsLongestAged() && !item.IsPromoted() {
		// Decrement quality
		if item.quality > 0 {
			item.quality--
		}
	} else {
		// Increment quality
		if item.quality < MAX_QUALITY {
			item.quality++
			if item.IsPromoted() {

				if item.sellIn < 11 {
					item.quality++
				}

				if item.sellIn < 6 {
					item.quality++
				}

			}

			if item.quality > MAX_QUALITY {
				item.quality = MAX_QUALITY
			}

		}
	}

	// Decrement Sell In
	item.sellIn--

	if item.sellIn >= 0 {
		return
	}

	if item.IsPromoted() {
		item.quality = 0
		return
	}

	// Sell In Passed Section
	if item.IsLongestAged() {
		if item.quality < MAX_QUALITY {
			item.quality++
		}

		return
	}

	if item.quality > 0 {
		item.quality--
	}
}
