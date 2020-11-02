package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/enflow.io/enf1/x/enf1/types"
)

type (
	Keeper struct {
		cdc        codec.Marshaler
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		bankKeeper types.BankKeeper
		stakingKeeper types.StakingKeeper
		accountKeeper types.AccountKeeper
	}
)

func NewKeeper(cdc codec.Marshaler, storeKey, memKey sdk.StoreKey, bankKeeper types.BankKeeper,
	stakingKeeper types.StakingKeeper, accountKeeper types.AccountKeeper ) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
		bankKeeper: bankKeeper,
		stakingKeeper: stakingKeeper,
		accountKeeper: accountKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
