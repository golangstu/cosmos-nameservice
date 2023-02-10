package nameservice

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler 处理器
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		//理解是获取接口实例实际的类型指针，以此调用实例所有可调用的方法，包括接口方法及自有方法。
		//需要注意的是该写法必须与switch case联合使用，case中列出实现该接口的类型。
		switch msg := msg.(type) {
		case MsgSetName:
			return handlerSetName(ctx, keeper, msg)
		case MsgBuyName:
			return handlerBuyName(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized nameservice Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// handlerSetName 处理方法
func handlerSetName(ctx sdk.Context, keeper Keeper, msg MsgSetName) sdk.Result {
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.Name)) {
		return sdk.ErrUnauthorized("Incorrect Owner").Result()
	}
	keeper.SetName(ctx, msg.Name, msg.Value)
	return sdk.Result{}
}

func handlerBuyName(ctx sdk.Context, keeper Keeper, msg MsgBuyName) sdk.Result {
	//判断当前出价是否符合要去
	if keeper.GetPrice(ctx, msg.Name).IsAllGT(msg.Bid) {
		return sdk.ErrInsufficientCoins("Bid not high enough").Result()
	}
	//判断当前域名是否已被寄售
	if keeper.HasOwner(ctx, msg.Name) {
		_, err := keeper.coinKeeper.SendCoins(ctx, msg.Buyer, keeper.GetOwner(ctx, msg.Name), msg.Bid)
		if err != nil {
			return sdk.ErrInsufficientCoins("Buyer not enough balance").Result()
		}
	} else {
		//直接销毁
		_, _, err := keeper.coinKeeper.SubtractCoins(ctx, msg.Buyer, msg.Bid)
		if err != nil {
			return sdk.ErrInsufficientCoins("Buyer not enough balance").Result()
		}
	}
	//这里应当优化，增加是否寄售的功能 todo
	keeper.SetOwner(ctx, msg.Name, msg.Buyer)
	keeper.SetPrice(ctx, msg.Name, msg.Bid)
	return sdk.Result{}
}
