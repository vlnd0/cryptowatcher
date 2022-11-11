package currencies

import "github.com/Rhymond/go-money"

func init() {
    money.AddCurrency("BTC", "BTC", "$ 1", ".", ",", 6)
    money.AddCurrency("ETH", "ETH", "$ 1", ".", ",", 6)
    money.AddCurrency("SOL", "SOL", "$ 1", ".", ",", 6)
    money.AddCurrency("TRX", "TRX", "$ 1", ".", ",", 6)
    money.AddCurrency("USDC", "USDC", "$ 1", ".", ",", 2)
    money.AddCurrency("USDT", "USDT", "$ 1", ".", ",", 2)
}
