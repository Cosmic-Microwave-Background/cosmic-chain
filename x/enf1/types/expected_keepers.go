package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankexported "github.com/cosmos/cosmos-sdk/x/bank/exported"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
)

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	SetBalances(ctx sdk.Context, addr sdk.AccAddress, balances sdk.Coins) error
	LockedCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins

	GetSupply(ctx sdk.Context) bankexported.SupplyI

	SendCoinsFromModuleToModule(ctx sdk.Context, senderPool, recipientPool string, amt sdk.Coins) error
	UndelegateCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	DelegateCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error

	BurnCoins(ctx sdk.Context, name string, amt sdk.Coins) error
}

type StakingKeeper interface {
	StakingTokenSupply(ctx sdk.Context) sdk.Int
	BondedRatio(ctx sdk.Context) sdk.Dec
	GetDelegation(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) (delegation types.Delegation, found bool)
	GetAllDelegatorDelegations(ctx sdk.Context, delegator sdk.AccAddress) []types.Delegation
}


type AccountKeeper interface {
	NewAccount(sdk.Context, authTypes.AccountI) authTypes.AccountI
	NewAccountWithAddress(ctx sdk.Context, addr sdk.AccAddress) authTypes.AccountI

	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authTypes.AccountI
	GetAllAccounts(ctx sdk.Context) []authTypes.AccountI
	SetAccount(ctx sdk.Context, acc authTypes.AccountI)

	IterateAccounts(ctx sdk.Context, process func(authTypes.AccountI) bool)

	ValidatePermissions(macc authTypes.ModuleAccountI) error

	GetModuleAddress(moduleName string) sdk.AccAddress
	GetModuleAddressAndPermissions(moduleName string) (addr sdk.AccAddress, permissions []string)
	GetModuleAccountAndPermissions(ctx sdk.Context, moduleName string) (authTypes.ModuleAccountI, []string)
	GetModuleAccount(ctx sdk.Context, moduleName string) authTypes.ModuleAccountI
	SetModuleAccount(ctx sdk.Context, macc authTypes.ModuleAccountI)
}

