package remote

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var wasmContracts []WasmContract
var lastQueryTime time.Time
var mutex sync.Mutex

type WasmContract struct {
	ContractAddress string
	Denom           string
}

type SmartContractBalanceRequest struct {
	Balance struct {
		Address string `json:"address"`
	} `json:"balance"`
}

type SmartContractBalanceResponse struct {
	Balance string `json:"balance"`
}

type TokenInfo struct {
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	Decimals    int    `json:"decimals"`
	TotalSupply string `json:"total_supply"`
}

func (s Source) GetAccountBalance(address string, height int64) ([]sdk.Coin, error) {
	tokens, err := s.getWasmTokens()
	if err != nil {
		return nil, err
	}

	coins := []sdk.Coin{}
	for _, v := range tokens {
		coin, err := s.getTokenBalance(v.ContractAddress, address, v.Denom)
		if err != nil {
			return nil, err
		}
		coins = append(coins, *coin)
	}

	return coins, nil
}

func (s Source) getTokenBalance(contractAddress, walletAddress, denom string) (*sdk.Coin, error) {
	balanceRequest := SmartContractBalanceRequest{}
	balanceRequest.Balance.Address = walletAddress

	data, err := json.Marshal(balanceRequest)
	if err != nil {
		return nil, err
	}

	contractState := wasmtypes.QuerySmartContractStateRequest{
		Address:   contractAddress,
		QueryData: data,
	}

	queryBalanceResponse, err := s.wasmClient.SmartContractState(context.Background(), &contractState)
	if err != nil {
		return nil, err
	}

	var balanceResponse SmartContractBalanceResponse
	json.Unmarshal(queryBalanceResponse.Data, &balanceResponse)
	balance, err := strconv.ParseInt(balanceResponse.Balance, 10, 64)
	if err != nil {
		return nil, err
	}

	return &sdk.Coin{
		Amount: sdk.NewInt(balance),
		Denom:  denom,
	}, nil
}

func (s *Source) getWasmTokens() ([]WasmContract, error) {
	// Check if at least one hour has passed since the last query
	if lastQueryTime.IsZero() || time.Since(lastQueryTime) > time.Hour {
		mutex.Lock()
		defer mutex.Unlock()

		// Check again inside the lock to avoid duplicate queries by concurrent threads
		if lastQueryTime.IsZero() || time.Since(lastQueryTime) > time.Hour {
			wasmContractInfos, err := s.db.GetWasmTokens()
			if err != nil {
				return nil, err
			}

			// Clear the existing wasmContracts
			wasmContracts = nil

			// Iterate over the contract infos and fill wasmContracts with the relevant data
			for _, contractInfo := range wasmContractInfos {
				var contractState map[string]string
				err := json.Unmarshal([]byte(contractInfo.ContractStates), &contractState)
				if err != nil {
					return nil, fmt.Errorf("error unmarshaling contract state JSON: %s", err)
				}

				var tokenInfo TokenInfo
				err = json.Unmarshal([]byte(contractState["token_info"]), &tokenInfo)
				if err != nil {
					return nil, fmt.Errorf("error unmarshaling contract token info JSON: %s", err)
				}

				wasmContracts = append(wasmContracts, WasmContract{
					ContractAddress: contractInfo.ContractAddress,
					Denom:           tokenInfo.Symbol,
				})
			}

			// Update the last query time
			lastQueryTime = time.Now().UTC()
		}
	}

	return wasmContracts, nil
}
