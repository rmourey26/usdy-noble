package keeper

import (
	"context"
	"errors"
	"fmt"

	"cosmossdk.io/collections"
	sdkerrors "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noble-assets/aura/x/aura/types"
)

var _ types.MsgServer = &msgServer{}

type msgServer struct {
	*Keeper
}

func NewMsgServer(keeper *Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

func (k msgServer) Burn(ctx context.Context, msg *types.MsgBurn) (*types.MsgBurnResponse, error) {
	allowance, err := k.Burners.Get(ctx, msg.Signer)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, types.ErrInvalidBurner
		}

		return nil, sdkerrors.Wrapf(err, "unable to retrieve burner from state")
	}
	if allowance.LT(msg.Amount) {
		return nil, sdkerrors.Wrapf(types.ErrInsufficientAllowance, "burner %s has an allowance of %s", msg.Signer, allowance.String())
	}

	from, err := k.accountKeeper.AddressCodec().StringToBytes(msg.From)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "unable to decode account address %s", msg.From)
	}

	coins := sdk.NewCoins(sdk.NewCoin(k.Denom, msg.Amount))
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, from, types.ModuleName, coins)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "unable to transfer from user to module")
	}
	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, coins)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "unable to burn from module")
	}

	err = k.Burners.Set(ctx, msg.Signer, allowance.Sub(msg.Amount))
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "unable to update burner's allowance")
	}

	// TODO(@john): Do we emit an event here?
	return &types.MsgBurnResponse{}, nil
}

func (k msgServer) Mint(ctx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	allowance, err := k.Minters.Get(ctx, msg.Signer)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, types.ErrInvalidMinter
		}

		return nil, sdkerrors.Wrapf(err, "unable to retrieve minter from state")
	}
	if allowance.LT(msg.Amount) {
		return nil, sdkerrors.Wrapf(types.ErrInsufficientAllowance, "minter %s has an allowance of %s", msg.Signer, allowance.String())
	}

	to, err := k.accountKeeper.AddressCodec().StringToBytes(msg.To)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "unable to decode account address %s", msg.To)
	}

	coins := sdk.NewCoins(sdk.NewCoin(k.Denom, msg.Amount))
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, coins)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "unable to mint to module")
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, to, coins)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "unable to transfer from module to user")
	}

	err = k.Minters.Set(ctx, msg.Signer, allowance.Sub(msg.Amount))
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "unable to update minter's allowance")
	}

	// TODO(@john): Do we emit an event here?
	return &types.MsgMintResponse{}, nil
}

func (k msgServer) Pause(ctx context.Context, msg *types.MsgPause) (*types.MsgPauseResponse, error) {
	has, err := k.Pausers.Has(ctx, msg.Signer)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "unable to retrieve pauser from state")
	}
	if !has {
		return nil, types.ErrInvalidPauser
	}

	if paused, _ := k.Paused.Get(ctx); paused {
		return nil, errors.New("module is already paused")
	}

	err = k.Paused.Set(ctx, true)
	if err != nil {
		return nil, errors.New("unable to set paused state")
	}

	return &types.MsgPauseResponse{}, k.eventService.EventManager(ctx).Emit(ctx, &types.Paused{
		Account: msg.Signer,
	})
}

func (k msgServer) Unpause(ctx context.Context, msg *types.MsgUnpause) (*types.MsgUnpauseResponse, error) {
	owner, err := k.Owner.Get(ctx)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "unable to retrieve owner from state")
	}
	if msg.Signer != owner {
		return nil, sdkerrors.Wrapf(types.ErrInvalidOwner, "expected %s, got %s", owner, msg.Signer)
	}

	if paused, _ := k.Paused.Get(ctx); !paused {
		return nil, errors.New("module is already unpaused")
	}

	err = k.Paused.Set(ctx, false)
	if err != nil {
		return nil, errors.New("unable to set paused state")
	}

	return &types.MsgUnpauseResponse{}, k.eventService.EventManager(ctx).Emit(ctx, &types.Unpaused{
		Account: msg.Signer,
	})
}

