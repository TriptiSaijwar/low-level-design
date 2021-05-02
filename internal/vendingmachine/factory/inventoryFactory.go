package factory

import (
	"fmt"
	"github.com/TriptiSaijwar/low-level-design/internal/vendingmachine/inventory"
)

type InventoryType string

const (
	CASH InventoryType = "CASH"
	ITEM InventoryType = "ITEM"
)

func GetInventory(inventoryType InventoryType) (inventory.IInventory, error) {
	switch inventoryType {
	case CASH:
		return inventory.NewInventory(), nil
	case ITEM:
		return inventory.NewInventory(), nil
	default:
		return nil, fmt.Errorf("Wrong invetory type passed")
	}
}
