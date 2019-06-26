package main

type PromotedItem struct {
	*Item
}

// Process is change quality depends on sell in
func (item *PromotedItem) Process() {
	if item.quality < MAX_QUALITY {
		item.IncrQuality()

		if item.sellIn < 11 {
			item.IncrQuality()
		}

		if item.sellIn < 6 {
			item.IncrQuality()
		}

		if item.quality > MAX_QUALITY {
			item.quality = MAX_QUALITY
		}
	}

	item.sellIn--

	if item.sellIn < 0 {
		item.Reset()
	}
}
