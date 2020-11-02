package cli

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/enflow.io/enf1/x/enf1/types"
)

func CmdCreateAction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-action [receiver] [amount] [denom]",
		Short: "Creates a new action",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsReceiver := sdk.AccAddress(args[0])
			argsAmount := string(args[1])
			argsDenom := string(args[2])
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgAction(clientCtx.GetFromAddress(), argsReceiver, string(argsAmount), string(argsDenom))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
