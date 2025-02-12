package utils

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/usrcoin-forks/usrcoin/lib/btc"
)

// https://blockchain.info/block/000000000000000000871f4f01a389bda59e568ead8d0fd45fc7cc1919d2666e?format=hex
// https://webbtc.com/block/0000000000000000000cdc0d2a9b33c2d4b34b4d4fa8920f074338d0dc1164dc.bin
// https://blockexplorer.com/api/rawblock/0000000000000000000cdc0d2a9b33c2d4b34b4d4fa8920f074338d0dc1164dc

// GetBlockFromExplorer downloads (and re-assembles) a raw block from blockexplorer.com.
func GetBlockFromExplorer(hash *btc.Uint256) (rawtx []byte) {
	url := "http://blockexplorer.com/api/rawblock/" + hash.String()
	r, er := http.Get(url)
	if er == nil {
		if r.StatusCode == 200 {
			defer r.Body.Close()
			c, _ := io.ReadAll(r.Body)
			var txx struct {
				Raw string `json:"rawblock"`
			}
			er = json.Unmarshal(c[:], &txx)
			if er == nil {
				rawtx, er = hex.DecodeString(txx.Raw)
			}
		} else {
			fmt.Println("blockexplorer.com StatusCode=", r.StatusCode)
		}
	}
	if er != nil {
		fmt.Println("blockexplorer.com:", er.Error())
	}
	return
}

// GetBlockFromWebBTC downloads a raw block from webbtc.com.
func GetBlockFromWebBTC(hash *btc.Uint256) (raw []byte) {
	url := "https://webbtc.com/block/" + hash.String() + ".bin"
	r, er := http.Get(url)
	if er == nil {
		if r.StatusCode == 200 {
			raw, _ = io.ReadAll(r.Body)
			r.Body.Close()
		} else {
			fmt.Println("webbtc.com StatusCode=", r.StatusCode)
		}
	}
	if er != nil {
		fmt.Println("webbtc.com:", er.Error())
	}
	return
}

// GetBlockFromBlockchainInfo downloads a raw block from blockchain.info.
func GetBlockFromBlockchainInfo(hash *btc.Uint256) (rawtx []byte) {
	url := "https://blockchain.info/block/" + hash.String() + "?format=hex"
	r, er := http.Get(url)
	if er == nil {
		if r.StatusCode == 200 {
			defer r.Body.Close()
			rawhex, _ := io.ReadAll(r.Body)
			rawtx, er = hex.DecodeString(string(rawhex))
		} else {
			fmt.Println("blockchain.info StatusCode=", r.StatusCode)
		}
	}
	if er != nil {
		fmt.Println("blockchain.info:", er.Error())
	}
	return
}

// GetBlockFromBlockstream downloads a raw block from blockstream
func GetBlockFromBlockstream(hash *btc.Uint256, api_url string) (raw []byte) {
	url := api_url + hash.String() + "/raw"
	r, er := http.Get(url)
	if er == nil {
		if r.StatusCode == 200 {
			raw, _ = io.ReadAll(r.Body)
			r.Body.Close()
		} else {
			fmt.Println("blockstream block StatusCode=", r.StatusCode)
		}
	}
	if er != nil {
		fmt.Println("blockstream block:", er.Error())
	}
	return
}

func IsBlockOK(raw []byte, hash *btc.Uint256) (bl *btc.Block) {
	var er error
	bl, er = btc.NewBlock(raw)
	if er != nil {
		return
	}
	if !bl.Hash.Equal(hash) {
		return nil
	}
	er = bl.BuildTxList()
	if er != nil {
		return nil
	}
	if !bl.MerkleRootMatch() {
		return nil
	}
	return
}

// GetBlockFromWeb downloads a raw block from a web server (try one after another).
func GetBlockFromWeb(hash *btc.Uint256) (bl *btc.Block) {
	var raw []byte

	raw = GetBlockFromBlockstream(hash, "https://blockstream.info/api/block/")
	if bl = IsBlockOK(raw, hash); bl != nil {
		//println("GetTxFromBlockstream - OK")
		return
	}

	raw = GetBlockFromBlockchainInfo(hash)
	if bl = IsBlockOK(raw, hash); bl != nil {
		//println("GetTxFromBlockchainInfo - OK")
		return
	}

	raw = GetBlockFromExplorer(hash)
	if bl = IsBlockOK(raw, hash); bl != nil {
		//println("GetTxFromExplorer - OK")
		return
	}

	raw = GetBlockFromWebBTC(hash)
	if bl = IsBlockOK(raw, hash); bl != nil {
		//println("GetTxFromWebBTC - OK")
		return
	}

	return
}
