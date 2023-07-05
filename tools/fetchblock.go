package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/usrcoin-forks/usrcoin/lib/btc"
	"github.com/usrcoin-forks/usrcoin/lib/others/utils"
)

func main() {
	const usrcoinVersion = "0.0.1"
	fmt.Println("Gocoin FetchBlock version", usrcoinVersion)

	if len(os.Args) < 2 {
		fmt.Println("Specify block hash on the command line (MSB).")
		return
	}

	hash := btc.NewUint256FromString(os.Args[1])
	bl := utils.GetBlockFromWeb(hash)
	if bl == nil {
		fmt.Println("Error fetching the block")
	} else {
		ioutil.WriteFile(bl.Hash.String()+".bin", bl.Raw, 0666)
	}
}
