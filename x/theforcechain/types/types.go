package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Order struct {
	Id string `json:"id"`

	Borrower string `json:"borrower"`
	Lender   string `json:"lender"`

	TokenGet  sdk.Coin `json:"tokenGet"`
	TokenGive sdk.Coin `json:"tokenGive"`

	Owner sdk.AccAddress `json:"owner"`
}

func NewOrder(id, borrower, lender string, tokenGet, tokenGive sdk.Coin, owner sdk.AccAddress) Order {
	return Order{
		Id:        id,
		Borrower:  borrower,
		Lender:    lender,
		TokenGet:  tokenGet,
		TokenGive: tokenGive,
		Owner:     owner,
	}
}

// implement fmt.Stringer
func (o Order) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Id: %s
Borrower: %s
Lender: %s
TokenGet: %v
TokenGive: %v
Owner: %s`, o.Id, o.Borrower, o.Lender, o.TokenGet, o.TokenGive, o.Owner))
}
