package data

import (
	"fmt"
	"github.com/Rhymond/go-money"
	"github.com/f-sev/cryptowatcher/config"
	"github.com/getlantern/systray"
	"github.com/nanmu42/etherscan-api"
	"math/big"
)

var EthWallet = EthDataSource{
	CryptoDataSource{
		Name: "ETH Wallet",
		Info: TrayElement{
			Content: "",
		},
		ChildElements: make(map[string]TrayElement),
	},
	EthCredentials{
		apiKey:        config.EthApiKey,
		walletAddress: config.EthWalletAddress,
		tokensAddress: []TokenInfo{
			{
				address: config.EthDaiAddress,
				token:   "DAI",
			},
			{
				address: config.EthUsdtAddress,
				token:   "USDT",
			},
			{
				address: config.EthUsdcAddress,
				token:   "USDC",
			},
		},
	},
}

type EthCredentials struct {
	apiKey        string
	walletAddress string
	tokensAddress []TokenInfo
}

type EthDataSource struct {
	CryptoDataSource
	EthCredentials
}

type TokenInfo struct {
	token   string
	address string
}

func (t *EthDataSource) Collect() {
	t.Balance = make(BalanceType)
	client := etherscan.New(etherscan.Mainnet, t.apiKey)

	eth, err := client.AccountBalance(t.walletAddress)
	if err != nil {
		fmt.Printf("ether error %s", err.Error())
	} else {
		value, _ := new(big.Float).SetInt(eth.Int()).Float64()
		t.Balance["ETH"] = value / 1000000000000000000
	}

	for _, token := range t.tokensAddress {
		tokenValue, valErr := client.TokenBalance(token.address, t.walletAddress)
		if valErr != nil {
			fmt.Printf("etherium error, %s", valErr.Error())
		} else {
			value, _ := new(big.Float).SetInt(tokenValue.Int()).Float64()
			if value > 0.0 {
				t.Balance[token.token] = value / 1000000000000000000
			}
		}
	}
}

func (t *EthDataSource) TotalFiat() *money.Money {
	return t.Balance.TotalFiat()
}

func (t *EthDataSource) Display() {
	isNew := t.Info.Content == ""
	value := t.Name + "(" + t.Balance.TotalFiat().Display() + ")"
	t.Info.Content = value
	if isNew {
		item := systray.AddMenuItem(t.Info.Content, "")
		t.Info.Item = item
	} else {
		t.Info.Item.SetTitle(t.Info.Content)
	}
	t.Balance.Display(t.Info.Item, t.ChildElements)
}
