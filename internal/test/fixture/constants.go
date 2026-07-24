package fixture

const (
	TestId int64 = 1

	TestName = "Test"

	TestAmount = int64(100000)

	TestFirstName = "Mykola"
	TestMidlName  = "Ihorovich"
	TestLastName  = "Kachmaryk"

	TestEmail       = "john.doe@example.com"
	TestPhoneNumber = "+380123456789"

	TestPassword     = "Password123"
	TestPasswordHash = "$2a$10$abcdefghijklmnopqrstuv"

	TestCardNumber = "1234567890123456789"

	TestTermMonths = 24

	TestIsoCode         = "USD"
	TestCurrencyName    = "US Dollar"
	TestSymbolString    = "$"
	TestSymbolRune      = '$'
	TestMinorUnitsInt32 = int32(2)
	TestMinorUnitsInt8  = int8(2)
)

// StringPointer returns a pointer to the supplied fixture value.
func StringPointer(v string) *string { return &v }

// RunePointer returns a pointer to the supplied fixture value.
func RunePointer(v rune) *rune { return &v }

// Int32Pointer returns a pointer to the supplied fixture value.
func Int32Pointer(v int32) *int32 { return &v }

// Int8Pointer returns a pointer to the supplied fixture value.
func Int8Pointer(v int8) *int8 { return &v }
