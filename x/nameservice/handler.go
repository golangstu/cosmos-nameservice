package nameservice

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler 处理器
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgSetName:
			return handlerSetName(ctx, keeper, msg)
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
