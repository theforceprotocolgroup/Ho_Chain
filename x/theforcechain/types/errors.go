package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	DefaultCodespace sdk.CodespaceType = ModuleName

	CodeOrderDoesNotExist sdk.CodeType = 101
)

func ErrOrderDoesNotExist(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeOrderDoesNotExist, "Order does not exist")
}
