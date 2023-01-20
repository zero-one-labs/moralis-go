package evm

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MoralisAPI struct {
	APIKey  string
	BaseURL string
}

type BalanceResponse struct {
	Balance string `json:"balance"`
}

func (m *MoralisAPI) Balance(address string, chain string) (BalanceResponse, error) {
	url := fmt.Sprintf("%s/%s/balance", m.BaseURL, address)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return BalanceResponse{}, err
	}

	q := req.URL.Query()
	q.Add("chain", chain)
	req.URL.RawQuery = q.Encode()

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
	TokenAddress string `json:"token_address"`
	Name         string `json:"name"`
	Symbol       string `json:"symbol"`
	Logo         string `json:"logo"`
	Thumbnail    string `json:"thumbnail"`
	Decimals     int16  `json:"decimals"`
	Balance      string `json:"balance"`
}

func (m *MoralisAPI) TokenBalance(walletAddress string, tokenAddress string, chain string) (TokenBalanceResponse, error) {
	url := fmt.Sprintf("%s/%s/erc20", m.BaseURL, walletAddress)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return TokenBalanceResponse{}, err
	}

	q := req.URL.Query()
	q.Add("chain", chain)
	q.Add("token_addresses", tokenAddress)
	req.URL.RawQuery = q.Encode()

	req.Header.Set("X-API-Key", m.APIKey)
	req.Header.Set("accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return TokenBalanceResponse{}, err
	}

	var result []TokenBalanceResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return TokenBalanceResponse{}, err
	}

	if len(result) == 0 {
		return TokenBalanceResponse{}, fmt.Errorf("no token found")
	}

	return result[0], nil
}

type WalletNFTsResponse struct {
	Total    int64  `json:"total"`
	Page     int64  `json:"page"`
	PageSize int64  `json:"page_size"`
	Cursor   string `json:"cursor"`
	Result   []NFT  `json:"result"`
}

type NFT struct {
	TokenAddress      string `json:"token_address"`
	TokenID           string `json:"token_id"`
	Amount            string `json:"amount"`
	OwnerOf           string `json:"owner_of"`
	TokenHash         string `json:"token_hash"`
	BlockNumberMinted string `json:"block_number_minted"`
	BlockNumber       string `json:"block_number"`
	ContractType      string `json:"contract_type"`
	Name              string `json:"name"`
	Symbol            string `json:"symbol"`
	TokenURI          string `json:"token_uri"`
	Metadata          string `json:"metadata"`
	LastTokenURISync  string `json:"last_token_uri_sync"`
	LastMetadataSync  string `json:"last_metadata_sync"`
	MinterAddress     string `json:"minter_address"`
}

// TODO @mgcm
// The API returns all NFTs in a wallet for a given contract
// It has cursor support (one wallet can have many NFTs of the same collection) but ignoring that for now
// We just want to check if a wallet has at least one
func (m *MoralisAPI) NFTsInWallet(walletAddress string, contractAddress string, chain string) (WalletNFTsResponse, error) {
	url := fmt.Sprintf("%s/%s/nft", m.BaseURL, walletAddress)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return WalletNFTsResponse{}, err
	}

	q := req.URL.Query()
	q.Add("chain", chain)
	q.Add("token_addresses", contractAddress)
	req.URL.RawQuery = q.Encode()

	req.Header.Set("X-API-Key", m.APIKey)
	req.Header.Set("accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return WalletNFTsResponse{}, err
	}

	var result WalletNFTsResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return WalletNFTsResponse{}, err
	}

	return result, nil
}
