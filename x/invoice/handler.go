package invoice

import (
	"fmt"
	"net/url"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler creates a new handler
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgCreateInvoice:
			return handleMsgCreateInvoice(ctx, keeper, msg)
		case MsgEditInvoice:
			return handleMsgEditInvoice(ctx, keeper, msg)
		case MsgPayInvoice:
			return handleMsgPayInvoice(ctx, keeper, msg)
		case MsgFactorInvoice:
			return handleMsgFactorInvoice(ctx, keeper, msg)
		case MsgAddAdmin:
			return handleMsgAddAdmin(ctx, keeper, msg)
		case MsgRemoveAdmin:
			return handleMsgRemoveAdmin(ctx, keeper, msg)
		case MsgUpdateParams:
			return handleMsgUpdateParams(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized invoice message type: %T", msg)
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgCreateInvoice(ctx sdk.Context, keeper Keeper, msg MsgCreateInvoice) sdk.Result {
	if err := msg.ValidateBasic(); err != nil {
		return err.Result()
	}

	// parse url from string
	sourceURL, urlError := url.Parse(msg.Source)
	if urlError != nil {
		return ErrInvalidSourceURL(msg.Source).Result()
	}

	invoice, err := keeper.SubmitInvoice(ctx, msg.Body, msg.MarketplaceID, msg.Merchant, *sourceURL)
	if err != nil {
		return err.Result()
	}

	res, codecErr := ModuleCodec.MarshalJSON(invoice)
	if codecErr != nil {
		return sdk.ErrInternal(fmt.Sprintf("Marshal result error: %s", codecErr)).Result()
	}

	return sdk.Result{
		Data: res,
	}
}

func handleMsgEditInvoice(ctx sdk.Context, keeper Keeper, msg MsgEditInvoice) sdk.Result {
	if err := msg.ValidateBasic(); err != nil {
		return err.Result()
	}

	invoice, err := keeper.EditInvoice(ctx, msg.ID, msg.Body, msg.Editor)
	if err != nil {
		return err.Result()
	}

	res, codecErr := ModuleCodec.MarshalJSON(invoice)
	if codecErr != nil {
		return sdk.ErrInternal(fmt.Sprintf("Marshal result error: %s", codecErr)).Result()
	}

	return sdk.Result{
		Data: res,
	}
}

// Pay Invoice

func handleMsgPayInvoice(ctx sdk.Context, keeper Keeper, msg MsgPayInvoice) sdk.Result {
	if err := msg.ValidateBasic(); err != nil {
		return err.Result()
	}

	invoice, err := keeper.PayInvoice(ctx, msg.ID, msg.Body, msg.Editor)
	if err != nil {
		return err.Result()
	}

	res, codecErr := ModuleCodec.MarshalJSON(invoice)
	if codecErr != nil {
		return sdk.ErrInternal(fmt.Sprintf("Marshal result error: %s", codecErr)).Result()
	}

	return sdk.Result{
		Data: res,
	}
}

// Factor Invoice

func handleMsgFactorInvoice(ctx sdk.Context, keeper Keeper, msg MsgFactorInvoice) sdk.Result {
	if err := msg.ValidateBasic(); err != nil {
		return err.Result()
	}

	invoice, err := keeper.FactorInvoice(ctx, msg.ID, msg.Body, msg.Editor)
	if err != nil {
		return err.Result()
	}

	res, codecErr := ModuleCodec.MarshalJSON(invoice)
	if codecErr != nil {
		return sdk.ErrInternal(fmt.Sprintf("Marshal result error: %s", codecErr)).Result()
	}

	return sdk.Result{
		Data: res,
	}
}

func handleMsgAddAdmin(ctx sdk.Context, k Keeper, msg MsgAddAdmin) sdk.Result {
	if err := msg.ValidateBasic(); err != nil {
		return err.Result()
	}

	err := k.AddAdmin(ctx, msg.Admin, msg.Creator)
	if err != nil {
		return err.Result()
	}

	res, jsonErr := ModuleCodec.MarshalJSON(true)
	if jsonErr != nil {
		return sdk.ErrInternal(fmt.Sprintf("Marshal result error: %s", jsonErr)).Result()
	}

	return sdk.Result{
		Data: res,
	}
}

func handleMsgRemoveAdmin(ctx sdk.Context, k Keeper, msg MsgRemoveAdmin) sdk.Result {
	if err := msg.ValidateBasic(); err != nil {
		return err.Result()
	}

	err := k.RemoveAdmin(ctx, msg.Admin, msg.Remover)
	if err != nil {
		return err.Result()
	}

	res, jsonErr := ModuleCodec.MarshalJSON(true)
	if jsonErr != nil {
		return sdk.ErrInternal(fmt.Sprintf("Marshal result error: %s", jsonErr)).Result()
	}

	return sdk.Result{
		Data: res,
	}
}

func handleMsgUpdateParams(ctx sdk.Context, k Keeper, msg MsgUpdateParams) sdk.Result {
	if err := msg.ValidateBasic(); err != nil {
		return err.Result()
	}

	err := k.UpdateParams(ctx, msg.Updater, msg.Updates, msg.UpdatedFields)
	if err != nil {
		return err.Result()
	}

	res, jsonErr := ModuleCodec.MarshalJSON(true)
	if jsonErr != nil {
		return sdk.ErrInternal(fmt.Sprintf("Marshal result error: %s", jsonErr)).Result()
	}

	return sdk.Result{
		Data: res,
	}
}