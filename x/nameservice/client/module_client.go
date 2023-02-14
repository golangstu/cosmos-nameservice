package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	"github.com/tendermint/go-amino"
)

type ModuleClient struct {
	storeKey string
	cdc      *amino.Codec
}

func NewModuleClient(storeKey string, cdc *amino.Codec) ModuleClient {
	return ModuleClient{storeKey, cdc}
}

func (mc ModuleClient) GetQueryCmd() *cobra.Command {
	namesvcQueryCmd:=&cobra.Command{
		Use: "nameservice",
		Short: "Querying commands for the nameservice module",
	}

	namesvcQueryCmd.AddCommand(client.GetCommands(
		cli.GetCmdRe
		))
}