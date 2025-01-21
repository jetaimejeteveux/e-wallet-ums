package interfaces

import (
	"context"

	pb "github.com/jetaimejeteveux/e-wallet-ums/cmd/proto/tokenvalidation"
	"github.com/jetaimejeteveux/e-wallet-ums/helpers"
)

type ITokenValidationHandler interface {
	ValidateToken(ctx context.Context, req *pb.TokenRequest) (*pb.TokenResponse, error)
	pb.UnimplementedTokenValidationServer
}

type ITokenValidationService interface {
	ValidateToken(ctx context.Context, token string) (*helpers.ClaimToken, error)
}
