package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = ModuleName // this was defined in your key.go file

// MsgSetOrder defines a SetOrder message
type MsgSetOrder struct {
	Id        string         `json:"id"`
	Borrower  string         `json:"borrower"`
	Lender    string         `json:"lender"`
	TokenGet  sdk.Coin       `json:"tokenGet"`
	TokenGive sdk.Coin       `json:"tokenGive"`
	Owner     sdk.AccAddress `json:"owner"`
}

// NewMsgSetName is a constructor function for MsgSetName
func NewMsgSetOrder(id, borrower, lender string, tokenGet, tokenGive sdk.Coin, owner sdk.AccAddress) MsgSetOrder {
	return MsgSetOrder{
		Id:        id,
		Borrower:  borrower,
		Lender:    lender,
		TokenGet:  tokenGet,
		TokenGive: tokenGive,
		Owner:     owner,
	}
}

// Route should return the name of the module
func (msg MsgSetOrder) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSetOrder) Type() string { return "set_order" }

// ValidateBasic runs stateless checks on the message
func (msg MsgSetOrder) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.Borrower) == 0 || len(msg.Lender) == 0 {
		return sdk.ErrUnknownRequest("Borrower and/or Lender cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSetOrder) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgSetOrder) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}
