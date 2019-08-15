package theforcechain

import (
	"github.com/theforceprotocolgroup/theforcechain/x/theforcechain/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewMsgSetOrder = types.NewMsgSetOrder
	NewOrder       = types.NewOrder
	ModuleCdc      = types.ModuleCdc
	RegisterCodec  = types.RegisterCodec
)

type (
	MsgSetOrder = types.MsgSetOrder
	Order       = types.Order
	QueryResIds = types.QueryResIds
)
