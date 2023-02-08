package nameservice

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
)

type Keeper struct {
	coinKeeper bank.Keeper

	storeKey sdk.StoreKey

	cdc codec.Codec
}

// SetWhois 解析域名值信息
func (k Keeper) SetWhois(ctx sdk.Context, name string, whois Whois) {
	// 如果拥有人为空
	if whois.Owner.Empty() {
		return
	}
	// 获取storeKey的上下文对象
	store := ctx.KVStore(k.storeKey)
	// 将name和value编码后写入
	store.Set([]byte(name), k.cdc.MustMarshalBinaryBare(whois))
}

// GetWhois 查询域名信息
func (k Keeper) GetWhois(ctx sdk.Context, name string) Whois {
	// 获取storeKey的上下文对象
	store := ctx.KVStore(k.storeKey)
	//判断是否存在这个key
	if !store.Has([]byte(name)) {
		return NewWhois()
	}
	value := store.Get([]byte(name))
	//定义一个无引用的对象
	var whois Whois
	// 解码，与引用的方式传参
	k.cdc.MustUnmarshalBinaryBare(value, &whois)
	return whois
}
