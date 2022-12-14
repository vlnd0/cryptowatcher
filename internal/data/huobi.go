package data

import (
	"fmt"
	"github.com/Rhymond/go-money"
	"github.com/f-sev/cryptowatcher/config"
	"github.com/getlantern/systray"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"strconv"
	"strings"
)

var HuobiWallet = HuobiDataSource{
	CryptoDataSource{
		Name: "Huobi",
		Info: TrayElement{
			Content: "",
		},
		ChildElements: make(map[string]TrayElement),
	},
	HuobiCredentials{
		accessKey: config.HuobiAccessKey,
		secretKey: config.HuobiSecretKey,
		host:      config.HuobiHost,
		accountId: config.HuobiAccountId,
	},
}

type HuobiCredentials struct {
	accessKey string
	secretKey string
	host      string
	accountId string
}

type HuobiDataSource struct {
	CryptoDataSource
	HuobiCredentials
}

func (h *HuobiDataSource) Collect() {
	h.Balance = make(BalanceType)
	huobiClient := new(client.AccountClient).Init(h.accessKey, h.secretKey, h.host)
	resp, err := huobiClient.GetAccountBalance(h.accountId)
	if err != nil {
		fmt.Printf("error. Huobi. %s", err.Error())
	}

	for _, token := range resp.List {
		if token.Balance != "0" {
			val, _ := strconv.ParseFloat(token.Balance, 64)
			h.Balance[strings.ToUpper(token.Currency)] = val
		}
	}

}

func (h *HuobiDataSource) TotalFiat() *money.Money {
	return h.Balance.TotalFiat()
}

func (h *HuobiDataSource) Display() {
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
