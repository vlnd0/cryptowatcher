package currencies

import (
	"errors"
	"fmt"
	"github.com/f-sev/cryptowatcher/config"
	"github.com/f-sev/cryptowatcher/internal/utils"
	"strconv"
	"strings"
)

var exchangeRates map[string]float64
var DefaultFiatCurrency string

type CurrencyListJson struct {
	Data []CurrencyInfoJson `json:"data"`
}

type CurrencyInfoJson struct {
	Symbol   string `json:"symbol"`
	PriceUsd string `json:"priceUsd"`
}

func init() {
	DefaultFiatCurrency = utils.GetEnv("DEFAULT_FIAT_CURRENCY", "USD")
	exchangeRates = loadExchangeRates()
}

func loadExchangeRates() map[string]float64 {
	fmt.Print("Loading exchange rates for ")
	response := make(map[string]float64, 10)

	var currencyJson CurrencyListJson
	err := utils.GetJson(fmt.Sprintf("https://api.coincap.io/v2/assets?ids=%s", strings.Join(config.CurrencyList[:], ",")), &currencyJson)
	if err != nil {
		fmt.Printf("Error getting trone data	%s\n", err.Error())
	}

	for _, token := range currencyJson.Data {
		value, _ := strconv.ParseFloat(token.PriceUsd, 64)
		response[token.Symbol] = value
	}

	return response
}

func GetFiatRate(crypto string) (res float64, err error) {
	if val, ok := exchangeRates[crypto]; ok {
		return val, nil
	}

	return 0, errors.New("no target currency in exchange rates list")
}
