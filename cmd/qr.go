package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var QrCmd = &cobra.Command{
	Use:   "whodis [DOMAIN] --qr [CURRENCY]",
	Short: "Generates a qr code for the requested currency.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("Test qrCmd")
	},
}
