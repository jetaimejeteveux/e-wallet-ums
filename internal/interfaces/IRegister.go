package interfaces

import (
	"context"

	"github.com/jetaimejeteveux/e-wallet-ums/models"
)

type IRegisterRepository interface {
	InsertUser(ctx context.Context, user *models.User) error
}

type IRegisterService interface {
	Register(ctx context.Context, request models.User) (interface{}, error)
}
