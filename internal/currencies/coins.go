package currencies

import "github.com/Rhymond/go-money"

func init() {
	money.AddCurrency("BTC", "BTC", "$ 1", ".", ",", 6)
	money.AddCurrency("ETH", "ETH", "$ 1", ".", ",", 6)
	money.AddCurrency("BNB", "BNB", "$ 1", ".", ",", 4)
	money.AddCurrency("SOL", "SOL", "$ 1", ".", ",", 6)
	money.AddCurrency("TRX", "TRX", "$ 1", ".", ",", 6)
	money.AddCurrency("ADA", "ADA", "$ 1", ".", ",", 2)
	money.AddCurrency("DOT", "DOT", "$ 1", ".", ",", 4)
	money.AddCurrency("DOGE", "DOGE", "$ 1", ".", ",", 2)
	money.AddCurrency("USDC", "USDC", "$ 1", ".", ",", 2)
	money.AddCurrency("USDT", "USDT", "$ 1", ".", ",", 2)
	money.AddCurrency("DAI", "DAI", "$ 1", ".", ",", 2)
}
