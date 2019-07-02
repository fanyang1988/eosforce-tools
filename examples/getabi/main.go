package main

import (
	"flag"
	"fmt"

	"github.com/fanyang1988/eosforce-tools/chain"
)

var account = flag.String("a", "eosio", "account to get abi")
var cfg = flag.String("cfg", "../eosforce.json", "cfg chain json file path")

func main() {
	flag.Parse()

	api := chain.NewAPIByCfg(*cfg, false)
	res, err := chain.GetABI(api, *account)
	if err != nil {
		fmt.Errorf("get abi err by %v", err.Error())
		return
	}

	fmt.Printf("abi %v", res)
}
