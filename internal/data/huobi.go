package data

import (
    "github.com/Rhymond/go-money"
    "github.com/getlantern/systray"
)

var HuobiWallet = HuobiDataSource{
	CryptoDataSource{
		Name: "Huobi",
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

func (h *HuobiDataSource) Collect() {
	h.Balance = make(BalanceType)
	// TODO: Load assets!

	h.Balance["USDT"] = 1005.0555
	h.Balance["TRX"] = 1005.0555
}

func (h *HuobiDataSource) TotalFiat() *money.Money {
	return h.Balance.TotalFiat()
}

func (h *HuobiDataSource) Display() {
    item := systray.AddMenuItem(h.Name + "(" + h.Balance.TotalFiat().Display() + ")", "")
    h.Balance.Display(item)
}
