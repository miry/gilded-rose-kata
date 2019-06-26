package main

type GeneralItem struct {
	*InventoryItem
}

func (item *GeneralItem) Process() {
	if item.GetQuality() > 0 {
		item.DecrQuality() // = quality -= speed
	}

	item.sellIn--

	if item.sellIn < 0 {
		item.DecrQuality() // <= && item.quality > 0
	}
}
