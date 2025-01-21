package api

import (
	"context"
	"fmt"

	pb "github.com/jetaimejeteveux/e-wallet-ums/cmd/proto/tokenvalidation"
	"github.com/jetaimejeteveux/e-wallet-ums/constants"
	"github.com/jetaimejeteveux/e-wallet-ums/helpers"
	"github.com/jetaimejeteveux/e-wallet-ums/internal/interfaces"
)

type TokenValidationHandler struct {
	TokenValidationService interfaces.ITokenValidationService
	pb.UnimplementedTokenValidationServer
}

func (h *TokenValidationHandler) ValidateToken(ctx context.Context, req *pb.TokenRequest) (*pb.TokenResponse, error) {
	var (
		token = req.Token
		log   = helpers.Logger
	)

	if token == "" {
		err := fmt.Errorf("token is empty")
		log.Error("token is empty = ", err)
		return &pb.TokenResponse{
			Message: err.Error(),
		}, nil
	}

	claimToken, err := h.TokenValidationService.ValidateToken(ctx, token)
	if err != nil {
		return &pb.TokenResponse{
			Message: err.Error(),
		}, nil
	}

	return &pb.TokenResponse{
		Message: constants.SuccessMessage,
		Data: &pb.UserData{
			UserId:   int64(claimToken.UserID),
			Username: claimToken.Username,
			FullName: claimToken.Fullname,
			Email:    claimToken.Email,
		},
	}, nil
}
