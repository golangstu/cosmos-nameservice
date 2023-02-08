package app

import (
	mba "github.com/cosmos/cosmos-sdk/baseapp"
	dbm "github.com/tendermint/tm-db"
	"log"
)

const (
	appName = "nameservice"
)

type nameServiceApp struct {
	*mba.BaseApp
}

func NewNameServiceApp(logger log.Logger, db dbm.MemDB) *nameServiceApp {

}
