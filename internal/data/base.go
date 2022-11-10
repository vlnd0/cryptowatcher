package data

type Collectible interface {
	Collect()
	Format()
}

type CryptoDataSource struct {
	Name    string
	Balance BalanceType
}

type BalanceType map[string]float32

func (b BalanceType) Format() {
    for key, value := range b {
        println(key, ":", value)
    }
}