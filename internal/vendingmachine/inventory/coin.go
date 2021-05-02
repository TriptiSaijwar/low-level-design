package inventory

type Coin int

const (
	PENNY Coin = iota
	NICKLE
	DIME
	QUARTER
	numOfCoins
)

func (c Coin) GetName() string {
	return [...]string{"PENNY", "NICKLE", "DIME", "QUARTER"}[c]
}

func (c Coin) GetValue() int {
	return [...]int{1, 5, 10, 25}[c]
}

func (c Coin) GetId() int {
	return int(c)
}

func GetCoins() []Coin {
	coins := make([]Coin, numOfCoins)
	for coin := Coin(0); coin < numOfCoins; coin++ {
		coins = append(coins, coin)
	}
	return coins
}
