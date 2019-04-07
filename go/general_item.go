package main

type GeneralItem struct {
	*Item
}

func (item *GeneralItem) Process() {
	if item.quality > 0 {
		item.quality--
	}

	item.sellIn--

	if item.sellIn < 0 && item.quality > 0 {
		item.quality--
	}
}
