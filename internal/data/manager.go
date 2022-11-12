package data

import (
	"github.com/Rhymond/go-money"
	"github.com/f-sev/cryptowatcher/internal/currencies"
	"github.com/getlantern/systray"
)

var Manager = CryptoManager{}

func init() {
	Manager.cryptoDataSources = []Collectible{
		&TronWallet,
		&EthWallet,
		&BscWallet,
		&HuobiWallet,
		//&KucoinWallet,
	}
}

type CryptoManager struct {
	cryptoDataSources []Collectible
}

func (m *CryptoManager) Collect() {
	total := money.New(0.0, currencies.DefaultFiatCurrency)

	for _, value := range m.cryptoDataSources {
		value.Collect()
		total, _ = total.Add(value.TotalFiat())
	}

	systray.SetTitle("🌕 Balance: " + total.Display())
	for _, value := range m.cryptoDataSources {
		value.Display()
	}
}
