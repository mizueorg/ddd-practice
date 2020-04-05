package value

//go:generate stringer -type Currency ./currency.go
type Currency int

const (
	_ Currency = iota
	JPY
	USD
	EUR
)

func (i Currency) IsUndefined() bool {
	return i < JPY || i > EUR
}
