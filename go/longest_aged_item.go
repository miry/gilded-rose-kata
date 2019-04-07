package main

type LongestAgedItem struct {
	*Item
}

func (item *LongestAgedItem) Process() {
	if item.quality < MAX_QUALITY {
		item.quality++
	}

	item.sellIn--

	if item.sellIn < 0 && item.quality < MAX_QUALITY {
		item.quality++
	}
}
