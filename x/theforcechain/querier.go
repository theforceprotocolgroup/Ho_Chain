package theforcechain

import (
	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// query endpoints supported by the theforcechain Querier
const (
	QueryIds = "ids"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryIds:
			return queryIds(ctx, req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown theforcechain query endpoint")
		}
	}
}

func queryIds(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	var idsList QueryResIds

	iterator := keeper.GetIdsIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		idsList = append(idsList, string(iterator.Key()))
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, idsList)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}
