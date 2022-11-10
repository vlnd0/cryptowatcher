package data

var TronWallet = TronDataSource{
	CryptoDataSource{
		Name: "Tron Blockchain",
	},
	TronCredentials{
		apiKey: "Qwerty",
	},
}

type TronCredentials struct {
	apiKey string
}

type TronDataSource struct {
	CryptoDataSource
	TronCredentials
}

func (t *TronDataSource) Collect() {
	t.Balance = make(BalanceType)

	t.Balance["USDT"] = 100
	t.Balance["TRX"] = 50
}

func (t *TronDataSource) Format() {
    t.Balance.Format()
}

