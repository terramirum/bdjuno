package local

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (s Source) GetAccountBalance(address string, height int64) ([]sdk.Coin, error) {
	coins := []sdk.Coin{}
	return coins, nil
}

