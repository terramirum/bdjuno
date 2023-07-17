package nft

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/nft"

	"github.com/forbole/bdjuno/v5/database"

	"github.com/forbole/juno/v5/modules"
)

var (
	_ modules.Module = &Module{}
)

// Module represent x/wasm module
type Module struct {
	cdc codec.Codec
	db  *database.Db
}

// NewModule returns a new Module instance
func NewModule(cdc codec.Codec, db *database.Db) *Module {
	return &Module{
		cdc: cdc,
		db:  db,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return nft.ModuleName
}
