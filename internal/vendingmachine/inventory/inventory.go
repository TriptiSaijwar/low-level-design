package inventory

type IInventory interface {
	Add(itemId int)
	Deduct(itemId int)
	HasItem(itemId int) bool
	Clear()
	Put(itemId, quantity int)
	GetQuantity(itemId int) int
}

type Inventory struct {
	inventory map[int]int
}

func NewInventory() *Inventory {
	inventory := make(map[int]int)
	return &Inventory{
		inventory: inventory,
	}
}

func (i *Inventory) Add(itemId int) {
	i.inventory[itemId] = i.inventory[itemId] + 1
}

func (i *Inventory) Deduct(itemId int) {
	i.inventory[itemId] = i.inventory[itemId] - 1
}

func (i *Inventory) Put(itemId, quantity int) {
	i.inventory[itemId] = quantity
}

func (i *Inventory) HasItem(itemId int) bool {
	if _, ok := i.inventory[itemId]; ok {
		return true
	}
	return false
}

func (i *Inventory) Clear()  {
	i.inventory = make(map[int]int)
}

func (i *Inventory) GetQuantity(itemId int) int {
	return i.inventory[itemId]
}

