package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/theforceprotocolgroup/theforcechain/x/theforcechain/types"
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	forcechainTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "theforcechain transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	forcechainTxCmd.AddCommand(client.PostCommands(
		GetCmdSetOrder(cdc),
	)...)

	return forcechainTxCmd
}

// GetCmdSetName is the CLI command for sending a SetName transaction
func GetCmdSetOrder(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-order [id] [borrower] [lender] [tokenGet] [tokenGive]",
		Short: "set the value associated with a id that you own",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			// if err := cliCtx.EnsureAccountExists(); err != nil {
			// 	return err
			// }

			tokenGet, err1 := sdk.ParseCoin(args[3])
			if err1 != nil {
				return err1
			}

			tokenGive, err2 := sdk.ParseCoin(args[4])
			if err2 != nil {
				return err2
			}

			msg := types.NewMsgSetOrder(args[0], args[1], args[2], tokenGet, tokenGive, cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			// return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, msgs)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
