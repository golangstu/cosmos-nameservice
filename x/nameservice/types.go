package nameservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MinNamePrice 定义无人持有时的最低价格
var MinNamePrice = sdk.Coins{sdk.NewInt64Coin("nametoken", 1)}

// Whois 定义域名结构
type Whois struct {
	Value string         `json:"value"`
	Owner sdk.AccAddress `json:"owner"`
	Price sdk.Coins      `json:"price"`
}

// NewWhois 返回最低价格
func NewWhois() Whois {
	return Whois{
		Price: MinNamePrice,
	}
}
