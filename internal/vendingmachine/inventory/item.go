package inventory

type Item int

const (
	COKE Item = iota
	PEPSI
	SODA
	numOfItems
)

func (i Item) GetName() string {
	return [...]string{"Coke", "Pepsi", "Soda"}[i]
}

func (i Item) GetPrice() int {
	return [...]int{25, 35, 40}[i]
}

func (c Item) GetId() int {
	return int(c)
}

func GetItems() []Item {
	items := make([]Item, numOfItems)
	for item := Item(0); item < numOfItems; item++ {
		items = append(items, item)
	}
	return items
}
