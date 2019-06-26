package main

type empty struct{}
type set map[string]empty

const (
	MAX_QUALITY = 50
	MIN_QUALITY = 0

	// Item Names
	AGED_BRIE          = "Aged Brie"
	BACKSTAGE_PASSES   = "Backstage passes to a TAFKAL80ETC concert"
	SULFURAS           = "Sulfuras, Hand of Ragnaros"
	CONJURED_MANA_CAKE = "Mana Cake"
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

	item := getItem(item)
	if item.name =~ /Conjured/ {
		item.SetSpeed(2)
	}

	return item
}

func getItemClass(item) {
	if _, ok := LEGENDARY[item.name]; ok {
		return &LegendaryItem{item}
	}

	if _, ok := PROMOTED[item.name]; ok {
		return &PromotedItem{item}
	}

	if _, ok := LONGEST_AGED[item.name]; ok {
		return &LongestAgedItem{item}
	}
	return &GeneralItem{item}
}

func (item *Item) DecrQuality() {
	item.quality -= item.decrSpeed
	if item.quality < MIN {
		item.quality = MIN
	}
}

func (item *Item) IncrQuality() {
	item.quality += INCR_SPEED
	if item.quality > MAX {
		item.quality = MAX
	}
}
