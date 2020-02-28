package cmd

import (
	"fmt"
	"log"

	"encoding/base64"

	"github.com/spf13/cobra"

	"github.com/terra-project/core/app"
	"github.com/terra-project/core/x/auth"
)

func init() {
	decodeCmd.AddCommand(decodeTxCmd)
	rootCmd.AddCommand(decodeCmd)
}

var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Runs amino decoding",
}

var decodeTxCmd = &cobra.Command{
	Use:   "tx [amino_encoded_tx]",
	Short: "decode amino encoded tx",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		txBytes, err := base64.StdEncoding.DecodeString(args[0])
		if err != nil {
			log.Fatal(err.Error())
			return
		}

		cdc := app.MakeCodec()

		var stdTx auth.StdTx
		err = cdc.UnmarshalBinaryLengthPrefixed(txBytes, &stdTx)
		if err != nil {
			log.Fatal(err.Error())
			return
		}

		bz, err := cdc.MarshalJSON(stdTx)

		fmt.Println(string(bz))
	},
}
