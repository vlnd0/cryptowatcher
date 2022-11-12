package data

import (
	"fmt"
	"github.com/Rhymond/go-money"
	"github.com/f-sev/cryptowatcher/config"
	"github.com/f-sev/cryptowatcher/internal/utils"
	"github.com/getlantern/systray"
	"strconv"
	"strings"
)

type TronJson struct {
	Tokens      []TronTokenJson `json:"tokens"`
	TotalFrozen int32           `json:"totalFrozen"`
}

type TronTokenJson struct {
	Balance   string `json:"balance"`
	TokenAbbr string `json:"tokenAbbr"`
}

var TronWallet = TronDataSource{
	CryptoDataSource{
		Name: "TRON Wallet",
	},
	TronCredentials{
		apiKey: "",
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

	var tronJson TronJson
	err := utils.GetJson(fmt.Sprintf("https://apilist.tronscan.org/api/account?address=%s", config.TronWalletAddress), &tronJson)
	if err != nil {
		fmt.Printf("Error getting trone data	%s\n", err.Error())
	}

	for _, token := range tronJson.Tokens {
		value, _ := strconv.Atoi(token.Balance)
		t.Balance[strings.ToUpper(token.TokenAbbr)] = float64(value) / 1000000
		if token.TokenAbbr == "trx" && tronJson.TotalFrozen > 0 {
			t.Balance[strings.ToUpper(token.TokenAbbr)] += float64(tronJson.TotalFrozen) / 1000000
		}
	}

}

func (t *TronDataSource) TotalFiat() *money.Money {
	return t.Balance.TotalFiat()
}

func (t *TronDataSource) Display() {
	item := systray.AddMenuItem(t.Name+"("+t.Balance.TotalFiat().Display()+")", "")
	t.Balance.Display(item)
}