func (k msgServer) TransferOwnership(ctx context.Context, msg *types.MsgTransferOwnership) (*types.MsgTransferOwnershipResponse, error) {
	owner, err := k.Owner.Get(ctx)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "unable to retrieve owner from state")
	}
	if msg.Signer != owner {
		return nil, sdkerrors.Wrapf(types.ErrInvalidOwner, "expected %s, got %s", owner, msg.Signer)
	}

	err = k.PendingOwner.Set(ctx, msg.NewOwner)
	if err != nil {
		return nil, errors.New("unable to set pending owner state")
	}

	return &types.MsgTransferOwnershipResponse{}, k.eventService.EventManager(ctx).Emit(ctx, &types.OwnershipTransferStarted{
		OldOwner: owner,
		NewOwner: msg.NewOwner,
	})
}

func (k msgServer) AcceptOwnership(ctx context.Context, msg *types.MsgAcceptOwnership) (*types.MsgAcceptOwnershipResponse, error) {
	pendingOwner, err := k.PendingOwner.Get(ctx)
	if err != nil {
		return nil, errors.New("there is no pending owner")
	}
	if msg.Signer != pendingOwner {
		return nil, sdkerrors.Wrapf(types.ErrInvalidPendingOwner, "expected %s, got %s", pendingOwner, msg.Signer)
	}

	err = k.Owner.Set(ctx, pendingOwner)
	if err != nil {
		return nil, errors.New("unable to set owner state")
	}
	err = k.PendingOwner.Remove(ctx)
	if err != nil {
		return nil, errors.New("unable to remove pending owner state")
	}

	return &types.MsgAcceptOwnershipResponse{}, nil
}

func (k msgServer) AddBurner(ctx context.Context, msg *types.MsgAddBurner) (*types.MsgAddBurnerResponse, error) {
	owner, err := k.Owner.Get(ctx)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "unable to retrieve owner from state")
	}
	if msg.Signer != owner {
		return nil, sdkerrors.Wrapf(types.ErrInvalidOwner, "expected %s, got %s", owner, msg.Signer)
	}

	has, err := k.Burners.Has(ctx, msg.Burner)
	if err != nil || has {
		return nil, fmt.Errorf("%s is already a burner", msg.Burner)
	}

	// TODO(@john): Validate allowance?

	err = k.Burners.Set(ctx, msg.Burner, msg.Allowance)
	if err != nil {
		return nil, errors.New("unable to set burner in state")
	}

	// TODO(@john): Do we emit an event here?
	return &types.MsgAddBurnerResponse{}, nil
}

func (k msgServer) RemoveBurner(ctx context.Context, msg *types.MsgRemoveBurner) (*types.MsgRemoveBurnerResponse, error) {
	owner, err := k.Owner.Get(ctx)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "unable to retrieve owner from state")
	}
	if msg.Signer != owner {
		return nil, sdkerrors.Wrapf(types.ErrInvalidOwner, "expected %s, got %s", owner, msg.Signer)
	}

	has, err := k.Burners.Has(ctx, msg.Burner)
	if err != nil || !has {
		return nil, fmt.Errorf("%s is not a burner", msg.Burner)
	}

	err = k.Burners.Remove(ctx, msg.Burner)
	if err != nil {
		return nil, errors.New("unable to remove burner from state")
	}

	// TODO(@john): Do we emit an event here?
	return &types.MsgRemoveBurnerResponse{}, nil
}

func (k msgServer) SetBurnerAllowance(ctx context.Context, msg *types.MsgSetBurnerAllowance) (*types.MsgSetBurnerAllowanceResponse, error) {
	owner, err := k.Owner.Get(ctx)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "unable to retrieve owner from state")
	}
	if msg.Signer != owner {
		return nil, sdkerrors.Wrapf(types.ErrInvalidOwner, "expected %s, got %s", owner, msg.Signer)
	}

	has, err := k.Burners.Has(ctx, msg.Burner)
	if err != nil || !has {
		return nil, fmt.Errorf("%s is not a burner", msg.Burner)
	}

	// TODO(@john): Validate allowance?

	err = k.Burners.Set(ctx, msg.Burner, msg.Allowance)
	if err != nil {
		return nil, errors.New("unable to set burner allowance in state")
	}

	// TODO(@john): Do we emit an event here?
	return &types.MsgSetBurnerAllowanceResponse{}, nil
}

