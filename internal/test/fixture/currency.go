package fixture

import (
	currency "github.com/Suinar/Bank-proto/repository/currency"
	core "github.com/Suinar/Bank-repository-service/pkg/core"
)

// NewCurrencyCore creates a ready-to-use currency core.
func NewCurrencyCore() core.Currency {
	return core.Currency{
		Id:         TestId,
		IsoCode:    TestIsoCode,
		Name:       TestCurrencyName,
		Symbol:     TestSymbolRune,
		MinorUnits: TestMinorUnitsInt8,
	}
}

// NewCurrencyCoreInputIdAndIsoAndName creates a ready-to-use currency core input id and iso and name.
func NewCurrencyCoreInputIdAndIsoAndName(id int64, isoCode string, name string) core.Currency {
	return core.Currency{
		Id:         id,
		IsoCode:    isoCode,
		Name:       name,
		Symbol:     TestSymbolRune,
		MinorUnits: TestMinorUnitsInt8,
	}
}

// NewCurrencyUpdateInputCore creates a ready-to-use currency update input core.
func NewCurrencyUpdateInputCore() *core.CurrencyUpdateInput {
	return &core.CurrencyUpdateInput{
		Name:       StringPointer(TestCurrencyName),
		Symbol:     RunePointer(TestSymbolRune),
		IsoCode:    StringPointer(TestIsoCode),
		MinorUnits: Int8Pointer(TestMinorUnitsInt8),
	}
}

// NewCurrencyProto creates a ready-to-use currency proto.
func NewCurrencyProto() *currency.Currency {
	return &currency.Currency{
		Id:         TestId,
		IsoCode:    TestIsoCode,
		Name:       TestCurrencyName,
		Symbol:     TestSymbolString,
		MinorUnits: TestMinorUnitsInt32,
	}
}

// NewCurrencyListProto creates a ready-to-use currency list proto.
func NewCurrencyListProto() *currency.CurrencyList {
	return &currency.CurrencyList{
		Currencies: []*currency.Currency{
			NewCurrencyProto(),
			NewCurrencyProto(),
		},
	}
}

// NewIsoCodeRequestProto creates a ready-to-use iso code request proto.
func NewIsoCodeRequestProto() *currency.IsoCodeRequest {
	return &currency.IsoCodeRequest{
		IsoCode: TestIsoCode,
	}
}

// NewSymbolRequestProto creates a ready-to-use symbol request proto.
func NewSymbolRequestProto() *currency.SymbolRequest {
	return &currency.SymbolRequest{
		Symbol: TestSymbolString,
	}
}

// NewCurrencyUpdateInputProto creates a ready-to-use currency update input proto.
func NewCurrencyUpdateInputProto() *currency.CurrencyUpdateInput {
	return &currency.CurrencyUpdateInput{
		Name:       StringPointer(TestCurrencyName),
		Symbol:     StringPointer(TestSymbolString),
		IsoCode:    StringPointer(TestIsoCode),
		MinorUnits: Int32Pointer(TestMinorUnitsInt32),
	}
}

// NewUpdateCurrencyRequestProto creates a ready-to-use update currency request proto.
func NewUpdateCurrencyRequestProto() *currency.UpdateCurrencyRequest {
	return &currency.UpdateCurrencyRequest{
		Id:    1,
		Input: NewCurrencyUpdateInputProto(),
	}
}
