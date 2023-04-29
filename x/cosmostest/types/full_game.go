package types

import (
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	"github.com/InnerPeace080/cosmos-test/x/cosmostest/rules"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetBlackAddress returns the black address from a StoredGame
func (storedGame StoredGame) GetBlackAddress() (black sdk.AccAddress, err error) {
	black, errBlack := sdk.AccAddressFromBech32(storedGame.Black)
	return black, sdkerrors.Wrapf(errBlack, ErrInvalidBlack.Error(), storedGame.Black)
}

// GetRedAddress returns the red address from a StoredGame
func (storedGame StoredGame) GetRedAddress() (red sdk.AccAddress, err error) {
	red, errRed := sdk.AccAddressFromBech32(storedGame.Red)
	return red, sdkerrors.Wrapf(errRed, ErrInvalidRed.Error(), storedGame.Red)
}

// ParseGame parses the game from a StoredGame
func (storedGame StoredGame) ParseGame() (game *rules.Game, err error) {
	board, errBoard := rules.Parse(storedGame.Board)
	if errBoard != nil {
		return nil, sdkerrors.Wrapf(errBoard, ErrGameNotParseable.Error())
	}
	board.Turn = rules.StringPieces[storedGame.Turn].Player
	if board.Turn.Color == "" {
		return nil, sdkerrors.Wrapf(fmt.Errorf("Turn: %s", storedGame.Turn), ErrGameNotParseable.Error())
	}
	return board, nil
}

// Validate checks that the storedGame is valid
func (storedGame StoredGame) Validate() (err error) {
	_, err = storedGame.GetBlackAddress()
	if err != nil {
		return err
	}
	_, err = storedGame.GetRedAddress()
	if err != nil {
		return err
	}
	_, err = storedGame.ParseGame()
	return err
}
