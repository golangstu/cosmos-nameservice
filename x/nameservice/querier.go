package nameservice

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"strings"
)

const (
	QueryResolve = "resolve"
	QueryWhois   = "whois"
	QueryNames   = "names"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryResolve:
			return queryResolve(ctx, path[1:], req, keeper)
		case QueryWhois:
			return queryWhois(ctx, path[1:], req, keeper)
		case QueryNames:
			return queryNames(ctx, req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown nameservice query endpoint")
		}
	}

}

func queryResolve(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	name := path[0]
	value := keeper.ResolveName(ctx, name)
	if value == "" {
		return []byte{}, sdk.ErrUnknownRequest("can not resolve name")
	}
	bt, err2 := codec.MarshalJSONIndent(keeper.cdc, QueryResResolve{value})
	if err2 != nil {
		panic("could not marshal result to JSON")
	}
	return bt, nil
}

func queryWhois(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	name := path[0]
	whois := keeper.GetWhois(ctx, name)
	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, whois)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}
	return bz, nil
}

func queryNames(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	var nameList QueryResNames
	it := keeper.GetNamesIterator(ctx)
	for ; it.Valid(); it.Next() {
		name := string(it.Key())
		nameList = append(nameList, name)
	}
	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, nameList)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}
	return bz, nil
}

/*
按照惯例，每个输出类型都应该是 JSON marshallable 和 stringable（实现 Golang fmt.Stringer 接口）。 返回的字节应该是输出结果的JSON编码。
因此，对于输出类型的解析，我们将解析字符串包装在一个名为 QueryResResolve 的结构中，该结构既是JSON marshallable 的又有.String（）方法。
对于 Whois 的输出，正常的 Whois 结构已经是 JSON marshallable 的，但我们需要在其上添加.String（）方法。
名称查询的输出也一样，[]字符串本身已经可 marshallable ，但我们需要在其上添加.String（）方法。
*/
// implement fmt.Stringer
func (w Whois) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Owner: %s
Value: %s
Price: %s`, w.Owner, w.Value, w.Price))
}

type QueryResNames []string

//实现string接口，换行
func (query QueryResNames) String() string {
	return strings.Join(query, "\n")
}

type QueryResResolve struct {
	Value string `json:"value"`
}
