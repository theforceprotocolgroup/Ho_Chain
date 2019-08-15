package theforcechain

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler returns a handler for "theforcechain" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgSetOrder:
			return handleMsgSetOrder(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized theforcechain Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// Handle a message to set name
func handleMsgSetOrder(ctx sdk.Context, keeper Keeper, msg MsgSetOrder) sdk.Result {
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.Id)) { // Checks if the the msg sender is the same as the current owner
		return sdk.ErrUnauthorized("Incorrect Owner").Result() // If not, throw an error
	}
	var order Order
	order.Id = msg.Id
	order.Borrower = msg.Borrower
	order.Lender = msg.Lender
	order.TokenGet = msg.TokenGet
	order.TokenGive = msg.TokenGive
	order.Owner = msg.Owner

	keeper.SetOrder(ctx, msg.Id, order) // If so, set the name to the value specified in the msg.
	return sdk.Result{}                 // return
}
