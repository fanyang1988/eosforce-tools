package chain

import (
	"github.com/codexnetwork/codex-go/eosforce"
	eos "github.com/eosforce/goeosforce"
	"github.com/pkg/errors"
)

// GetABI get abi from chain by account
func GetABI(api *eosforce.API, account string) (*eos.ABI, error) {
	out, err := api.GetABI(AN(account))
	if err != nil {
		return nil, errors.Wrapf(err, "get %s abi err", account)
	}

	return &out.ABI, nil
}
