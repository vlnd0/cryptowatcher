package currencies

import (
	"errors"
	"github.com/f-sev/cryptowatcher/internal/utils"
)

var exchangeRates map[string]float64
var DefaultFiatCurrency string

func init() {
	DefaultFiatCurrency = utils.GetEnv("DEFAULT_FIAT_CURRENCY", "USD")
	exchangeRates = loadExchangeRates(DefaultFiatCurrency)
}

func loadExchangeRates(currency string) map[string]float64 {
	//	println("Loading exchange rates for ", currency, "...")

	return map[string]float64{
		"BTC":  17765.10,
		"ETH":  1317.21,
		"SOL":  17.97,
		"TRX":  0.0590,
		"USDC": 1.008,
		"USDT": 0.9988,
		"DAI":  0.9998,
	}
}

func GetFiatRate(crypto string) (res float64, err error) {
	if val, ok := exchangeRates[crypto]; ok {
		return val, nil
	}

	return 0, errors.New("no target currency in exchange rates list")
}
