package cmd

import (
	"os"

	"github.com/mercury-protocol/data-dao-helper/pkg"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// 12345678912345678912345678912345
var encryptCmd = &cobra.Command{
	Use:   "encrypt [key, path]",
	Short: "encrypts the given file with the given key",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("Running encryption")
		key := []byte(args[0])
		file, err := os.ReadFile(args[1])
		if err != nil {
			log.Fatal(err)
		}

		encoded, err := pkg.Encrypt(key, file)
		if err != nil {
			log.Fatal(err)
		}
		if err := os.WriteFile("encrpyted.bin", encoded, 0644); err != nil {
			log.Fatal(err)
		}
		log.Info("Encryption done")
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)
}
