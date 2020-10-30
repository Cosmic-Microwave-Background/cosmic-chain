package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/enflow.io/enf1/x/enf1/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
)

func (k Keeper) CreateAction(ctx sdk.Context, action types.MsgAction) {
	print("IM KEEPER")
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionKey))
	b := k.cdc.MustMarshalBinaryBare(&action)
	store.Set(types.KeyPrefix(types.ActionKey), b)
}

func (k Keeper) GetAllAction(ctx sdk.Context) (msgs []types.MsgAction) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.ActionKey))

	defer iterator.Close()

	//k.bankKeeper.DelegateCoinsFromAccountToModule(ctx, )

	for ; iterator.Valid(); iterator.Next() {
		var msg types.MsgAction
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
        msgs = append(msgs, msg)
	}

    return
}
