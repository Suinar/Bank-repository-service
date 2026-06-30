package fixture

import "Bank-repository-service/proto/repository/currency"

func NewCurrency(isoCode string, name string, symbol string, minorUnits int32) *currency.Currency {
	return &currency.Currency{
		Id:         TestId,
		IsoCode:    isoCode,
		Name:       name,
		Symbol:     symbol,
		MinorUnits: minorUnits,
	}
}

func NewCurrencyList(
	isoCodeFirst string, isoCodeSecond string,
	nameFirst string, nameSecond string,
	symbolFirst string, symbolSecond string,
	minorUnitsFirst int32, minorUnitsSecond int32) *currency.CurrencyList {
	return &currency.CurrencyList{
		Currencies: []*currency.Currency{
			NewCurrency(isoCodeFirst, nameFirst, symbolFirst, minorUnitsFirst),
			NewCurrency(isoCodeSecond, nameSecond, symbolSecond, minorUnitsSecond),
		},
	}
}

func NewIsoCodeRequest(iso string) *currency.IsoCodeRequest {
	return &currency.IsoCodeRequest{
		IsoCode: iso,
	}
}

func NewSymbolRequest(symbol string) *currency.SymbolRequest {
	return &currency.SymbolRequest{
		Symbol: symbol,
	}
}

func NewCurrencyUpdateInput() *currency.CurrencyUpdateInput {
	return &currency.CurrencyUpdateInput{
		Name:       String(TestUpdateCurrencyName),
		Symbol:     String(TestUpdateSymbol),
		IsoCode:    String(TestUpdateIsoCode),
		MinorUnits: Int32(TestUpdateMinorUnits),
	}
}

func NewUpdateCurrencyRequest() *currency.UpdateCurrencyRequest {
	return &currency.UpdateCurrencyRequest{
		Id:    1,
		Input: NewCurrencyUpdateInput(),
	}
}
