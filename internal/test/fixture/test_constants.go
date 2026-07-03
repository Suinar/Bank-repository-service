package fixture

const (
	TestId int64 = 1

	TestName       = "Test"
	TestUpdateName = "TestUpdate"

	TestAmount = int64(100_000)

	TestFirstName       = "Mykola"
	TestMidlName        = "Ihorovich"
	TestLastName        = "Kachmaryk"
	TestUpdateFirstName = "Jane"
	TestUpdateMidlName  = "Ann"
	TestUpdateLastName  = "Smith"

	TestEmail       = "john.doe@example.com"
	TestPhoneNumber = "+380123456789"

	TestPassword     = "Password123"
	TestPasswordHash = "$2a$10$abcdefghijklmnopqrstuv"

	TestCardNumber = "1234567890123456"

	TestTermMonths = 24

	TestIsoCode               = "USD"
	TestCurrencyName          = "US Dollar"
	TestSymbolString          = "$"
	TestSymbolRune            = '$'
	TestMinorUnitsInt32       = int32(2)
	TestMinorUnitsInt8        = int8(2)
	TestUpdateIsoCode         = "EUR"
	TestUpdateCurrencyName    = "Euro"
	TestUpdateSymbolString    = "€"
	TestUpdateSymbolRune      = '€'
	TestUpdateMinorUnitsInt32 = int32(2)
	TestUpdateMinorUnitsInt8  = int8(2)
)

func StringPointer(v string) *string { return &v }

func RunePointer(v rune) *rune { return &v }

func Int32Pointer(v int32) *int32 { return &v }

func Int8Pointer(v int8) *int8 { return &v }
