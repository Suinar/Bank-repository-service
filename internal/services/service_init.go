package service

type Services struct {
	AccountService  IAccountService
	CardService     ICardService
	CreditService   ICreditService
	CurrencyService ICurrencyService
	DepositService  IDepositService
	UserService     IUserService
}

func InitServices(
	accountService IAccountService,
	cardService ICardService,
	creditService ICreditService,
	currencyService ICurrencyService,
	depositService IDepositService,
	userService IUserService) *Services {
	return &Services{
		AccountService:  accountService,
		CardService:     cardService,
		CreditService:   creditService,
		CurrencyService: currencyService,
		DepositService:  depositService,
		UserService:     userService,
	}
}
