package nameservice

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
)

type Keeper struct {
	coinKeeper bank.Keeper

	storeKey sdk.StoreKey

	cdc *codec.Codec
}

// NewKeeper 创建一个keeper实例
func NewKeeper(coinKeeper bank.Keeper, storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		coinKeeper: coinKeeper,
		storeKey:   storeKey,
		cdc:        cdc,
	}
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

// ResolveName - 返回解析过的域名的值
func (k Keeper) ResolveName(ctx sdk.Context, name string) string {
	return k.GetWhois(ctx, name).Value
}

// SetName - 设置域名信息
func (k Keeper) SetName(ctx sdk.Context, name string, value string) {
	whois := k.GetWhois(ctx, name)
	whois.Value = value
	k.SetWhois(ctx, name, whois)
}

// HasOwner - 返回该域名是否有人持有
func (k Keeper) HasOwner(ctx sdk.Context, name string) bool {
	return !k.GetWhois(ctx, name).Owner.Empty()
}

// GetOwner - 获得当前域名持有人的地址
func (k Keeper) GetOwner(ctx sdk.Context, name string) sdk.AccAddress {
	return k.GetWhois(ctx, name).Owner
}

// SetOwner - 设置域名持有人地址
func (k Keeper) SetOwner(ctx sdk.Context, name string, owner sdk.AccAddress) {
	whois := k.GetWhois(ctx, name)
	whois.Owner = owner
	k.SetWhois(ctx, name, whois)
}

// GetPrice - 获得域名的价格
func (k Keeper) GetPrice(ctx sdk.Context, name string) sdk.Coins {
	return k.GetWhois(ctx, name).Price
}

// SetPrice - 设置域名的价格
func (k Keeper) SetPrice(ctx sdk.Context, name string, price sdk.Coins) {
	whois := k.GetWhois(ctx, name)
	whois.Price = price
	k.SetWhois(ctx, name, whois)
}

// GetNamesIterator 获得该域名的迭代器
func (k Keeper) GetNamesIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte{})
}
