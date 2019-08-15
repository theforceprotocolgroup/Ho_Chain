package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
	"github.com/theforceprotocolgroup/theforcechain/x/theforcechain/types"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	forcechainQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the forcechain module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	forcechainQueryCmd.AddCommand(client.GetCommands(
		GetCmdIds(storeKey, cdc),
	)...)
	return forcechainQueryCmd
}

// GetCmdIds queries a list of all ids
func GetCmdIds(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "ids",
		Short: "ids",
		// Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/forcechain", queryRoute), nil)
			if err != nil {
				fmt.Printf("could not get query forcechain\n")
				return nil
			}

			var out types.QueryResIds
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
