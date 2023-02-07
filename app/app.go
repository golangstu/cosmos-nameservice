package app

import (
	mba "github.com/cosmos/cosmos-sdk/baseapp"
)

const (
	appName = "nameservice"
)

type nameServiceApp struct {
	*mba.BaseApp
}
