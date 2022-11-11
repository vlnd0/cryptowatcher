package data

import (
	"github.com/Rhymond/go-money"
	"github.com/f-sev/cryptowatcher/internal/currencies"
    "github.com/getlantern/systray"
	"strings"
)

type Collectible interface {
	Collect()
    TotalFiat() *money.Money
	Display()
}

type CryptoDataSource struct {
	Name    string
	Balance BalanceType
}

type BalanceType map[string]float64

func (b BalanceType) Display(item *systray.MenuItem) {

    for key, value := range b {
        // TODO: Handle error, when there is no exchange rate for currency
		fiatRate, _ := currencies.GetFiatRate(key)
		cryptoMoney := money.NewFromFloat(value, key)
		fiatMoney := money.NewFromFloat(fiatRate * value, currencies.DefaultFiatCurrency)

        var sb strings.Builder
        sb.WriteString(cryptoMoney.Display())
		sb.WriteString(" ")
		sb.WriteString("(" + fiatMoney.Display() + ")")
        item.AddSubMenuItem(sb.String(), "")
	}
}

func (b BalanceType) TotalFiat() *money.Money {
    sum := 0.0
    for key, value := range b {
        // TODO: Handle error, when there is no exchange rate for currency
        fiatRate, _ := currencies.GetFiatRate(key)

        sum += fiatRate * value
    }
    return money.NewFromFloat(sum, currencies.DefaultFiatCurrency)
}