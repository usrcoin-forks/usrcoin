package main

import (
	"fmt"
	"io/ioutil"
	"os"

	// "github.com/usrcoin-forks/usrcoin"
	"github.com/usrcoin-forks/usrcoin/lib/btc"
	"github.com/usrcoin-forks/usrcoin/lib/others/utils"
)

func main() {
	const usrcoinVersion = "0.0.1"
	fmt.Println("Gocoin FetchTx version", usrcoinVersion)

	if len(os.Args) < 2 {
		fmt.Println("Specify transaction id on the command line (MSB).")
		return
	}

	txid := btc.NewUint256FromString(os.Args[1])
	if txid == nil {
		println("Incorrect transaction ID")
		return
	}

	rawtx := utils.GetTxFromWeb(txid)
	if rawtx == nil {
		fmt.Println("Error fetching the transaction")
	} else {
		ioutil.WriteFile(txid.String()+".tx", rawtx, 0666)
	}
}
