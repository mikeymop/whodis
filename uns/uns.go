package uns

import (
	"encoding/hex"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/unstoppabledomains/resolution-go"
)

const ETH_MAINNET = "0xfEe4D4F0aDFF8D84c12170306507554bC7045878"
const POLY_MAINNET = "0xA3f32c8cd786dc089Bd1fC175F2707223aeE5d00"

func UDToHash(domain string) string {
	arr := strings.Split(domain, ".")
	label, ext := arr[0], arr[1]
	labelHash := common.BytesToHash([]byte(domain))
	extHash := common.BytesToHash([]byte(ext))
	log.Printf("label: %v ext: %v", label, ext)

	return hex.EncodeToString(crypto.Keccak256(labelHash[:], extHash[:]))
}

func Locate(domain string) {
	uns, _ := resolution.NewUnsBuilder().Build()
	res, _ := uns.Resolver("brad.crypto")
	log.Println("Resolver game me this", res)
}

func ShowRecords(domainName string) map[string]string {
	uns, err := resolution.NewUnsBuilder().Build()
	if err != nil {
		log.Fatalf("Error connecting to ethereum %v", err)
	}

	records, err := uns.AllRecords(domainName)
	if err != nil {
		log.Fatalf("Error getting records for %v %v", domainName, err)
	}

	return records
}

func ResolveAddress(domainName string, token string) string {
	uns, _ := resolution.NewUnsBuilder().Build()
	address, _ := uns.Addr(domainName, token)

	return address
}

func resolve() {
	// log.Print(udToHash("brad.crypto"))
	// conn, err := ethclient.Dial("https://mainnet.infura.io/")
	// if err != nil {
	// 	log.Fatalf("Failed to connect to Ethereum network %v", err)
	// }

}
