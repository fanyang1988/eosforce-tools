package chain

import (
	"encoding/json"
	"io/ioutil"

	eos "github.com/eosforce/goeosforce"
)

// AN convert string to eos.AccountName
func AN(account string) eos.AccountName {
	return eos.AN(account)
}

// LoadJSONFile load a json file to obj
func LoadJSONFile(path string, obj interface{}) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, obj)
}
