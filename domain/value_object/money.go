package value_object

import "errors"

type money struct {
	amount   float64
	currency Currency
}

func NewMoney(amount float64, currency Currency) (*money, error) {
	if currency.IsUndefined() {
		return nil, errors.New("currency is undefined")
	}

	return &money{
		amount:   amount,
		currency: currency,
	}, nil
}

func (m *money) Add(arg *money) (*money, error) {
	if arg == nil {
		return nil, errors.New("arg is nil")
	}

	if m.currency != arg.currency {
		return nil, errors.New("currency is different")
	}

	return NewMoney(m.amount+arg.amount, m.currency)
}
