package cmd

import (
	"fmt"
	"strings"

	"github.com/mikeymop/whodis/qr"
	"github.com/mikeymop/whodis/uns"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(QrCmd)
	RootCmd.Flags().StringVarP(&qrToken, "qr", "q", "", "Prints qr code for the specified token")
	RootCmd.Flags().StringVarP(&token, "token", "t", "", "Prints address for the specified token")
}

func Execute() {
	RootCmd.Execute()
}

var (
	qrToken string
	token   string
)

var RootCmd = &cobra.Command{
	Use:   "whodis",
	Args:  cobra.RangeArgs(1, 1),
	Short: "whodis, whois for the decentralized web.",
	Long: `whodis is a web3 domain look-up utility written
			with love by mikeymop in Go.
			Complete documentation is available at https://githubhub.com/mikymop/whodis`,
	Run: func(cmd *cobra.Command, args []string) {
		domain := args[0]

		if token != "" {
			addr := uns.ResolveAddress(domain, strings.ToUpper(token))
			fmt.Printf("%s Address for %s: %s\n", strings.ToUpper(token), domain, addr)
		}

		if qrToken != "" {
			addr := uns.ResolveAddress(domain, strings.ToUpper(qrToken))
			qr.Generate(addr)
			fmt.Printf("%s address for %s: %s\n", strings.ToUpper(qrToken), domain, addr)
		}

		if qrToken == "" && token == "" {
			fmt.Printf("rootCmd args: %s\n", args)
			records := uns.ShowRecords(domain)

			for i, record := range records {
				fmt.Println(i, record)
			}
		}

	},
}