func (k msgServer) AddMinter(ctx context.Context, msg *types.MsgAddMinter) (*types.MsgAddMinterResponse, error) {
	owner, err := k.Owner.Get(ctx)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "unable to retrieve owner from state")
	}
	if msg.Signer != owner {
		return nil, sdkerrors.Wrapf(types.ErrInvalidOwner, "expected %s, got %s", owner, msg.Signer)
	}

	has, err := k.Minters.Has(ctx, msg.Minter)
	if err != nil || has {
		return nil, fmt.Errorf("%s is already a minter", msg.Minter)
	}

	// TODO(@john): Validate allowance?

	err = k.Minters.Set(ctx, msg.Minter, msg.Allowance)
	if err != nil {
		return nil, errors.New("unable to set minter in state")
	}

	// TODO(@john): Do we emit an event here?
	return &types.MsgAddMinterResponse{}, nil
}

func (k msgServer) RemoveMinter(ctx context.Context, msg *types.MsgRemoveMinter) (*types.MsgRemoveMinterResponse, error) {
	owner, err := k.Owner.Get(ctx)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "unable to retrieve owner from state")
	}
	if msg.Signer != owner {
		return nil, sdkerrors.Wrapf(types.ErrInvalidOwner, "expected %s, got %s", owner, msg.Signer)
	}

	has, err := k.Minters.Has(ctx, msg.Minter)
	if err != nil || !has {
		return nil, fmt.Errorf("%s is not a minter", msg.Minter)
	}

	err = k.Minters.Remove(ctx, msg.Minter)
	if err != nil {
		return nil, errors.New("unable to remove minter from state")
	}

	// TODO(@john): Do we emit an event here?
	return &types.MsgRemoveMinterResponse{}, nil
}

func (k msgServer) SetMinterAllowance(ctx context.Context, msg *types.MsgSetMinterAllowance) (*types.MsgSetMinterAllowanceResponse, error) {
	owner, err := k.Owner.Get(ctx)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "unable to retrieve owner from state")
	}
	if msg.Signer != owner {
		return nil, sdkerrors.Wrapf(types.ErrInvalidOwner, "expected %s, got %s", owner, msg.Signer)
	}

	has, err := k.Minters.Has(ctx, msg.Minter)
	if err != nil || !has {
		return nil, fmt.Errorf("%s is not a minter", msg.Minter)
	}

	// TODO(@john): Validate allowance?

	err = k.Minters.Set(ctx, msg.Minter, msg.Allowance)
	if err != nil {
		return nil, errors.New("unable to set minter allowance in state")
	}

	// TODO(@john): Do we emit an event here?
	return &types.MsgSetMinterAllowanceResponse{}, nil
}

func (k msgServer) AddPauser(ctx context.Context, msg *types.MsgAddPauser) (*types.MsgAddPauserResponse, error) {
	owner, err := k.Owner.Get(ctx)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "unable to retrieve owner from state")
	}
	if msg.Signer != owner {
		return nil, sdkerrors.Wrapf(types.ErrInvalidOwner, "expected %s, got %s", owner, msg.Signer)
	}

	has, err := k.Pausers.Has(ctx, msg.Pauser)
	if err != nil || has {
		return nil, fmt.Errorf("%s is already a pauser", msg.Pauser)
	}

	err = k.Pausers.Set(ctx, msg.Pauser)
	if err != nil {
		return nil, errors.New("unable to set pauser in state")
	}

	// TODO(@john): Do we emit an event here?
	return &types.MsgAddPauserResponse{}, nil
}

func (k msgServer) RemovePauser(ctx context.Context, msg *types.MsgRemovePauser) (*types.MsgRemovePauserResponse, error) {
	owner, err := k.Owner.Get(ctx)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "unable to retrieve owner from state")
	}
	if msg.Signer != owner {
		return nil, sdkerrors.Wrapf(types.ErrInvalidOwner, "expected %s, got %s", owner, msg.Signer)
	}

	has, err := k.Pausers.Has(ctx, msg.Pauser)
	if err != nil || !has {
		return nil, fmt.Errorf("%s is not a pauser", msg.Pauser)
	}

	err = k.Pausers.Remove(ctx, msg.Pauser)
	if err != nil {
		return nil, errors.New("unable to remove pauser from state")
	}

	// TODO(@john): Do we emit an event here?
	return &types.MsgRemovePauserResponse{}, nil
}
