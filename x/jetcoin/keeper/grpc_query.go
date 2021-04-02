package keeper

import (
	"github.com/compamint/jetcoin/x/jetcoin/types"
)

var _ types.QueryServer = Keeper{}
