package fixture

import (
	core "Bank-repository-service/pkg/core"
	currency "Bank-repository-service/proto/repository/currency"
)

func NewCurrencyCore() core.Currency {
	return core.Currency{
		Id:         TestId,
		IsoCode:    TestIsoCode,
		Name:       TestName,
		Symbol:     TestSymbolRune,
		MinorUnits: TestMinorUnitsInt8,
	}
}

func NewCurrencyUpdateInputCore() *core.CurrencyUpdateInput {
	return &core.CurrencyUpdateInput{
		Name:       StringPointer(TestCurrencyName),
		Symbol:     RunePointer(TestSymbolRune),
		IsoCode:    StringPointer(TestIsoCode),
		MinorUnits: Int8Pointer(TestMinorUnitsInt8),
	}
}

func NewCurrencyProto() *currency.Currency {
	return &currency.Currency{
		Id:         TestId,
		IsoCode:    TestIsoCode,
		Name:       TestName,
		Symbol:     TestSymbolString,
		MinorUnits: TestMinorUnitsInt32,
	}
}

func NewCurrencyListProto() *currency.CurrencyList {
	return &currency.CurrencyList{
		Currencies: []*currency.Currency{
			NewCurrencyProto(),
			NewCurrencyProto(),
		},
	}
}

func NewIsoCodeRequestProto() *currency.IsoCodeRequest {
	return &currency.IsoCodeRequest{
		IsoCode: TestIsoCode,
	}
}

func NewSymbolRequestProto() *currency.SymbolRequest {
	return &currency.SymbolRequest{
		Symbol: TestSymbolString,
	}
}

func NewCurrencyUpdateInputProto() *currency.CurrencyUpdateInput {
	return &currency.CurrencyUpdateInput{
		Name:       StringPointer(TestCurrencyName),
		Symbol:     StringPointer(TestSymbolString),
		IsoCode:    StringPointer(TestIsoCode),
		MinorUnits: Int32Pointer(TestMinorUnitsInt32),
	}
}

func NewUpdateCurrencyRequestProto() *currency.UpdateCurrencyRequest {
	return &currency.UpdateCurrencyRequest{
		Id:    1,
		Input: NewCurrencyUpdateInputProto(),
	}
}
