package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetName{}

type MsgSetName struct {
  Name string `json:"name"`
  Value string `json:"value"`
  Owner sdk.AccAddress `json:"owner"`
}

func NewMsgSetName(name string, value string, owner sdk.AccAddress) MsgSetName {
  return MsgSetName{
    Name:  name,
    Value: value,
    Owner: owner,
  }
}

func (msg MsgSetName) Route() string {
  return RouterKey
}

func (msg MsgSetName) Type() string {
  return "SetWhois"
}

func (msg MsgSetName) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{msg.Owner}
}

func (msg MsgSetName) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgSetName) ValidateBasic() error {
  if msg.Owner.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }

  if len(msg.Name) == 0 || len(msg.Value) == 0 {
    return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Name and/or Value cannot be empty")
  }

  return nil
}