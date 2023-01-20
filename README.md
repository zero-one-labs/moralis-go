# Moralis Go

A Go wrapper for [Moralis Web3 APIs](https://docs.moralis.io/reference/introduction)

You need to get an API key from [Moralis](https://moralis.io)

## Usage

### Add the module to your go project

```sh
go get github.com/zero-one-labs/moralis-go/moralis
```

### Sample Code

```go
package main

import (
	"fmt"

	"github.com/zero-one-labs/moralis-go/moralis"
)

const MORALIS_API_KEY = "YOUR_MORALIS_API_KEY"

func main() {
	/*********
	* ETH / EVM
	**********/
	evm := moralis.NewEVM(MORALIS_API_KEY)

	// get eth address balance
	balance, err := evm.Balance("0xB7C9C2560828F585518D0283656A680AA919DE12", "eth")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", balance.Balance)

	// eth address + token address
	tokenBalance, err := evm.TokenBalance("0xB7C9C2560828F585518D0283656A680AA919DE12", "0x1cf0b4989ad877438aa2e7804192e2ae83a94081", "eth")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", tokenBalance)

	// check if joe random wallet has BAYC token
	ethNFT, err := evm.NFTsInWallet("0x35d06c66d3b4ac4a2b554721e5aedb58b0e6858a", "0xCF4F43EC9E61C0E8723015831B27AF90908EE7C1", "eth")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", ethNFT)

	/*********
	* SOL
	**********/
	sol := moralis.NewSOL(MORALIS_API_KEY)
	solBalance, err := sol.Balance("7Vdeq3nAj6RMwBtBxAvH93WogdGQdfhkLHbLP5nTTLcg")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", solBalance)

	solTokenBalance, err := sol.TokenBalance("7Vdeq3nAj6RMwBtBxAvH93WogdGQdfhkLHbLP5nTTLcg")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", solTokenBalance)

	solNFT, err := sol.NFTsInWallet("7Vdeq3nAj6RMwBtBxAvH93WogdGQdfhkLHbLP5nTTLcg")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", solNFT)

	data, err := sol.NFTMetaData("4NRfLVSoWuChSs949EnBmwLBMidspCMB33xaTiL5tEX8")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", data)
}
```

## Implemented Endpoints

- EVM API
    - [Get Native Balance by Wallet](https://docs.moralis.io/web3-data-api/reference/get-native-balance)
    - [Get ERC20 token balance by wallet](https://docs.moralis.io/reference/getwallettokenbalances)
    - [Get NFTs by wallet](https://docs.moralis.io/reference/getwalletnfts)
- SOLANA API
    - [Get native balance by wallet](https://docs.moralis.io/reference/solbalance)
    - [Get token balance by wallet](https://docs.moralis.io/reference/getspl)
    - [Get NFTs by wallet](https://docs.moralis.io/reference/getsolnfts)
    - [Get NFT metadata](https://docs.moralis.io/reference/getsolnftmetadata)

## TODO

- Add GoDocs
- Add Tests
- Implement more API endpoints
