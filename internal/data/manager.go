package data

var Manager = CryptoManager{}

func init() {
    Manager.cryptoDataSources = []Collectible{
        &TronWallet,
        &HuobiWallet,
        &KucoinWallet,
    }
}

type CryptoManager struct {
	cryptoDataSources []Collectible
}

func (m *CryptoManager) Collect() {
	for _, value := range m.cryptoDataSources {
		value.Collect()
        value.Format()
	}
}

