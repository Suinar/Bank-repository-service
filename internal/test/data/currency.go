package data

import "Bank-repository-service/proto/repository/currency"

func NewCurrency() *currency.Currency {
	return &currency.Currency{
		Id:         1,
		IsoCode:    "USD",
		Name:       "Dollar",
		Symbol:     "$",
		MinorUnits: 2,
	}
}
