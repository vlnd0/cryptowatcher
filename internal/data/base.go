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
	Name          string
	Balance       BalanceType
	Info          TrayElement
	ChildElements ChildrenElementsType
}

type TrayElement struct {
	Content string
	Item    *systray.MenuItem
}

type BalanceType map[string]float64

type ChildrenElementsType map[string]TrayElement

func (b BalanceType) Display(item *systray.MenuItem, children ChildrenElementsType) {
	for key, value := range b {
		child, isExists := children[key]
		// TODO: Handle error, when there is no exchange rate for currency
		fiatRate, _ := currencies.GetFiatRate(key)
		cryptoMoney := money.NewFromFloat(value, key)
		fiatMoney := money.NewFromFloat(fiatRate*value, currencies.DefaultFiatCurrency)

		var sb strings.Builder
		sb.WriteString(cryptoMoney.Display())
		sb.WriteString(" ")
		sb.WriteString("(" + fiatMoney.Display() + ")")
		content := sb.String()

		if isExists {
			child.Content = content
			child.Item.SetTitle(content)
		} else {
			child := item.AddSubMenuItem(content, "")
			children[key] = TrayElement{
				Item:    child,
				Content: content,
			}
		}

	}
}

func (b BalanceType) TotalFiat() *money.Money {
	sum := 0.0
	for key, value := range b {
		fiatRate, err := currencies.GetFiatRate(key)
		if err != nil {
			sum += 1 * value
		} else {
			sum += fiatRate * value
		}
	}
	return money.NewFromFloat(sum, currencies.DefaultFiatCurrency)
}
