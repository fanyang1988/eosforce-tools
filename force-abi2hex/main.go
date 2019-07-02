package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	eos "github.com/eosforce/goeosforce"
)

var abiPath = flag.String("abi", "", "abi in json file path")

// loadJSONFile load a json file to obj
func loadJSONFile(path string, obj interface{}) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, obj)
}

func main() {
	flag.Parse()

	abi := &eos.ABI{}

	err := loadJSONFile(*abiPath, abi)
	if err != nil {
		fmt.Printf("ERROR: load json file %s err by %s\n", *abiPath, err.Error())
		os.Exit(1)
		return
	}

	var buffer bytes.Buffer
	encoder := eos.NewEncoder(&buffer)
	err = encoder.Encode(abi)

	if err != nil {
		fmt.Printf("ERROR: encoder %s err by %s\n", *abiPath, err.Error())
		os.Exit(2)
		return
	}

	fmt.Printf("%x", buffer.Bytes())
}
