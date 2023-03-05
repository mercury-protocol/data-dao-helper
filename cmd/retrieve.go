package cmd

import (
	"fmt"

	"github.com/mercury-protocol/data-dao-helper/pkg"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var retrieveCmd = &cobra.Command{
	Use:   "retrieve [cid]",
	Short: "Retrieves data from the filecoin network",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := pkg.Retrieve(args[0], fmt.Sprintf("%s.txt", args[0])); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(retrieveCmd)
}
