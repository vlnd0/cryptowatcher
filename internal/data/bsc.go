package data

import (
	"fmt"
	"github.com/Rhymond/go-money"
	"github.com/f-sev/cryptowatcher/config"
	"github.com/getlantern/systray"
	"github.com/nanmu42/etherscan-api"
	"math/big"
	"time"
)

var BscWallet = BscDataSource{
	CryptoDataSource{
		Name: "BSC Wallet",
	},
	BscCredentials{
		apiKey:        config.BscApiKey,
		walletAddress: config.BscWalletAddress,
		tokensAddress: []TokenInfo{
			{
				address: config.BscDaiAddress,
				token:   "DAI",
			},
			{
				address: config.BscBusdAddress,
				token:   "BUSD",
			},
			{
				address: config.BscUsdcAddress,
				token:   "USDC",
			},
		},
	},
}

type BscCredentials struct {
	apiKey        string
	walletAddress string
	tokensAddress []TokenInfo
}

type BscDataSource struct {
	CryptoDataSource
	BscCredentials
}

func (t *BscDataSource) Collect() {
	t.Balance = make(BalanceType)
	client := etherscan.NewCustomized(etherscan.Customization{
		Timeout: 15 * time.Second,
		Key:     t.apiKey,
		BaseURL: "https://api.bscscan.com/api?",
		Verbose: false,
	})

	eth, err := client.AccountBalance(t.walletAddress)
	if err != nil {
		fmt.Printf("ether error %s", err.Error())
	} else {
		value, _ := new(big.Float).SetInt(eth.Int()).Float64()
		t.Balance["BNB"] = value / 1000000000000000000
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

func (t *BscDataSource) TotalFiat() *money.Money {
	return t.Balance.TotalFiat()
}

func (t *BscDataSource) Display() {
	item := systray.AddMenuItem(t.Name+"("+t.Balance.TotalFiat().Display()+")", "")
	t.Balance.Display(item)
}
