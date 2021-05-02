package inventory

type Item int

const (
	COKE Item = iota
	PEPSI
	SODA
	numOfItems
)

func (i Item) GetName() string {
	return [...]string{"Coke", "Pepsi", "Soda"}[i-1]
}

func (i Item) GetPrice() int {
	return [...]int{20, 25, 30}[i-1]
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
