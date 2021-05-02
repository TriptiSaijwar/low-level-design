package main

import (
	"github.com/TriptiSaijwar/low-level-design/internal/vendingmachine"
	"github.com/TriptiSaijwar/low-level-design/internal/vendingmachine/inventory"
	"log"
)

func main() {
	vm, err := vendingmachine.NewVendingMachine()
	if err != nil {
		log.Fatal("Error in initialising vending machine ", err.Error())
	}
	vm.SelectItemAndGetPrice(inventory.COKE)
	vm.InsertCoin(inventory.QUARTER)
	item, coins, err := vm.GetItemAndChange()
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("Bon Appetit")
	log.Println(item)
	for _, coin := range coins {
		log.Println(coin.GetName(), " : ", coin.GetValue())
	}

}
