package data

import (
	"fmt"
	"github.com/Rhymond/go-money"
	"github.com/f-sev/cryptowatcher/config"
	"github.com/getlantern/systray"
	"github.com/nanmu42/etherscan-api"
	"math/big"
)

type ETHJson struct {
	Tokens      []ETHTokenJson `json:"tokens"`
	TotalFrozen int32          `json:"totalFrozen"`
}

type ETHTokenJson struct {
	Balance   string `json:"balance"`
	TokenAbbr string `json:"tokenAbbr"`
}

var ETHWallet = ETHDataSource{
	CryptoDataSource{
		Name: "ETH Wallet",
	},
	ETHCredentials{
		apiKey: "Qwerty",
	},
}

type ETHCredentials struct {
	apiKey string
}

type ETHDataSource struct {
	CryptoDataSource
	ETHCredentials
}

func (t *ETHDataSource) Collect() {
	t.Balance = make(BalanceType)
	client := etherscan.New(etherscan.Mainnet, config.ETHApiKey)

	// check account balance
	balance, err := client.AccountBalance(config.ETHWalletAddress)
	if err != nil {
		fmt.Printf("ether error %s", err.Error())
	} else {
		value, _ := new(big.Float).SetInt(balance.Int()).Float64()
		t.Balance["ETH"] = value / 1000000000000000000
	}

	// check token balance
	tokenBalance, balanceErr := client.TokenBalance(config.ETHDaiAddress, config.ETHWalletAddress)
	if balanceErr != nil {
		fmt.Printf("etherium error, %s %s", balanceErr.Error(), tokenBalance)
	} else {
		value, _ := new(big.Float).SetInt(tokenBalance.Int()).Float64()
		t.Balance["DAI"] = value / 1000000000000000000
	}
	return
}

func (t *ETHDataSource) TotalFiat() *money.Money {
	return t.Balance.TotalFiat()
}

func (t *ETHDataSource) Display() {
	item := systray.AddMenuItem(t.Name+"("+t.Balance.TotalFiat().Display()+")", "")
	t.Balance.Display(item)
}
