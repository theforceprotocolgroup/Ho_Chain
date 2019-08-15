package theforcechain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	OrderRecords []Order `json:"order_records"`
}

func NewGenesisState(orderRecords []Order) GenesisState {
	return GenesisState{OrderRecords: nil}
}

func ValidateGenesis(data GenesisState) error {
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		OrderRecords: []Order{},
	}
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.OrderRecords {
		keeper.SetOrder(ctx, record.Id, record)
	}
	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []Order
	iterator := k.GetIdsIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		id := string(iterator.Key())
		var order Order
		order = k.GetOrder(ctx, id)
		records = append(records, order)
	}
	return GenesisState{OrderRecords: records}
}
