package chain

import (
	"github.com/codexnetwork/codex-go/config"
	"github.com/codexnetwork/codex-go/eosforce"
)

// NewAPIByCfg new api by cfg json file
func NewAPIByCfg(cfgPath string, isDebug bool) *eosforce.API {
	cfg := &config.ConfigData{}
	err := LoadJSONFile(cfgPath, cfg)
	if err != nil {
		panic(err)
	}

	api := &eosforce.API{}
	err = api.Init(cfg)

	api.Cfg.IsDebug = isDebug
	api.API.Debug = isDebug

	if err != nil {
		panic(err)
	}

	return api
}
