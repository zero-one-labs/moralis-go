package moralis

import (
	"github.com/zero-one-labs/moralis-go/moralis/evm"
	"github.com/zero-one-labs/moralis-go/moralis/solana"
)

func NewEVM(apiKey string) *evm.MoralisAPI {
	return &evm.MoralisAPI{
		APIKey:  apiKey,
		BaseURL: "https://deep-index.moralis.io/api/v2",
	}
}

func NewSOL(apiKey string) *solana.MoralisAPI {
	return &solana.MoralisAPI{
		APIKey:  apiKey,
		Network: "mainnet",
		BaseURL: "https://solana-gateway.moralis.io",
	}
}

func EVMSupportedChain(chainID string) bool {
	switch chainID {
	case "eth", "polygon", "bsc", "avalanche":
		return true
	default:
		return false
	}
}
