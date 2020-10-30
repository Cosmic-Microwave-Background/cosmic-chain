package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/enflow.io/enf1/x/enf1/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AllAction(c context.Context, req *types.QueryAllActionRequest) (*types.QueryAllActionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var actions []*types.MsgAction
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	actionStore := prefix.NewStore(store, types.KeyPrefix(types.ActionKey))

	pageRes, err := query.Paginate(actionStore, req.Pagination, func(key []byte, value []byte) error {
		var action types.MsgAction
		if err := k.cdc.UnmarshalBinaryBare(value, &action); err != nil {
			return err
		}

		actions = append(actions, &action)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllActionResponse{Action: actions, Pagination: pageRes}, nil
}
