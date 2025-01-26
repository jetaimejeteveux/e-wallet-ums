package interfaces

import (
	"context"

	externalModels "github.com/jetaimejeteveux/e-wallet-ums/external/models"
)

type IExternal interface {
	CreateWallet(ctx context.Context, userID int) (*externalModels.Wallet, error)
}
