package keeper

import (
	"github.com/enflow.io/enf1/x/enf1/types"
)

var _ types.QueryServer = Keeper{}
