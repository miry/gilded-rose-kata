package main

type ConjuredItem struct {
	*Item
}

func (item *ConjuredItem) Process() {
	if item.quality > 0 {
		item.quality -= 2

		if item.quality < 0 {
			item.quality = 0
		}
	}

	item.sellIn--

	if item.quality == 0 {
		return
	}

	if item.sellIn < 0 {
		item.quality -= 2

		if item.quality < 0 {
			item.quality = 0
		}
	}
}
