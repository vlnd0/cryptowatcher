package data

import (
    "github.com/Rhymond/go-money"
    "github.com/getlantern/systray"
)

var KucoinWallet = KucoinDataSource{
	CryptoDataSource{
		Name: "Kucoin",
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
	// TODO: Load assets!

	k.Balance["USDC"] = 200.0
	k.Balance["SOL"] = 1005.0555
}

func (k *KucoinDataSource) TotalFiat() *money.Money {
	return k.Balance.TotalFiat()
}

func (k *KucoinDataSource) Display() {
	item := systray.AddMenuItem(k.Name + "(" + k.Balance.TotalFiat().Display() + ")", "")
    k.Balance.Display(item)
}
