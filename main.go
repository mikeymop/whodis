package main

import (
	"github.com/mikeymop/whodis/cmd"
	// "github.com/mikeymop/whodis/qr"
	// "github.com/mikeymop/whodis/uns"
)

func main() {
	// addr := "mikeymop.x"
	// log.Printf("Records for %s %v", addr, uns.ShowRecords(addr))
	// log.Printf("Address for %s: %s", addr, uns.ResolveAddress(addr, "ETH"))
	// qr.Generate(uns.ResolveAddress(addr, "ETH"))

	cmd.RootCmd.Execute()

}
