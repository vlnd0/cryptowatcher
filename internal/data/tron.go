package data

import (
    "github.com/Rhymond/go-money"
    "github.com/getlantern/systray"
)

var TronWallet = TronDataSource{
	CryptoDataSource{
		Name: "Tron Wallet",
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
	// TODO: Load assets!

	t.Balance["USDT"] = 100.0
	t.Balance["TRX"] = 50.0
}

func (t *TronDataSource) TotalFiat() *money.Money {
    return t.Balance.TotalFiat()
}

func (t *TronDataSource) Display() {
	item := systray.AddMenuItem(t.Name + "(" + t.Balance.TotalFiat().Display() + ")", "")
    t.Balance.Display(item)
}
