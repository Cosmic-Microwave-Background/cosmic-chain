package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgAction{}

func NewMsgAction(creator sdk.AccAddress, receiver sdk.AccAddress, amount string, denom string) *MsgAction {
	return &MsgAction{
		Id:       uuid.New().String(),
		Creator:  creator,
		Receiver: receiver,
		Amount:   amount,
		Denom: denom,
	}
}

func (msg *MsgAction) Route() string {
	return RouterKey
}

func (msg *MsgAction) Type() string {
	return "CreateAction"
}

func (msg *MsgAction) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg *MsgAction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAction) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
