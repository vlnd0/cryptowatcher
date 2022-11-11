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

type HuobiJson struct {
	Tokens      []TronTokenJson `json:"tokens"`
	TotalFrozen int32           `json:"totalFrozen"`
}

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
	huobiClient := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
	//resp, err := huobiClient.GetAccountInfo()
	resp, err := huobiClient.GetAccountBalance(config.AccountId)
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
	item := systray.AddMenuItem(h.Name+"("+h.Balance.TotalFiat().Display()+")", "")
	h.Balance.Display(item)
}
