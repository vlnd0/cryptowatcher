package data

var KucoinWallet = KucoinDataSource{
	CryptoDataSource{
		Name: "Kucoin Exchange",
	},
	KucoinCredentials{
		apiKey:    "Qwerty",
		secretKey: "Secret",
	},
}

type KucoinCredentials struct {
	apiKey    string
	secretKey string
}

type KucoinDataSource struct {
	CryptoDataSource
	KucoinCredentials
}

func (k *KucoinDataSource) Collect() {
	k.Balance = make(BalanceType)

	k.Balance["USDC"] = 200
	k.Balance["SOL"] = 50
}

func (k *KucoinDataSource) Format() {
	k.Balance.Format()
}
