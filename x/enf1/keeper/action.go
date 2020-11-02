package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/enflow.io/enf1/x/enf1/types"
	"os"
	"strconv"

	//"runtime"
)

func (k Keeper) CreateAction(ctx sdk.Context, action types.MsgAction) error {

	defer func() { //catch or finally
		if err := recover(); err != nil { //catch
			fmt.Fprintf(os.Stderr, "Exception: %v\n", err)
			os.Exit(1)
		}
	}()


	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionKey))
	b := k.cdc.MustMarshalBinaryBare(&action)
	//bondDenom := sdk.DefaultBondDenom


	fmt.Printf("\n%v", action)
	fmt.Printf("amount: \n%v", action.Amount)
	intAmount, err := strconv.ParseInt(action.Amount,  10, 64)
	fmt.Printf("\n%v", intAmount)
	fmt.Print("------")

	if err != nil {
		fmt.Printf("%v", err)
		panic("cant convert")
	}
	amt := sdk.NewInt(intAmount)
	fmt.Printf("NEW AMOUNT \n%v", amt)
	denom := action.Denom
	fmt.Printf("DENOM \n%v", denom)

	balance := k.bankKeeper.GetBalance(ctx, action.Creator, denom)
	fmt.Printf("BALANCE \n%v", balance)

	coins := sdk.NewCoins(sdk.NewCoin(denom, amt))
	fmt.Printf("COINS \n%v", coins)

	fmt.Println("COINS", coins)



	//err := k.bankKeeper.DelegateCoinsFromAccountToModule(ctx, action.Creator, types.ModuleName, coins)
	//runtime.Breakpoint()
	print("BEFORE ERROR")
	if err := k.bankKeeper.DelegateCoinsFromAccountToModule(ctx, action.Creator, types.ModuleName, coins); err != nil {
		print("ERROR")
		fmt.Printf("\n%s\n", err)
		panic(err)
		return err
	}else{
		print("NO ERROR")
	}

	recipientAcc := k.accountKeeper.GetModuleAccount(ctx, types.ModuleName)
	//delegation, isFound := k.stakingKeeper.GetDelegation(ctx, action.Creator, sdk.ValAddress(recipientAcc.GetAddress()))
	delegations := k.stakingKeeper.GetAllDelegatorDelegations(ctx, action.Creator)

	fmt.Printf("\n%s\n", delegations)
	//fmt.Printf("\nisFound: %s\n", isFound)
	fmt.Printf("\naddress: %s\n", recipientAcc.GetAddress())
	//
	//fmt.Printf("%v\n", err)

	//if err := k.bankKeeper.DelegateCoinsFromAccountToModule(ctx, action.Creator, types.ModuleName, coins); err != nil {
	//	fmt.Printf("%v\n", err)
	//	print("ERROR")
	//}else{
	//	print("SUCCESS")
	//}

	fmt.Printf("%v\n", balance)
	print("it was balance")
	store.Set(types.KeyPrefix(types.ActionKey), b)
	return nil
}

func (k Keeper) GetAllAction(ctx sdk.Context) (msgs []types.MsgAction) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.ActionKey))

	defer iterator.Close()

	//k.bankKeeper.DelegateCoinsFromAccountToModule(ctx, msgs)

	for ; iterator.Valid(); iterator.Next() {
		var msg types.MsgAction
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}
