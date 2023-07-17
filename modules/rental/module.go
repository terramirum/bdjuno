package rental

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/terramirum/mirumd/x/rental/types"

	"github.com/forbole/bdjuno/v5/database"

	"github.com/forbole/juno/v5/modules"
)

var (
	_ modules.Module = &Module{}
)

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
	return types.ModuleName
}
