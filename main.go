package main

import (
	"log"

	"github.com/mikeymop/whodis/qr"
	"github.com/mikeymop/whodis/uns"
)

func main() {
	log.Printf("Records for mikeymop.crypto %v", uns.ShowRecords("mikeymop.crypto"))
	log.Printf("Address for mikeymop.x: %v", uns.ResolveAddress("mikeymop.x", "BTC"))
	qr.Generate(uns.ResolveAddress("mikeymop.crypto", "ETH"))
}
