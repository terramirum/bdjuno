package handlers

import (
	"fmt"

	"github.com/forbole/bdjuno/v5/modules/actions/types"

	"github.com/rs/zerolog/log"
)

func AccountBalanceHandler(ctx *types.Context, payload *types.Payload) (interface{}, error) {
	log.Debug().Str("address", payload.GetAddress()).
		Int64("height", payload.Input.Height).
		Msg("executing account balance action")

	height, err := ctx.GetHeight(payload)
	if err != nil {
		return nil, err
	}

	balance, err := ctx.Sources.BankSource.GetAccountBalance(payload.GetAddress(), height)
	if err != nil {
		return nil, fmt.Errorf("error while getting account balance: %s", err)
	}

	wasmBalance, err := ctx.Sources.WasmSource.GetAccountBalance(payload.GetAddress(), height)
	if err != nil {
		return nil, fmt.Errorf("error while getting account contract balance: %s", err)
	}

	coins := types.ConvertCoins(balance)
	wasmCoins := types.ConvertCoins(wasmBalance)
	coins = append(coins, wasmCoins...)

	return types.Balance{
		Coins: coins,
	}, nil
}
