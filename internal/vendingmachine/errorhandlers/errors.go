package errorhandlers

type ItemSoldOutError string

func (err ItemSoldOutError) Error() string {
	return string(err)
}

type NotSufficientChange string

func (err NotSufficientChange) Error() string {
	return string(err)
}

type PriceNotPaid string

func (err PriceNotPaid) Error() string {
	return string(err)
}