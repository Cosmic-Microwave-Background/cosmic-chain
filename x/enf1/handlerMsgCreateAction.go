package enf1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/enflow.io/enf1/x/enf1/types"
	"github.com/enflow.io/enf1/x/enf1/keeper"
)

func handleMsgCreateAction(ctx sdk.Context, k keeper.Keeper, action *types.MsgAction) (*sdk.Result, error) {
	k.CreateAction(ctx, *action)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
