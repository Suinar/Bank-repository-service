package fixture

import (
	core "Bank-repository-service/pkg/core"
	currency "Bank-repository-service/proto/repository/currency"
)

func NewCurrencyCore(isoCode string, name string, symbol rune, minorUnits int8) core.Currency {
	return core.Currency{
		Id:         TestId,
		IsoCode:    isoCode,
		Name:       name,
		Symbol:     symbol,
		MinorUnits: minorUnits,
	}
}

func NewCurrencyUpdateInputCore(isoCode string, name string, symbol rune, minorUnits int8) *core.CurrencyUpdateInput {
	return &core.CurrencyUpdateInput{
		Name:       StringPointer(name),
		Symbol:     RunePointer(symbol),
		IsoCode:    StringPointer(isoCode),
		MinorUnits: Int8Pointer(minorUnits),
	}
}

func NewCurrencyProto(isoCode string, name string, symbol string, minorUnits int32) *currency.Currency {
	return &currency.Currency{
		Id:         TestId,
		IsoCode:    isoCode,
		Name:       name,
		Symbol:     symbol,
		MinorUnits: minorUnits,
	}
}

func NewCurrencyListProto(
	isoCodeFirst string, isoCodeSecond string,
	nameFirst string, nameSecond string,
	symbolFirst string, symbolSecond string,
	minorUnitsFirst int32, minorUnitsSecond int32) *currency.CurrencyList {
	return &currency.CurrencyList{
		Currencies: []*currency.Currency{
			NewCurrencyProto(isoCodeFirst, nameFirst, symbolFirst, minorUnitsFirst),
			NewCurrencyProto(isoCodeSecond, nameSecond, symbolSecond, minorUnitsSecond),
		},
	}
}

func NewIsoCodeRequestProto(iso string) *currency.IsoCodeRequest {
	return &currency.IsoCodeRequest{
		IsoCode: iso,
	}
}

func NewSymbolRequestProto(symbol string) *currency.SymbolRequest {
	return &currency.SymbolRequest{
		Symbol: symbol,
	}
}

func NewCurrencyUpdateInputProto(isoCode string, name string, symbol string, minorUnits int32) *currency.CurrencyUpdateInput {
	return &currency.CurrencyUpdateInput{
		Name:       StringPointer(name),
		Symbol:     StringPointer(symbol),
		IsoCode:    StringPointer(isoCode),
		MinorUnits: Int32Pointer(minorUnits),
	}
}

func NewUpdateCurrencyRequestProto(isoCode string, name string, symbol string, minorUnits int32) *currency.UpdateCurrencyRequest {
	return &currency.UpdateCurrencyRequest{
		Id:    1,
		Input: NewCurrencyUpdateInputProto(isoCode, name, symbol, minorUnits),
	}
}
