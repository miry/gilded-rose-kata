package main

type PromotedItem struct {
	*Item
}

func (item *PromotedItem) Process() {
	if item.quality < MAX_QUALITY {
		item.quality++

		if item.sellIn < 11 {
			item.quality++
		}

		if item.sellIn < 6 {
			item.quality++
		}

		if item.quality > MAX_QUALITY {
			item.quality = MAX_QUALITY
		}
	}

	item.sellIn--

	if item.sellIn < 0 {
		item.quality = 0
	}
}
