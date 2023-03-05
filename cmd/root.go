package cmd

import (
	"os"

	"github.com/mercury-protocol/data-dao-helper/internal"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var version = "0.0.1"

var rootCmd = &cobra.Command{
	Use:     "mcydao",
	Version: version,
	Short:   "Mercury Data DAO helper tools CLI",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		internal.LoadConfig()
		internal.ConfigLogger()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal("Error starting CLI: ", err)
		os.Exit(1)
	}
}
