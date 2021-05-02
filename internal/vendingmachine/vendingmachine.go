package vendingmachine

import (
	"fmt"
	errors "github.com/TriptiSaijwar/low-level-design/internal/vendingmachine/errorhandlers"
	"github.com/TriptiSaijwar/low-level-design/internal/vendingmachine/factory"
	"github.com/TriptiSaijwar/low-level-design/internal/vendingmachine/inventory"
)

type IVendingMachine interface {
	InsertCoin(coin inventory.Coin)
	GetItemAndChange() (item inventory.Item, coins []inventory.Coin, err error)
	SelectItemAndGetPrice(item inventory.Item) (price int, err error)
	Refund() (coins []inventory.Coin, err error)
}

type VendingMachine struct {
	TotalSales     int
	CurrentItem    *inventory.Item
	CurrentBalance int
	CashInventory  inventory.IInventory
	ItemInventory  inventory.IInventory
}

func NewVendingMachine() (*VendingMachine, error) {
	cashInventory, err := factory.GetInventory(factory.CASH)
	itemInventory, err := factory.GetInventory(factory.ITEM)
	if err != nil {
		return nil, err
	}

	for _, coin := range inventory.GetCoins() {
		cashInventory.Put(coin.GetId(), 5)
	}

	for _, item := range inventory.GetItems() {
		itemInventory.Put(item.GetId(), 5)
	}

	return &VendingMachine{
		TotalSales:     0,
		CurrentItem:    nil,
		CurrentBalance: 0,
		CashInventory:  cashInventory,
		ItemInventory:  itemInventory,
	}, nil
}

func (v *VendingMachine) InsertCoin(coin inventory.Coin) {
	v.CashInventory.Add(coin.GetId())
	v.CurrentBalance = v.CurrentBalance + coin.GetValue()
}

func (v *VendingMachine) GetItemAndChange() (item inventory.Item, coins []inventory.Coin, err error) {
	if !v.isFullyPaid() {
		err = errors.PriceNotPaid(fmt.Sprintf("Price is not fully paid. Remaining balance %d",v.CurrentItem.GetPrice()-v.CurrentBalance))
		return
	}
	coins, err = v.collectChange(v.CurrentBalance - v.CurrentItem.GetPrice())
	if err != nil {
		return
	}
	item = v.collectItem()
	v.TotalSales = v.TotalSales + item.GetPrice()
	return
}

func (v *VendingMachine) isFullyPaid() bool {
	if v.CurrentBalance > v.CurrentItem.GetPrice() {
		return true
	}
	return false
}

func (v *VendingMachine) collectItem() (item inventory.Item) {
	v.ItemInventory.Deduct(v.CurrentItem.GetId())
	item = *v.CurrentItem
	v.CurrentItem = nil
	return item
}

func (v *VendingMachine) collectChange(amount int) (coins []inventory.Coin, err error) {
	coins, err = v.getRequiredChange(amount)
	if err != nil {
		return
	}
	for _, coin := range coins {
		v.CashInventory.Deduct(coin.GetId())
	}
	v.CurrentBalance = 0
	return
}

func (v *VendingMachine) getRequiredChange(amount int) (changeCoins []inventory.Coin, err error) {
	change := amount
	for ; change > 0; {
		switch {
		case change >= inventory.QUARTER.GetValue() && v.CashInventory.HasItem(inventory.QUARTER.GetId()):
			change = change - inventory.QUARTER.GetValue()
			changeCoins = append(changeCoins, inventory.QUARTER)
			break
		case change >= inventory.DIME.GetValue() && v.CashInventory.HasItem(inventory.DIME.GetId()):
			change = change - inventory.DIME.GetValue()
			changeCoins = append(changeCoins, inventory.DIME)
			break
		case change >= inventory.NICKLE.GetValue() && v.CashInventory.HasItem(inventory.NICKLE.GetId()):
			change = change - inventory.NICKLE.GetValue()
			changeCoins = append(changeCoins, inventory.NICKLE)
			break
		case change >= inventory.PENNY.GetValue() && v.CashInventory.HasItem(inventory.PENNY.GetId()):
			change = change - inventory.PENNY.GetValue()
			changeCoins = append(changeCoins, inventory.PENNY)
			break
		default:
			return nil, errors.NotSufficientChange(fmt.Sprintf("Not Sufficient Change Rs %d", change))
		}
	}
	return
}

func (v *VendingMachine) SelectItemAndGetPrice(item inventory.Item) (price int, err error) {
	if !v.ItemInventory.HasItem(item.GetId()) {
		return 0, errors.ItemSoldOutError("Item Sold Out.")
	}

	v.CurrentItem = &item
	price = item.GetPrice()
	return
}

func (v *VendingMachine) Refund() (coins []inventory.Coin, err error) {
	coins, err = v.getRequiredChange(v.CurrentBalance)
	if err != nil {
		return
	}
	for _, coin := range coins {
		v.CashInventory.Deduct(coin.GetId())
	}
	v.CurrentBalance = 0
	v.CurrentItem = nil
	return
}

func (v *VendingMachine) clear() {
	v.CashInventory = nil
	v.ItemInventory = nil
	v.CurrentItem = nil
	v.CurrentBalance = 0
	v.TotalSales = 0
}
