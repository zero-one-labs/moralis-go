package solana

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MoralisAPI struct {
	APIKey  string
	Network string
	BaseURL string
}

type BalanceResponse struct {
	Solana   string `json:"solana"`
	Lamports string `json:"lamports"`
}

func (m *MoralisAPI) Balance(address string) (BalanceResponse, error) {
	url := fmt.Sprintf("%s/account/%s/%s/balance", m.BaseURL, m.Network, address)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return BalanceResponse{}, err
	}

	req.Header.Set("X-API-Key", m.APIKey)
	req.Header.Set("accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return BalanceResponse{}, err
	}

	var result BalanceResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return BalanceResponse{}, err
	}

	return result, nil
}

type TokenBalanceResponse struct {
	AssociatedTokenAddress string `json:"associatedTokenAddress"`
	Mint                   string `json:"mint"`
	Name                   string `json:"name"`
	Symbol                 string `json:"symbol"`
	Amount                 string `json:"amount"`
	AmountRaw              string `json:"amountRaw"`
	Decimals               string `json:"decimals"`
}

func (m *MoralisAPI) TokenBalance(address string) ([]TokenBalanceResponse, error) {
	url := fmt.Sprintf("%s/account/%s/%s/tokens", m.BaseURL, m.Network, address)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-API-Key", m.APIKey)
	req.Header.Set("accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var result []TokenBalanceResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type NFT struct {
	AssociatedTokenAddress string `json:"associatedTokenAddress"`
	Mint                   string `json:"mint"`
	Name                   string `json:"name"`
	Symbol                 string `json:"symbol"`
}

func (m *MoralisAPI) NFTsInWallet(address string) ([]NFT, error) {
	url := fmt.Sprintf("%s/account/%s/%s/nft", m.BaseURL, m.Network, address)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-API-Key", m.APIKey)
	req.Header.Set("accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var result []NFT
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type NFTOwner struct {
	Address  string `json:"address"`
	Verified int16  `json:"verified"`
	Share    int16  `json:"share"`
}

type Metaplex struct {
	MetadataURI          string     `json:"metadataUri"`
	MasterEdition        bool       `json:"masterEdition"`
	IsMutable            bool       `json:"isMutable"`
	PrimarySaleHappened  int16      `json:"primarySaleHappened"`
	SellerFeeBasisPoints int64      `json:"sellerFeeBasisPoints"`
	UpdateAuthority      string     `json:"updateAuthority"`
	Owners               []NFTOwner `json:"owners"`
}

type NFTMetaDataResponse struct {
	Mint     string   `json:"mint"`
	Standard string   `json:"standard"`
	Name     string   `json:"name"`
	Symbol   string   `json:"symbol"`
	Metaplex Metaplex `json:"metaplex"`
}

func (m *MoralisAPI) NFTMetaData(address string) (NFTMetaDataResponse, error) {
	url := fmt.Sprintf("%s/nft/%s/%s/metadata", m.BaseURL, m.Network, address)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return NFTMetaDataResponse{}, err
	}

	req.Header.Set("X-API-Key", m.APIKey)
	req.Header.Set("accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return NFTMetaDataResponse{}, err
	}

	var result NFTMetaDataResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return NFTMetaDataResponse{}, err
	}

	return result, nil
}
