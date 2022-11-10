package data

var HuobiWallet = HuobiDataSource{
	CryptoDataSource{
		Name: "Huobi Exchange",
	},
	HuobiCredentials{
		apiKey:    "Qwerty",
		secretKey: "Secret",
	},
}

type HuobiCredentials struct {
	apiKey    string
	secretKey string
}

type HuobiDataSource struct {
	CryptoDataSource
	HuobiCredentials
}

func (t *HuobiDataSource) Collect() {
	t.Balance = make(map[string]float32)

	t.Balance["USDT"] = 100
	t.Balance["TRX"] = 50
}

func (h *HuobiDataSource) Format() {
	h.Balance.Format()
}
