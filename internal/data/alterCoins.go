package data

import (
	"github.com/Rhymond/go-money"
	"github.com/getlantern/systray"
)

var AlterWallet = CoinsDataSource{
	CryptoDataSource{
		Name: "Coins",
		Info: TrayElement{
			Content: "",
		},
		ChildElements: make(map[string]TrayElement),
	},
	CoinsCredentials{
		apiKey: "",
	},
}

type CoinsCredentials struct {
	apiKey string
}

type CoinsDataSource struct {
	CryptoDataSource
	CoinsCredentials
}

func (h *CoinsDataSource) Collect() {
	h.Balance = make(BalanceType)
	h.Balance["ADA"] = 69.9
	h.Balance["DOT"] = 3.869
	h.Balance["DOGE"] = 268.37
}

func (h *CoinsDataSource) TotalFiat() *money.Money {
	return h.Balance.TotalFiat()
}

func (h *CoinsDataSource) Display() {
	isNew := h.Info.Content == ""
	value := h.Name + "(" + h.Balance.TotalFiat().Display() + ")"
	h.Info.Content = value
	if isNew {
		item := systray.AddMenuItem(h.Info.Content, "")
		h.Info.Item = item
	} else {
		h.Info.Item.SetTitle(h.Info.Content)
	}
	h.Balance.Display(h.Info.Item, h.ChildElements)
}
